## Deepin GoLang Library

Deepin GoLang 库是一个包含许多有用的 go 例程的库，用于 glib、gettext、存档、图形等。

## 依赖


### 编译依赖

* gio-2.0
* glib-2.0
* x11
* gdk-pixbuf-2.0
* libpulse
* mobile-broadband-provider-info

## 安装

go-lib需要预安装以下包

```
$ go get github.com/smartystreets/goconvey
$ go get github.com/howeyc/fsnotify
$ go get gopkg.in/check.v1
$ go get github.com/linuxdeepin/go-x11-client
```

安装

```
mkdir -p $GOPATH/src/github.com/linuxdeepin/
cp -r go-lib $GOPATH/src/github.com/linuxdeepin/go-lib
```

## 获得帮助

如果您遇到任何其他问题，您可能还会发现这些渠道很有用：

* [Gitter](https://gitter.im/orgs/linuxdeepin/rooms)
* [IRC channel](https://webchat.freenode.net/?channels=deepin)
* [Forum](https://bbs.deepin.org/)
* [WiKi](http://wiki.deepin.org/)

## 贡献指南

我们鼓励您报告问题并做出更改。

* [Contribution guide for users](http://wiki.deepin.org/index.php?title=Contribution_Guidelines_for_Users)
* [Contribution guide for developers](http://wiki.deepin.org/index.php?title=Contribution_Guidelines_for_Developers)

## 开源协议

go-lib项目在 [GPL-3.0-or-later](LICENSE)开源协议下发布。
