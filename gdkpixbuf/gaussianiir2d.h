/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
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
