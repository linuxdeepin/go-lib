/**
 * Copyright (c) 2011 ~ 2013 Deepin, Inc.
 *               2011 ~ 2013 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
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

package utils

const (
        URI_STRING_FILE  = "file://"
        URI_STRING_FTP   = "ftp://"
        URI_STRING_HTTP  = "http://"
        URI_STRING_HTTPS = "https://"
        URI_STRING_SMB   = "smb://"
)

func (op *Manager) URIToPath(uri string) (string, bool) {
        tmp := deleteStartSpace(uri)

        if op.IsContainFromStart(tmp, URI_STRING_FILE) {
                return tmp[7:], true
        } else if op.IsContainFromStart(tmp, URI_STRING_FTP) {
                return tmp[6:], true
        } else if op.IsContainFromStart(tmp, URI_STRING_HTTP) {
                return tmp[7:], true
        } else if op.IsContainFromStart(tmp, URI_STRING_HTTPS) {
                return tmp[8:], true
        } else if op.IsContainFromStart(tmp, URI_STRING_SMB) {
                return tmp[6:], true
        } else if op.IsContainFromStart(tmp, "/") {
                return tmp, true
        }

        return "", false
}

func (op *Manager) PathToFileURI(path string) (string, bool) {
        tmp := deleteStartSpace(path)

        if op.IsContainFromStart(tmp, "/") {
                return URI_STRING_FILE + path, true
        } else if op.IsContainFromStart(tmp, URI_STRING_FILE) {
                return tmp, true
        }

        return "", false
}

func (op *Manager) PathToFtpURI(path string) (string, bool) {
        tmp := deleteStartSpace(path)

        if op.IsContainFromStart(tmp, "/") {
                return URI_STRING_FTP + path, true
        } else if op.IsContainFromStart(tmp, URI_STRING_FTP) {
                return tmp, true
        }

        return "", false
}

func (op *Manager) PathToHttpURI(path string) (string, bool) {
        tmp := deleteStartSpace(path)

        if op.IsContainFromStart(tmp, "/") {
                return URI_STRING_HTTP + path, true
        } else if op.IsContainFromStart(tmp, URI_STRING_HTTP) {
                return tmp, true
        }

        return "", false
}

func (op *Manager) PathToHttpsURI(path string) (string, bool) {
        tmp := deleteStartSpace(path)

        if op.IsContainFromStart(tmp, "/") {
                return URI_STRING_HTTPS + path, true
        } else if op.IsContainFromStart(tmp, URI_STRING_HTTPS) {
                return tmp, true
        }

        return "", false
}

func (op *Manager) PathToSmbURI(path string) (string, bool) {
        tmp := deleteStartSpace(path)

        if op.IsContainFromStart(tmp, "/") {
                return URI_STRING_SMB + path, true
        } else if op.IsContainFromStart(tmp, URI_STRING_SMB) {
                return tmp, true
        }

        return "", false
}
