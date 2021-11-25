# dbusutil

[包文档](https://godoc.org/github.com/linuxdeepin/go-lib/dbusutil)

## 一般用法

首先创建 Service。

用 NewSessionService 创建 session 服务； 用 NewSystemService 创建 system 服务。
```go
service := dbusutil.NewSessionServer()
```


然后导出对象
用 Export 。

```go
e := &Exportable1{}
service.Export(e)
```

e 要实现 dbusutil.Exportable 接口，并且必须类型必须是 struct 指针。

```
func (e *Exportable1) GetDBusExportInfo() dbusutil.ExportInfo {
    return dbusutil.ExportInfo{
        Path: "/the/object/path"
        Interface: "the.interface.Name"
    }
}
```

之后申请服务名，调用 service 的 RequestName 方法, 如果服务名已被其他连接申请了，就会返回错误。

```
service.RequestName("the.service.Name")
```

## 临时服务

服务在运行一段时间内没有接到其他调用请求，就去自动退出。

```
service.SetAutoQuitHanlder(1*time.Minute, func() bool {
    // 这个函数叫 canQuit 回调
    if canQuit {
        // 满足条件退出
        return true
    }

    // 不满足条件，不退出
    return false
})

// 等待退出信号
service.Wait()
```

由于没有修改 dbus 库，所有必须在每个导出方法的实现内加上调用 Service.DelayAutoQuit 的方法，如：

```
type Exportable1 struct{
    service *dbusutil.Service
}


func (e Exportable1) Method1() *dbus.Error {
    e.service.DelayAutoQuit()
    return nil
}
```

这样，每次有其他程序通过 dbus 调用了 Method1，就会延长本服务的存活时间。
目前的实现是 2 倍于 SetAutoQuitHandler 第一个参数的值的时间段内没有接到调用请求就会退出。

## 自省

dbusutil 包使用了 go 语言的反射机制自动生成导出对象的 introspection xml。

### 方法
由于反射机制不能获取方法的参数名，只能通过添加额外信息的方式提供方法的参数名，要求必须提供，否则就崩溃。

在结构体里面增加字段,名为 methods， 类型为结构体指针，具体字段立即定义，无需命名结构。
```
type Exportable1 {
    methods *struct{
        Method1 func() `in:"a,b" out:"result"`
        Method2 func() `in:"a,b"`
        Method3 func() `out:"result"`
    }
}
```

结构中每一个字段的名与导出的方法名对应，类型可为任意类型，但一般写成 func() 即函数指针，字段可以使用 in、out tag, in tag 内容为传入参数名列表，out tag 内容为返回参数名列表，参数名列表用逗号(,) 分割。 注意 in tag 和 out tag 之间要用空格分割。
如果导出方法没有输入和输出参数，则可以省略掉该字段。

### 信号
在结构图里面增加字段，名为 signals, 类型为结构体指针，具体字段立即定义，无需命名结构。
```
type Exportable1 struct {
    signals *struct{
        Signal1 struct{}

        Signal2 struct{
            name string
        }

        Signal3 struct{
            name string
            value uint32
        }
    }
}
```
这个匿名结构的每个字段，表示一个信号，类型必须为 struct, 这个更内层的 struct 内每个字段表示信号的参数。

为什么不用　func() 表示信号，是因为 go 的反射机制是不能获取方法的参数名，只能获取参数类型，虽然　d-feet 不能显示出信号的每个参数名是什么，但 D-Bus Introspection 规范是允许的。

### 属性

结构体内每个大写字母开头的字段只要同时满足如下所有条件都可以成为属性。

0. 不是 PropsMaster
1. 不是前一个字段的锁字段
2. 不带 prop tag 不为 "-"
3. 不带空结构，空结构在 dbus 中是非法的。

比如:
```
type Exportable1 struct{
    PropsMaster dbusutil.PropsMaster

    Prop1 string
    Prop1Mu sync.RWMutex

    Prop2 uint32

    Prop3 struct{}

    Prop4 string `prop:"-"`
}
```

最后导出的属性只有 Prop1 和 Prop2。（目前还有 Prop3,但是不能用）
* PropsMaster 是一个特殊的字段

* Prop1Mu 是 Prop1 的锁字段，它紧跟 Prop1 字段，并且名称是前一个字段名 Prop1 + Mu，并且类型是 sync.RWMutex。

* Prop3 的类型是空结构
* Prop4 使用了 prop tag 来明确忽略它。

属性字段可以使用 prop tag, 内容的格式是：

```
选项1:选项值1,选项2:选项值2
```

如 access:rw,emit:false

支持的选项有 access 和 emit。

access 选项值可以为 r, read, w, write, rw, readwrite 之一，默认为 read。
* r 或 read 表示 只读
* w 或 write 表示 只写
* rw 或 readwrite 表示 可读可写

emit 选项值可以为 true， false, invaliates 之一，默认为 true。
* true 表示属性被修改后，发送属性改变信号时带上改变的具体值；
* false 表示属性被修改后，不发送属性改变信号；
* invalidates 表示属性被修改后，发送属性改变信号是不带上具体的值；


#### 属性锁
 org.freedestkop.DBus.Properties interface 下的 Get、GetAll 和Set 方法如果有可用的锁，那么就都会自动加锁。Get 与 GetAll 加读锁, Set 加写锁。

目前有三种锁：

全局锁 PropsMaster dbusutil.PropsMaster, 用于值类型的字段，比如 uint32，uint64， string 类型，没有局部锁或内部锁的属性都会使用全局锁。

局部锁 PropXXXMu sync.RWMutex, 用于引用类型的字段，比如 map 类型；要求局部锁字段的名为前一个字段名加 Mu，类型为 sync.RWMutex，作为前一个字段的局部锁。

内部锁, 实现了 dbusutil.Property 接口类型的字段，比如
dbusutil/gsprop 包内的各种类型, 要求 gsprop 包内部实现加锁。

```
type Exportable1 struct {

  PropsMaster dbusutil.PropsMaster

  Prop1 uint32
  Prop2 uint32


  Prop3 map[string]string
  Prop3Mu sync.RWMutex


  Prop4 *gsprop.String
}
```

Prop1 与 Prop2 使用 PropsMaster 加锁

Prop3 使用 Prop3Mu 加锁

Prop4 使用 gsprop.String 内部的锁，在 dbusutil 代码里就是不加锁。


#### dbusutil-gen
利用 go 语言的 generate 特性自动生成属性的 set 与 get 方法，比如操作属性 Name 的 setPropName 与 getPropName 方法，get 方法内实现了加读锁， set 方法内部实现了加写锁与发送改变信号。

先安装 dbusutil-gen 到 $GOPATH/bin 中，执行命令：
```
go install github.com/linuxdeepin/go-lib/dbusutil/_tool/dbusutil-gen
```

在代码里写上特殊注释
```
//go:generate dbusutil-gen -type Type1,Type2 file1.go file2.go
```

Type1,Type2 是在 dbus 上导出的类型，用逗号分割。
file1.go file2.go 是要扫描的 go 源代码文件， 其中就定义了 Type1 和 Type2。

然后进入go源码文件所在目录，执行命令 `go generate`, 将看到 dbusutil-gen 的输出信息，自动生成文件 包名 + _dbusutil.go， 如包 pkg1 生成文件名为 pkg1_dbusutil.go, 还可以用 -output 选项指定。

具体选项查看 dbusutil-gen 的帮助信息

可导出属性的字段的注释第一行可以写上
```
// dbusutil-gen: XXXX
```
作为 dbusutil-gen 工具的指导指令

支持的指导指令有
* ignore 忽略此字段
* ignore-below 忽略之后的所有字段
* equal=EQUAL 用于比较新旧值

  EQUAL 可以为：
  * nil 不进行比较
  * 任意，如 bytes.Equal，作为函数，比较是使用 bytes.Equal(old, new)
  * 以method: 开头，如 method:equal, 作为方法，比较时使用 old.equal(new)
