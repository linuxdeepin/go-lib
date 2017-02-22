/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package graphic

// Format defines the type of image format.
type Format string

// Registered image format.
const (
	FormatPng  Format = "png"
	FormatJpeg Format = "jpeg"
	FormatBmp  Format = "bmp"
	FormatTiff Format = "tiff"
)
