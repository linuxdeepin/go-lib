// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gdkpixbuf

type GdkInterpType uint

const (
	GDK_INTERP_NEAREST GdkInterpType = iota
	GDK_INTERP_TILES
	GDK_INTERP_BILINEAR
	GDK_INTERP_HYPER
)

type GdkPixbufRotation uint

const (
	GDK_PIXBUF_ROTATE_NONE             GdkPixbufRotation = 0
	GDK_PIXBUF_ROTATE_COUNTERCLOCKWISE GdkPixbufRotation = 90
	GDK_PIXBUF_ROTATE_UPSIDEDOWN       GdkPixbufRotation = 180
	GDK_PIXBUF_ROTATE_CLOCKWISE        GdkPixbufRotation = 270
)

type GdkPixbufAlphaMode uint

const (
	GDK_PIXBUF_ALPHA_BILEVEL GdkPixbufAlphaMode = iota
	GDK_PIXBUF_ALPHA_FULL
)

type GdkColorspace uint

const (
	GDK_COLORSPACE_RGB GdkColorspace = iota
)

type GdkPixbufError uint

const (
	/* image data hosed */
	GDK_PIXBUF_ERROR_CORRUPT_IMAGE GdkPixbufError = iota
	/* no mem to load image */
	GDK_PIXBUF_ERROR_INSUFFICIENT_MEMORY
	/* bad option passed to save routine */
	GDK_PIXBUF_ERROR_BAD_OPTION
	/* unsupported image type (sort of an ENOSYS) */
	GDK_PIXBUF_ERROR_UNKNOWN_TYPE
	/* unsupported operation (load, save) for image type */
	GDK_PIXBUF_ERROR_UNSUPPORTED_OPERATION
	GDK_PIXBUF_ERROR_FAILED
)

type GdkPixdataType uint

const (
	/* colorspace + alpha */
	GDK_PIXDATA_COLOR_TYPE_RGB  GdkPixdataType = 0x01
	GDK_PIXDATA_COLOR_TYPE_RGBA GdkPixdataType = 0x02
	GDK_PIXDATA_COLOR_TYPE_MASK GdkPixdataType = 0xff
	/* width, support 8bits only currently */
	GDK_PIXDATA_SAMPLE_WIDTH_8    GdkPixdataType = 0x01 << 16
	GDK_PIXDATA_SAMPLE_WIDTH_MASK GdkPixdataType = 0x0f << 16
	/* encoding */
	GDK_PIXDATA_ENCODING_RAW  GdkPixdataType = 0x01 << 24
	GDK_PIXDATA_ENCODING_RLE  GdkPixdataType = 0x02 << 24
	GDK_PIXDATA_ENCODING_MASK GdkPixdataType = 0x0f << 24
)

type GdkPixdataDumpType uint

const (
	/* type of source to save */
	GDK_PIXDATA_DUMP_PIXDATA_STREAM GdkPixdataDumpType = 0
	GDK_PIXDATA_DUMP_PIXDATA_STRUCT GdkPixdataDumpType = 1
	GDK_PIXDATA_DUMP_MACROS         GdkPixdataDumpType = 2
	/* type of variables to use */
	GDK_PIXDATA_DUMP_GTYPES GdkPixdataDumpType = 0
	GDK_PIXDATA_DUMP_CTYPES GdkPixdataDumpType = 1 << 8
	GDK_PIXDATA_DUMP_STATIC GdkPixdataDumpType = 1 << 9
	GDK_PIXDATA_DUMP_CONST  GdkPixdataDumpType = 1 << 10
	/* save RLE decoder macro? */
	GDK_PIXDATA_DUMP_RLE_DECODER GdkPixdataDumpType = 1 << 16
)
