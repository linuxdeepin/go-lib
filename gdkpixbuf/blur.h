/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

#ifndef _BLUR_H
#define _BLUR_H 1

#include <gdk-pixbuf/gdk-pixbuf.h>

int blur(GdkPixbuf *src_pixbuf, double sigma, double numsteps);

#endif /* _BLUR_H */
