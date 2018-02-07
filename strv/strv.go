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

// string vector

type Strv []string

func (vector Strv) Contains(str string) bool {
	for _, e := range vector {
		if e == str {
			return true
		}
	}
	return false
}

func (v1 Strv) Equal(v2 Strv) bool {
	if len(v1) != len(v2) {
		return false
	}
	for i, e1 := range v1 {
		if e1 != v2[i] {
			return false
		}
	}
	return true
}

func (vector Strv) Uniq() Strv {
	var newVector Strv
	for _, e := range vector {
		if !newVector.Contains(e) {
			newVector = append(newVector, e)
		}
	}
	return newVector
}

func (vector Strv) FilterFunc(fn func(string) bool) Strv {
	if fn == nil {
		return vector
	}

	newVector := make(Strv, 0, len(vector))
	for _, e := range vector {
		if fn(e) {
			continue
		}
		newVector = append(newVector, e)
	}
	return newVector
}

func (vector Strv) FilterEmpty() Strv {
	newVector := make(Strv, 0, len(vector))
	for _, e := range vector {
		if len(e) == 0 {
			continue
		}
		newVector = append(newVector, e)
	}
	return newVector
}

func (vector Strv) Add(str string) (Strv, bool) {
	if vector.Contains(str) {
		return vector, false
	}
	return Strv(append(vector, str)), true
}

func (vector Strv) Delete(str string) (Strv, bool) {
	var found bool
	ret := make(Strv, 0, len(vector))

	for _, e := range vector {
		if e == str {
			found = true
			continue
		}
		ret = append(ret, e)
	}

	return ret, found
}
