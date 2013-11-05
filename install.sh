#!/bin/bash
# don't direct use this
PWD1=`pwd`
mkdir src && cd src && ln -sv ../dlib .
echo ${PWD1}
GOPATH=${PWD1} go install dlib/pseudo
GOPATH=${PWD1} go install dlib/dbus/proxyer
GOLIB=`go env GOHOSTOS`_`go env GOHOSTARCH`
rm ${PWD1}/pkg/${GOLIB}/dlib/pseudo.a
