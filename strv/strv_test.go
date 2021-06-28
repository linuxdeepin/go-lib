/*
 * Copyright (C) 2017 ~ 2018 Deepin Technology Co., Ltd.
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

package strv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	vector := Strv([]string{"a", "b", "c"})
	assert.True(t, vector.Contains("b"))
	assert.False(t, vector.Contains("d"))
}

func TestEqual(t *testing.T) {
	v1 := Strv([]string{"a", "b", "c"})
	v2 := Strv([]string{"a", "b", "c", "d"})
	v3 := Strv(v1[:])
	assert.False(t, v1.Equal(v2))
	assert.True(t, v1.Equal(v3))
}

func TestUniq(t *testing.T) {
	vector := Strv([]string{"a", "b", "c", "c", "b", "a", "c"})
	vector = vector.Uniq()
	assert.Equal(t, vector, Strv([]string{"a", "b", "c"}))
}

func TestFilterFunc(t *testing.T) {
	vector := Strv([]string{"hello", "", "world", "", "!"})
	vector = vector.FilterFunc(func(str string) bool {
		return len(str) == 0
	})
	assert.Equal(t, vector, Strv([]string{"hello", "world", "!"}))
}

func TestFilterEmpty(t *testing.T) {
	vector := Strv([]string{"hello", "", "world", "", "!"})
	vector = vector.FilterEmpty()
	assert.Equal(t, vector, Strv([]string{"hello", "world", "!"}))
}

func TestAdd(t *testing.T) {
	vector := Strv([]string{"a", "b", "c"})

	vector0, b0 := vector.Add("d")
	assert.Equal(t, vector, Strv([]string{"a", "b", "c"}))
	assert.Equal(t, vector0, Strv([]string{"a", "b", "c", "d"}))
	assert.True(t, b0)

	vector1, b1 := vector.Add("c")
	assert.Equal(t, vector, Strv([]string{"a", "b", "c"}))
	assert.Equal(t, vector1, Strv([]string{"a", "b", "c"}))
	assert.False(t, b1)
}

func TestDelete(t *testing.T) {
	vector := Strv([]string{"a", "b", "c"})
	vector0, b0 := vector.Delete("d")
	assert.Equal(t, vector, Strv([]string{"a", "b", "c"}))
	assert.Equal(t, vector0, Strv([]string{"a", "b", "c"}))
	assert.False(t, b0)

	vector1, b1 := vector.Delete("c")
	assert.Equal(t, vector, Strv([]string{"a", "b", "c"}))
	assert.Equal(t, vector1, Strv([]string{"a", "b"}))
	assert.True(t, b1)
}
