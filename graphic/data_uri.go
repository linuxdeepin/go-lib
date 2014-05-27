/**
 * Copyright (c) 2013 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 Xu FaSheng
 *
 * Author:      Xu FaSheng <fasheng.xu@gmail.com>
 * Maintainer:  Xu FaSheng <fasheng.xu@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package graphic

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

// ConvertImageToDataUri convert image file to data uri.
func ConvertImageToDataUri(imgfile string) (dataUri string, err error) {
	data, err := ioutil.ReadFile(imgfile)
	if err != nil {
		return
	}
	format, err := getImageFormat(imgfile)
	if err != nil {
		return
	}
	contentType := imageFormatToDataUriContentType(format)
	dataUri = fmt.Sprintf("data:%s;base64,%s", contentType, base64.StdEncoding.EncodeToString(data))
	return
}

func imageFormatToDataUriContentType(f Format) (contentType string) {
	contentType = fmt.Sprintf("image/%s", f)
	return
}
