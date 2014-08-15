/**
 * Copyright (c) 2014 Deepin, Inc.
 *               2014 Xu FaSheng
 *
 * Author:      Xu FaSheng <fasheng.xu@gmail.com>
 * Maintainer:  Xu FaSheng <fasheng.xu@gmail.com>
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License as
 * published by the Free Software Foundation; either version 3, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; see the file COPYING.  If not, write to
 * the Free Software Foundation, Inc., 51 Franklin Street, Fifth
 * Floor, Boston, MA 02110-1301, USA.
 **/

#include <stdio.h>
#include <glib.h>
#include <gdk/gdk.h>
#include <gdk/gdkx.h>
#include <X11/Xlib.h>
#include <gdk-pixbuf/gdk-pixbuf.h>
#include <gdk-pixbuf-xlib/gdk-pixbuf-xlib.h>

int init_gdk_xlib() {
        XInitThreads();         /* should be called before gdk_init() */
        if (!init_gdk()) {
                return FALSE;
        }
        gdk_pixbuf_xlib_init(GDK_DISPLAY_XDISPLAY(gdk_display_get_default()), 0);
        return TRUE;
}

int init_gdk() {
        return gdk_init_check(NULL, NULL);
}

const char *get_image_format(const char *img_file) {
        GdkPixbufFormat *format = gdk_pixbuf_get_file_info(img_file, NULL, NULL);
        if (format) {
                return (char *)gdk_pixbuf_format_get_name(format);
        }
        return NULL;
}

int get_image_size(const char *img_file, int *width, int *height) {
        GdkPixbufFormat *format = gdk_pixbuf_get_file_info(img_file, width, height);
        if (format) {
                return TRUE;
        }
        return FALSE;
}

int get_dominant_color(const GdkPixbuf *pixbuf, double *r, double *g, double *b) {
        if (pixbuf == NULL) {
                g_warning("pixbuf is NULL\n");
                return FALSE;
        }
        guint length = 0;
        guchar *pixels = gdk_pixbuf_get_pixels_with_length(pixbuf, &length);
        if (length == 0) {
                g_warning("zero length of pixbuf\n");
                return FALSE;
        }

        // calculate dominant color of pixbuf
        long long sum_r = 0;
        long long sum_g = 0;
        long long sum_b = 0;
        long count = 0;
        int skip = gdk_pixbuf_get_n_channels(pixbuf);
        guint i = 0;
        for (i=0; i<length; i += skip) {
                if (skip == 4 && pixels[i+3] < 125) {
                        continue;
                }
                sum_r += pixels[i];
                sum_g += pixels[i+1];
                sum_b += pixels[i+2];
                count++;
        }
        if (count == 0) {
                g_warning("count is zero when calculating dominant color of pixbuf\n");
                return FALSE;
        }
        *r = sum_r / count;
        *g = sum_g / count;
        *b = sum_b / count;
        return TRUE;
}

int save(GdkPixbuf *pixbuf, const char *dest_file, const char *format) {
        GError *err = NULL;
        gdk_pixbuf_save (pixbuf, dest_file, format, &err, NULL);
        if (err) {
                g_debug("save gdk pixbuf to file failed: %s", err->message);
                g_error_free(err);
                return FALSE;
        }
        return TRUE;
}

GdkPixbuf *new_pixbuf_from_file(const char *img_file) {
        GError *err = NULL;
        GdkPixbuf *pixbuf = gdk_pixbuf_new_from_file(img_file, &err);
        if (err) {
                g_debug("new gdk pixbuf from file failed: %s", err->message);
                g_error_free(err);
                return NULL;
        }
        return pixbuf;
}

Pixmap convert_pixbuf_to_xpixmap(GdkPixbuf *pixbuf) {
        Pixmap xpixmap = 0;
        gdk_pixbuf_xlib_render_pixmap_and_mask(pixbuf, &xpixmap, NULL, 0);
        return xpixmap;
}

GdkPixbuf *convert_xpixmap_to_pixbuf(const Pixmap xpixmap, int width, int height) {
        return gdk_pixbuf_xlib_get_from_drawable(NULL, (Drawable)xpixmap,
                                                 xlib_rgb_get_cmap(),
                                                 xlib_rgb_get_visual(),
                                                 0, 0, 0, 0,
                                                 width, height);
}

GdkPixbuf *copy_area_simple(const GdkPixbuf *src_pixbuf, int src_x, int src_y, int width, int height) {
        GdkPixbuf *dest_pixbuf = gdk_pixbuf_new(GDK_COLORSPACE_RGB, gdk_pixbuf_get_has_alpha(src_pixbuf),
                                                8, width, height);
        gdk_pixbuf_copy_area(src_pixbuf, src_x, src_y, width, height,
                             dest_pixbuf, 0, 0);
        return dest_pixbuf;
}

GdkPixbuf *screenshot() {
        GdkWindow *root;
        GdkPixbuf *pixbuf;
        root = gdk_get_default_root_window();
        int width = gdk_window_get_width(root);
        int height = gdk_window_get_height(root);
        pixbuf = gdk_pixbuf_get_from_window(root, 0, 0, width, height);
        return pixbuf;
}
