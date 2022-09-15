// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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
