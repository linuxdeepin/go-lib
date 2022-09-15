// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package kv

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReader(t *testing.T) {
	f, err := os.Open("./testdata/a")
	require.NoError(t, err)
	assert.NotNil(t, f)
	defer f.Close()

	r := NewReader(f)
	r.Delim = ':'
	r.TrimSpace = TrimDelimRightSpace | TrimTailingSpace

	var resultPair *Pair
	for {
		pair, err := r.Read()
		if err != nil {
			break
		}
		if pair.Key == "Pid" {
			resultPair = pair
			break
		}
	}
	assert.NotNil(t, resultPair)
	assert.Equal(t, resultPair.Value, "21722")

	pairs, err := r.ReadAll()
	require.NoError(t, err)
	assert.Equal(t, len(pairs), 43)

	f, err = os.Open("./testdata/b")
	require.NoError(t, err)
	assert.NotNil(t, f)
	defer f.Close()

	r = NewReader(f)
	pairs, err = r.ReadAll()
	require.Nil(t, pairs)
	assert.Equal(t, err, ErrBadLine)

	f, err = os.Open("./testdata/c")
	require.NoError(t, err)
	assert.NotNil(t, f)
	defer f.Close()

	r = NewReader(f)
	r.TrimSpace = TrimLeadingTailingSpace
	r.Comment = '#'

	pair, err := r.Read()
	require.NoError(t, err)
	assert.Equal(t, pair, &Pair{"LANG", "zh_CN.UTF-8"})

	pair, err = r.Read()
	require.NoError(t, err)
	assert.Equal(t, pair, &Pair{"LANGUAGE", "zh_CN"})

	pair, err = r.Read()
	require.Nil(t, pair)
	assert.Equal(t, err, io.EOF)
}
