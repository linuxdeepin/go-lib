/*
 * Copyright (C) 2016 ~ 2018 Deepin Technology Co., Ltd.
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
	require.Nil(t, err)
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
	require.Nil(t, err)
	assert.Equal(t, len(pairs), 43)

	f, err = os.Open("./testdata/b")
	require.Nil(t, err)
	assert.NotNil(t, f)
	defer f.Close()

	r = NewReader(f)
	pairs, err = r.ReadAll()
	require.Nil(t, pairs)
	assert.Equal(t, err, ErrBadLine)

	f, err = os.Open("./testdata/c")
	require.Nil(t, err)
	assert.NotNil(t, f)
	defer f.Close()

	r = NewReader(f)
	r.TrimSpace = TrimLeadingTailingSpace
	r.Comment = '#'

	pair, err := r.Read()
	require.Nil(t, err)
	assert.Equal(t, pair, &Pair{"LANG", "zh_CN.UTF-8"})

	pair, err = r.Read()
	require.Nil(t, err)
	assert.Equal(t, pair, &Pair{"LANGUAGE", "zh_CN"})

	pair, err = r.Read()
	require.Nil(t, pair)
	assert.Equal(t, err, io.EOF)
}
