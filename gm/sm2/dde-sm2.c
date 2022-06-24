#include <openssl/bio.h>
#include <openssl/pem.h>
#include <openssl/evp.h>
#include <openssl/ec.h>
#include <openssl/bn.h>
#include <openssl/crypto.h>
#include <openssl/sm2.h>

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

int get_ciphertext_size(const sm2_context *context, size_t plen) {
    size_t ret = 0;
    if (1 == sm2_ciphertext_size(context->key, EVP_sm3(), plen, &ret)) {
	return (int)ret;
    }

    return -1;
}

int get_plaintext_size(const uint8_t *ctext, size_t clen) {
    size_t ret = 0;
    if (1 == sm2_plaintext_size(ctext, clen, &ret)) {
	return (int)ret;
    }

    return -1;
}

int encrypt(const sm2_context* context, const uint8_t *ptext, size_t psize, uint8_t *ctext, size_t csize) {
    if (1 == sm2_encrypt(context->key, EVP_sm3(), ptext, psize, ctext, &csize)) {
	return (int)csize;
    }

    return -1;
}

int decrypt(const sm2_context* context, const uint8_t *ctext, size_t clen, uint8_t *ptext, size_t psize) {
    if (1 == sm2_decrypt(context->key, EVP_sm3(), ctext, clen, ptext, &psize)) {
	return (int)psize;
    }

    return -1;
}
