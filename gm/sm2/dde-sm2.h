// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

#ifndef DDE_SM2_H
#define DDE_SM2_H

#include <stdint.h>
#include <stddef.h>

typedef struct _sm2_context sm2_context;

sm2_context *new_sm2_context();
void free_sm2_context(sm2_context *context);

const char* get_sm2_public_key(sm2_context *context);
const char* get_sm2_private_key(sm2_context *context);

int get_ciphertext_size(const sm2_context *context, size_t plen);
int get_plaintext_size(const uint8_t *ctext, size_t clen);

int encrypt(const sm2_context *context, const uint8_t *ptext, size_t psize, uint8_t *ctext, size_t csize);
int decrypt(const sm2_context *context, const uint8_t *ctext, size_t csize, uint8_t *ptext, size_t psize);

#endif
