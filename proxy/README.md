网络代理相关库, 监听并将 gsettings(com.deepin.wrap.gnome.system.proxy)
的值同步给当前进程的环境变量, 用到的环境变量包括:

```
http_proxy="http://user:pass@127.0.0.1:8080/"
https_proxy="https://127.0.0.1:8080/"
ftp_proxy="ftp://127.0.0.1:8080/"
all_proxy="http://127.0.0.1:8080/"
SOCKS_SERVER=socks5://127.0.0.1:8000/
no_proxy="localhost,127.0.0.0/8,::1"
```

因为 Linux 没有统一的接口处理系统代理, Deepin 在兼容 GNOME 系统代理的
基础上同时会设置环境变量, 以求适配更多的网络应用, 由于环境变量是进程内
设置的, 所以抽象出这个库用于给 Deepin 相关程序(startdde/launcher/dock)
动态更新系统代理环境变量, 使用方法很简单:

```go
import (
	"github.com/linuxdeepin/go-gir/glib-2.0"
	"github.com/linuxdeepin/go-lib/proxy"
)

func main() {
	SetupProxy()
	glib.StartLoop()
}
```
