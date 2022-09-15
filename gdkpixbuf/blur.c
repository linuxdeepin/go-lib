// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

#include <stdio.h>
#include <glib.h>
#include <gdk-pixbuf/gdk-pixbuf.h>
#include "gaussianiir2d.h"

int blur(GdkPixbuf *pixbuf, double sigma, double numsteps) {
    int width = gdk_pixbuf_get_width(pixbuf);
    int height = gdk_pixbuf_get_height(pixbuf);
    int rowstride = gdk_pixbuf_get_rowstride(pixbuf);
    int n_channels = gdk_pixbuf_get_n_channels(pixbuf);
    guchar *image_data = gdk_pixbuf_get_pixels(pixbuf);
    gaussianiir2d_pixbuf_c(image_data, width, height,
                           rowstride, n_channels, sigma, numsteps);
    return TRUE;
}
