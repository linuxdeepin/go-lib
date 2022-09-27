## Deepin GoLang Library

Deepin GoLang Library is a library containing many useful go routines for things such as glib, gettext, archive, graphic,etc.

## Dependencies


### Build dependencies

* gio-2.0
* glib-2.0
* x11
* gdk-pixbuf-2.0
* libpulse
* mobile-broadband-provider-info

## Installation

Install prerequisites

```
$ go get github.com/smartystreets/goconvey
$ go get github.com/howeyc/fsnotify
$ go get gopkg.in/check.v1
$ go get github.com/linuxdeepin/go-x11-client
```

Install

```
mkdir -p $GOPATH/src/github.com/linuxdeepin/
cp -r go-lib $GOPATH/src/github.com/linuxdeepin/go-lib
```

## Getting help

Any usage issues can ask for help via

* [Gitter](https://gitter.im/orgs/linuxdeepin/rooms)
* [IRC channel](https://webchat.freenode.net/?channels=deepin)
* [Forum](https://bbs.deepin.org/)
* [WiKi](http://wiki.deepin.org/)

## Getting involved

We encourage you to report issues and contribute changes.

* [Contribution guide for users](http://wiki.deepin.org/index.php?title=Contribution_Guidelines_for_Users)
* [Contribution guide for developers](http://wiki.deepin.org/index.php?title=Contribution_Guidelines_for_Developers)

## License

Deepin GoLang Library is licensed under [GPL-3.0-or-later](LICENSE).
