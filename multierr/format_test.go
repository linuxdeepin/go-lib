// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package multierr

import (
	"errors"
	"testing"
)

func TestListFormatFuncSingle(t *testing.T) {
	expected := `1 error occurred:
	* foo`

	errs := []error{
		errors.New("foo"),
	}

	actual := ListFormatFunc(errs)
	if actual != expected {
		t.Fatalf("bad: %#v", actual)
	}
}

func TestListFormatFuncMultiple(t *testing.T) {
	expected := `2 errors occurred:
	* foo
	* bar`

	errs := []error{
		errors.New("foo"),
		errors.New("bar"),
	}

	actual := ListFormatFunc(errs)
	if actual != expected {
		t.Fatalf("bad: %#v", actual)
	}
}
