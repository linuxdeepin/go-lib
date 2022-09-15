// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTarWriterCompressFiles(t *testing.T) {
	var testFiles []string
	err := TarWriterCompressFiles(nil, testFiles)
	assert.NotEqual(t, nil, err)
}

func TestTarReaderExtracte(t *testing.T) {
	_, err := TarReaderExtracte(nil, "")
	assert.NotEqual(t, nil, err)
}

func TestTarWriterCompressDir(t *testing.T) {
	err := tarWriterCompressDir(nil, "", "")
	assert.NotEqual(t, nil, err)
}

func TestTarWriterCompressFile(t *testing.T) {
	err := tarWriterCompressFile(nil, "", "")
	assert.NotEqual(t, nil, err)
}
