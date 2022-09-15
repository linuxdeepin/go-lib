// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

#ifndef _BLUR_H
#define _BLUR_H 1

#include <gdk-pixbuf/gdk-pixbuf.h>

int blur(GdkPixbuf *src_pixbuf, double sigma, double numsteps);

#endif /* _BLUR_H */
