/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

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
