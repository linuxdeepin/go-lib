// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package graphic

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"io/ioutil"
	"strings"
)

// ConvertImageToDataUri convert image file to data uri.
func ConvertImageToDataUri(imgfile string) (dataUri string, err error) {
	data, err := ioutil.ReadFile(imgfile)
	if err != nil {
		return
	}
	f, err := GetImageFormat(imgfile)
	if err != nil {
		return
	}
	contentType := imageFormatToDataUriContentType(f)
	dataUri = fmt.Sprintf("data:%s;base64,%s", contentType, base64.StdEncoding.EncodeToString(data))
	return
}

// ConvertImageToDataUri convert image.Image object to data uri.
func ConvertImageObjectToDataUri(img image.Image, f Format) (dataUri string, err error) {
	byteBuf := new(bytes.Buffer)
	err = doSaveImage(byteBuf, img, f)
	if err != nil {
		return
	}
	data := byteBuf.Bytes()
	contentType := imageFormatToDataUriContentType(f)
	dataUri = fmt.Sprintf("data:%s;base64,%s", contentType, base64.StdEncoding.EncodeToString(data))
	byteBuf.Reset()
	return
}

// ConvertDataUriToImage convert data uri to image file.
func ConvertDataUriToImage(dataUri string, dstfile string, f Format) (err error) {
	img, err := LoadImageFromDataUri(dataUri)
	if err != nil {
		return
	}
	return SaveImage(dstfile, img, f)
}

// LoadImageFromDataUri convert data uri to image object.
func LoadImageFromDataUri(dataUri string) (img image.Image, err error) {
	strs := strings.Split(dataUri, ",") // ;base64,
	if len(strs) != 2 {
		err = fmt.Errorf("invalid data uri: %s", dataUri)
		return
	}
	data, err := base64.StdEncoding.DecodeString(strs[1])
	if err != nil {
		return
	}
	r := bytes.NewReader(data)
	img, _, err = image.Decode(r)
	return
}

func imageFormatToDataUriContentType(f Format) (contentType string) {
	contentType = fmt.Sprintf("image/%s", f)
	return
}
