/*
 * Copyright (C) 2014 ~ 2017 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

#ifndef _GDK_PIXBUF_UTILS_H
#define _GDK_PIXBUF_UTILS_H 1

#include <stdbool.h>
#include <X11/Xlib.h>
#include <gdk-pixbuf/gdk-pixbuf.h>

int init_gdk_xlib();
int init_gdk();
const char *get_image_format(const char *img_file);
int get_image_size(const char *img_file, int *width, int *height);
int get_dominant_color(const GdkPixbuf *pixbuf, double *r, double *g, double *b);
int save(GdkPixbuf *pixbuf, const char *dest_file, const char *format);
GdkPixbuf *new_pixbuf_from_file(const char *img_file);
Pixmap convert_pixbuf_to_xpixmap(GdkPixbuf *pixbuf);
GdkPixbuf *convert_xpixmap_to_pixbuf(const Pixmap xpixmap, int width, int height);
GdkPixbuf *copy_area_simple(const GdkPixbuf *src_pixbuf, int src_x, int src_y, int width, int height);
GdkPixbuf *screenshot();

#endif /* _GDK_PIXBUF_UTILS_H */
