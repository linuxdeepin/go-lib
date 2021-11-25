## 接口简介

- **NewLogger**

  创建 Logger 对象, 需要传递一个日志名称, 如 "daemon/network"

- **Logger.SetLogLevel**

  设置日志记录级别, 默认为 Logger.LevelInfo

- **Logger.Debug**, **Logger.Info**, **Logger.Warning**, **Logger.Error**

  日志记录接口, 其中 Logger.Error() 会额外打印 go 函数调用轨迹 (trace)

- **Logger.Debugf**, **Logger.Infof**, **Logger.Warningf**, **Logger.Errorf**

  日志记录接口, 支持 format 格式, 语法风格和 fmt.Printf() 相同

- **Logger.Panic**

  记录日志并额外执行一条 panic() 语句

- **Logger.Fatal**

  记录日志并额外执行 os.Exit(1)

## 环境变量

- **DDE_DEBUG**, 若值不为空, 则打印所有日志, 启用其他 **DDE_DEBUG_XXX** 变
  量时, 该变量会默认启用
- **DDE_DEBUG_MATCH**, 仅允许匹配该环境变量的 logger 对象打印i日志, 对
  logger name 进行匹配, 不区分大小写
- **DDE_DEBUG_LEVEL**, 用于从外部设置日志打印级别, 可选值为 "debug", "info", "warning", "error", "fatal"
- **DDE_DEBUG_CONSOLE**, 若值不为空, 则以 syslog 格式打印终端日志

示例 1, 打印 startdde 所有日志:
```
env DDE_DEBUG=1 /usr/bin/startdde
```

示例 2, 仅打印 startdde 警告级别以上的日志:
```
env DDE_DEBUG_LEVEL="warning" startddek
```

示例 3, 打印 dde-daemon 中网络模块的所有日志:
```
env DDE_DEBUG_MATCH="network" dde-session-daemon
```

示例 4, 打印 dde-daemon 中网络模块的日志, 且仅打印警告级别以上的日志:
```
env DDE_DEBUG_MATCH="network" DDE_DEBUG_LEVEL="warning" dde-session-daemon
```

## 示例代码

```go
import "github.com/linuxdeepin/go-lib/log"
import "flag"

var (
  l = log.NewLogger("daemon/test")
  argDebug bool
)

func main() {
  defer func() {
      if err := recover(); err != nil {
          l.Fatal(err)
      }
  }()

  // parse arguments
  flag.BoolVar(&argDebug, "d", false, "debug")
  flag.Parse()

  // setup logger
  if argDebug {
      l.SetLogLevel(l.LevelDebug)
  }
}
```

## 查看日志

对于 rsyslog, 可以配合辅助工具 logtool 高亮显示日志条目, 更加方便
```sh
# apt-get install logtool
$ tailf /var/log/syslog | logtool # follow syslog
$ tailf /var/log/syslog | grep "startdde"  | logtool # filter startdde syslog items
```

对于 systemd，则可以使用 journalctl
```sh
$ journalctl -f # follow syslog
$ journalctl /usr/bin/startdde # filter startdde syslog items
```
