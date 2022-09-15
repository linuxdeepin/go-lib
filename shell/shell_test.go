// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package shell

import (
	"os/exec"
	"testing"
)

func TestEncode(t *testing.T) {
	_, err := exec.LookPath("sh")
	if err != nil {
		t.Skip("not found sh")
	}

	for _, s := range []string{
		"hello world",
		"hello$world",
		"hello\t\r\nworld",
		"中文 english",
		"`~!#$&*()|\\;'\"<>? ",
	} {
		r := Encode(s)
		t.Log(r)

		cmd := exec.Command("sh", "-c", "echo -n "+r)
		output, err := cmd.Output()
		if err != nil {
			t.Fatal(err)
		}
		if s != string(output) {
			t.Errorf("%q != %q", s, string(output))
		}
	}
}
