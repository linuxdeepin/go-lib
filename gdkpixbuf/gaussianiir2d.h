/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

#ifndef _GAUSSIANIIR2D_H_
#define _GAUSSIANIIR2D_H_

#include <glib.h>

void gaussianiir2d_f(double* image_f,
		     long width, long height,
		     double sigma, long numsteps);
/*
 *	image data format
 *
 *	1. _pixbuf_c: use GdkPixbuf format.
 *	   p = pixels + y * rowstride + x* n_channels
 *
 *	2. gaussianiir2d_c: use cairo image data
 */
void gaussianiir2d_pixbuf_c(unsigned char* image_data,
			    int width, int height,
			    int rowstride, int n_channels,
			    double sigma, double numsteps);
#if 0
void gaussianiir2d_c(unsigned char* image_c,
		     long width, long height,
		     double sigma, long numsteps);

#endif
#endif /* _GAUSSIANIIR2D_H_ */
