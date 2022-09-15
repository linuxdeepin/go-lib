// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package mime

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueryURI(t *testing.T) {
	var infos = []struct {
		uri  string
		mime string
	}{
		{
			uri:  "testdata/data.txt",
			mime: "text/plain",
		},
		{
			uri:  "testdata/Deepin/index.theme",
			mime: MimeTypeGtk,
		},
	}

	for _, info := range infos {
		m, err := Query(info.uri)
		require.NoError(t, err)
		assert.Equal(t, m, info.mime)
	}
}
