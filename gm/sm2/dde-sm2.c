// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

#include <openssl/bio.h>
#include <openssl/pem.h>
#include <openssl/evp.h>
#include <openssl/ec.h>
#include <openssl/bn.h>
#include <openssl/crypto.h>

#include "dde-sm2.h"

struct _sm2_context {
    EC_KEY* key;
    char* private_key;
    char* publick_key;
};

static EC_KEY* gen_ec_key() {
    EC_KEY* ec_key = NULL;
    EC_GROUP* ec_group = NULL;

    if (NULL == (ec_key = EC_KEY_new())) {
	return NULL;
    }

    if (NULL == (ec_group = EC_GROUP_new_by_curve_name(NID_sm2))) {
	EC_KEY_free(ec_key);
	return NULL;
    }

    if (1 != EC_KEY_set_group(ec_key, ec_group)) {
	EC_GROUP_free(ec_group);
	EC_KEY_free(ec_key);
	return NULL;
    }

    if (!EC_KEY_generate_key(ec_key)) {
	EC_GROUP_free(ec_group);
	EC_KEY_free(ec_key);
	return NULL;
    }

    return ec_key;
}

static char* get_public_key(EC_KEY *key) {
    BIO* bio = NULL;
    char *ret = NULL;
    size_t len = 0;

    bio = BIO_new(BIO_s_mem());
    PEM_write_bio_EC_PUBKEY(bio, key);

    len = BIO_pending(bio);
    ret = OPENSSL_zalloc(len+1);
    if (ret != NULL) {
	BIO_read(bio, ret, len);
	ret[len] = '\0';
    }

    BIO_free(bio);

    return ret;
}

static char* get_private_key(EC_KEY *key) {
    BIO* bio = NULL;
    char *ret = NULL;
    size_t len = 0;

    bio = BIO_new(BIO_s_mem());
    PEM_write_bio_ECPrivateKey(bio, key, NULL, NULL, 0, NULL, NULL);

    len = BIO_pending(bio);
    ret = OPENSSL_zalloc(len+1);
    if (ret != NULL) {
	BIO_read(bio, ret, len);
	ret[len] = '\0';
    }

    BIO_free(bio);

    return ret;
}

/*openssl sm2 cipher evp using*/
static int openssl_evp_sm2_encrypt(EC_KEY *ec_key,
                                   const unsigned char *plain_text, size_t plain_len,
                                   unsigned char *cipher_text, size_t *cipher_len)
{
    int ret = 0;
    BIO *bp = NULL;
    EVP_PKEY* public_evp_key = NULL;
    EVP_PKEY_CTX *ctx = NULL;

    /*Check the user input.*/
    if (plain_text == NULL || plain_len == 0) {
        ret = -1;
        return ret;
    }

    //OpenSSL_add_all_algorithms();
    bp = BIO_new(BIO_s_mem());
    if (bp == NULL) {
        printf("BIO_new is failed.\n");
        ret = -1;
        return ret;
    }

    if (ec_key == NULL) {
        ret = -1;
        printf("open_public_key failed to PEM_read_bio_EC_PUBKEY Failed, ret=%d\n", ret);
        goto finish;
    }
    public_evp_key = EVP_PKEY_new();
    if (public_evp_key == NULL) {
        ret = -1;
        printf("open_public_key EVP_PKEY_new failed\n");
        goto finish;
    }
    ret = EVP_PKEY_set1_EC_KEY(public_evp_key, ec_key);
    if (ret != 1) {
        ret = -1;
        printf("EVP_PKEY_set1_EC_KEY failed\n");
        goto finish;
    }
    ret = EVP_PKEY_set_alias_type(public_evp_key, EVP_PKEY_SM2);
    if (ret != 1) {
        printf("EVP_PKEY_set_alias_type to EVP_PKEY_SM2 failed! ret = %d\n", ret);
        ret = -1;
        goto finish;
    }
    /*modifying a EVP_PKEY to use a different set of algorithms than the default.*/

    /*do cipher.*/
    ctx = EVP_PKEY_CTX_new(public_evp_key, NULL);
    if (ctx == NULL) {
        ret = -1;
        printf("EVP_PKEY_CTX_new failed\n");
        goto finish;
    }
    ret = EVP_PKEY_encrypt_init(ctx);
    if (ret < 0) {
        printf("sm2_pubkey_encrypt failed to EVP_PKEY_encrypt_init. ret = %d\n", ret);
        goto finish;
    }
    ret = EVP_PKEY_encrypt(ctx, cipher_text, cipher_len, plain_text, plain_len);
    if (ret < 0) {
        printf("sm2_pubkey_encrypt failed to EVP_PKEY_encrypt. ret = %d\n", ret);
        goto finish;
    }
    ret = 0;

finish:
    if (public_evp_key != NULL)
        EVP_PKEY_free(public_evp_key);
    if (ctx != NULL)
        EVP_PKEY_CTX_free(ctx);
    if (bp != NULL)
        BIO_free(bp);

    return ret;
}

/*openssl sm2 decrypt evp using*/
static int openssl_evp_sm2_decrypt(EC_KEY* ec_key,
                                   const unsigned char *cipher_text, size_t cipher_len,
                                   unsigned char *plain_text, size_t *plain_len)
{
    int ret = 0;
    size_t out_len = 0;
    BIO *bp = NULL;
    EVP_PKEY* private_evp_key = NULL;
    EVP_PKEY_CTX *ctx = NULL;

    /*Check the user input.*/
    if (cipher_len == 0 || cipher_text == NULL) {
        ret = -1;
        return ret;
    }

    //OpenSSL_add_all_algorithms();
    bp = BIO_new(BIO_s_mem());
    if (bp == NULL) {
        printf("BIO_new is failed.\n");
        ret = -1;
        return ret;
    }

    if (ec_key == NULL) {
        ret = -1;
        printf("open_private_key failed to PEM_read_bio_ECPrivateKey Failed, ret=%d\n", ret);
        goto finish;
    }
    private_evp_key = EVP_PKEY_new();
    if (private_evp_key == NULL) {
        ret = -1;
        printf("open_public_key EVP_PKEY_new failed\n");
        goto finish;
    }
    ret = EVP_PKEY_set1_EC_KEY(private_evp_key, ec_key);
    if (ret != 1) {
        ret = -1;
        printf("EVP_PKEY_set1_EC_KEY failed\n");
        goto finish;
    }
    ret = EVP_PKEY_set_alias_type(private_evp_key, EVP_PKEY_SM2);
    if (ret != 1) {
        printf("EVP_PKEY_set_alias_type to EVP_PKEY_SM2 failed! ret = %d\n", ret);
        ret = -1;
        goto finish;
    }
    /*modifying a EVP_PKEY to use a different set of algorithms than the default.*/

    /*do cipher.*/
    ctx = EVP_PKEY_CTX_new(private_evp_key, NULL);
    if (ctx == NULL) {
        ret = -1;
        printf("EVP_PKEY_CTX_new failed\n");
        goto finish;
    }
    ret = EVP_PKEY_decrypt_init(ctx);
    if (ret < 0) {
        printf("sm2 private_key decrypt failed to EVP_PKEY_decrypt_init. ret = %d\n", ret);
        goto finish;
    }

    ret = EVP_PKEY_decrypt(ctx, plain_text, plain_len, cipher_text, cipher_len);
    if (ret < 0) {
        printf("sm2_prikey_decrypt failed to EVP_PKEY_decrypt. ret = %d\n", ret);
        goto finish;
    }
    ret = 0;
finish:
    if (private_evp_key != NULL)
        EVP_PKEY_free(private_evp_key);
    if (ctx != NULL)
        EVP_PKEY_CTX_free(ctx);
    if (bp != NULL)
        BIO_free(bp);

    return ret;
}


sm2_context* new_sm2_context() {
    EC_KEY* key = gen_ec_key();
    if (key == NULL) {
	return NULL;
    }

    sm2_context *ret = malloc(sizeof(sm2_context));
    if (ret == NULL) {
	EC_KEY_free(key);
	return NULL;
    }
    ret->key = key;
    ret->private_key = get_private_key(ret->key);
    ret->publick_key = get_public_key(ret->key);

    return ret;
}

void free_sm2_context(sm2_context* context) {
    if (context->key) {
	EC_KEY_free(context->key);
    }
    if (context->publick_key) {
	OPENSSL_free(context->publick_key);
    }
    if (context->private_key) {
	OPENSSL_free(context->private_key);
    }
    free(context);
}

const char* get_sm2_public_key(sm2_context* context) {
    return context->publick_key;
}

const char* get_sm2_private_key(sm2_context* context) {
    return context->private_key;
}

int get_ciphertext_size(const sm2_context *context, const uint8_t *ptext, size_t plen) {
    size_t ret = 0;
    if (0 == openssl_evp_sm2_encrypt(context->key, ptext, plen, NULL, &ret)) {
	return (int)ret;
    }

    return -1;
}

int get_plaintext_size(const sm2_context *context, const uint8_t *ctext, size_t clen) {
    size_t ret = 0;
    if (0 == openssl_evp_sm2_decrypt(context->key, ctext, clen, NULL, &ret)) {
	return (int)ret;
    }

    return -1;
}

int encrypt(const sm2_context* context, const uint8_t *ptext, size_t psize, uint8_t *ctext, size_t csize) {
    if (0 == openssl_evp_sm2_encrypt(context->key, ptext, psize, ctext, &csize)) {
	return (int)csize;
    }

    return -1;
}

int decrypt(const sm2_context* context, const uint8_t *ctext, size_t clen, uint8_t *ptext, size_t psize) {
    if (0 == openssl_evp_sm2_decrypt(context->key, ctext, clen, ptext, &psize)) {
	return (int)psize;
    }

    return -1;
}
