// SPDX-FileCopyrightText: 2026 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package allowcaller

import (
	"bytes"
	"io/ioutil"
	"os"
	"reflect"
	"syscall"
	"testing"
)

func TestParseLoaderArgs(t *testing.T) {
	info, args, err := parseLoaderArgs([]string{
		"dde-session-daemon", "--verbose", "--fd1", "7", "--fd2", "8", "--force",
	})
	if err != nil {
		t.Fatal(err)
	}
	if !info.loaded || info.fd1 != 7 || info.fd2 != 8 {
		t.Fatalf("unexpected loader info: %+v", info)
	}
	want := []string{"dde-session-daemon", "--verbose", "--force"}
	if !reflect.DeepEqual(args, want) {
		t.Fatalf("unexpected cleaned args: got %v, want %v", args, want)
	}
}

func TestParseLoaderArgsRequiresBothFDs(t *testing.T) {
	if _, _, err := parseLoaderArgs([]string{"startdde", "--fd1", "7"}); err == nil {
		t.Fatal("expected incomplete loader arguments to fail")
	}
}

func TestParseLoaderArgsRejectsDuplicateFDs(t *testing.T) {
	tests := [][]string{
		{"dde-session-daemon", "--fd1", "7", "--fd1", "9", "--fd2", "8"},
		{"dde-session-daemon", "--fd1", "7", "--fd2", "8", "--fd2", "9"},
	}
	for _, args := range tests {
		if _, _, err := parseLoaderArgs(args); err == nil {
			t.Fatalf("expected duplicate loader arguments to fail: %v", args)
		}
	}
}

func TestParseLoaderArgsRequiresDistinctFDs(t *testing.T) {
	if _, _, err := parseLoaderArgs([]string{
		"dde-session-daemon", "--fd1", "7", "--fd2", "7",
	}); err == nil {
		t.Fatal("expected identical loader file descriptors to fail")
	}
}

func TestParseLoaderArgsPassesThroughUnrelatedArgs(t *testing.T) {
	info, args, err := parseLoaderArgs([]string{"dde-session-daemon", "--verbose"})
	if err != nil {
		t.Fatal(err)
	}
	if info.loaded {
		t.Fatal("unrelated args were reported as a security-loader invocation")
	}
	want := []string{"dde-session-daemon", "--verbose"}
	if !reflect.DeepEqual(args, want) {
		t.Fatalf("unexpected cleaned args: got %v, want %v", args, want)
	}
}

func TestReadResponseLimit(t *testing.T) {
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		_, _ = writer.Write(bytes.Repeat([]byte{'a'}, int(maxLoaderResponseSize)))
		_ = writer.Close()
	}()
	data, err := readResponse(reader)
	if err != nil {
		t.Fatal(err)
	}
	_ = reader.Close()
	if len(data) != int(maxLoaderResponseSize) {
		t.Fatalf("unexpected response length: got %d, want %d", len(data), maxLoaderResponseSize)
	}

	reader, writer, err = os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		_, _ = writer.Write(bytes.Repeat([]byte{'a'}, int(maxLoaderResponseSize)+1))
		_ = writer.Close()
	}()
	if _, err := readResponse(reader); err == nil {
		_ = reader.Close()
		t.Fatal("expected oversized loader response to fail")
	}
	_ = reader.Close()
}

func TestValidateLoaderFD(t *testing.T) {
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()
	defer writer.Close()

	if err := validateLoaderFD(writer, syscall.O_WRONLY); err != nil {
		t.Fatalf("expected pipe writer to be valid: %v", err)
	}
	if err := validateLoaderFD(reader, syscall.O_RDONLY); err != nil {
		t.Fatalf("expected pipe reader to be valid: %v", err)
	}
	if err := validateLoaderFD(reader, syscall.O_WRONLY); err == nil {
		t.Fatal("expected pipe access mode mismatch to fail")
	}

	regularFile, err := ioutil.TempFile(t.TempDir(), "security-loader")
	if err != nil {
		t.Fatal(err)
	}
	defer regularFile.Close()
	if err := validateLoaderFD(regularFile, syscall.O_RDWR); err == nil {
		t.Fatal("expected regular file descriptor to fail")
	}
}

func TestHandshakeReportsLoaderState(t *testing.T) {
	args, loaded, err := Handshake([]string{"dde-session-daemon", "--verbose"}, []Destination{{}})
	if err != nil {
		t.Fatal(err)
	}
	if loaded {
		t.Fatal("direct invocation was reported as security-loader invocation")
	}
	want := []string{"dde-session-daemon", "--verbose"}
	if !reflect.DeepEqual(args, want) {
		t.Fatalf("unexpected args: got %v, want %v", args, want)
	}

	_, loaded, err = Handshake([]string{"dde-session-daemon", "--fd1", "7"}, []Destination{{}})
	if err == nil {
		t.Fatal("expected incomplete loader arguments to fail")
	}
	if !loaded {
		t.Fatal("invalid injected arguments did not report security-loader invocation")
	}
}
