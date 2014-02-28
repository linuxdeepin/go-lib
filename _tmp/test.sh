#!/usr/bin/env bash

GOPATH=$(pwd) go install -x xdlib/pseudo
cp -vf ./pkg/linux_amd64/xdlib/* /usr/lib/go/pkg/linux_amd64/dlib/
cp -rvf ./src/xdlib/* /usr/lib/go/src/pkg/dlib/
