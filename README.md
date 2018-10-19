# bingo-log

基于Go语言的日志处理包

## 简介

`bingo-log` 是为了 [Bingo]() 框架而制作的日志模块，我将其单独抽出来，使其可以单独使用而不依赖 `Bingo` 框架

`bingo-log` 不依赖官方日志包，完全依靠自己的逻辑实现

Features：

  - [x] 完整的单元测试，覆盖率超过90%
  - [x] 支持5种日志级别( `Fatal`,`Error`,`Warning`,`Debug`,`Info` )
  - [x] 同步写入日志
  - [x] 通过协程池异步写入日志
  - [x] 自定义日志命名格式
  - [x] 自定义日志文件分片存储
  - [x] 自定义输出日志或调试信息的抬头
  - [x] 可以自定义连接器，实现日志输出的定制化
  - [x] 内置常见的日志输出以及分片格式，开箱即用(目前实现两种连接器 `BaseConnector`和`KirinConnector`)

## 安装

```
  go get github.com/silsuer/bingo-log
```

## 使用

### 组成

日志结构体如下：

```go
 type Log struct {
	 Connector // 内嵌连接器，用来定制化功能
	 sync.Mutex
	 initialized     bool         // 该日志对象是否初始化
	 mode            int          // 日志记录模式  同步记录 或 协程池记录
	 pool            *grpool.Pool // 协程池
	 poolExpiredTime int          // 协程池模式下，每个空闲协程的存活时间(秒)
	 poolWorkerNum   int          // 协程池模式下，允许的最高协程数
 }
```

该结构体的 `mode` 属性目前支持两种模式: `LogSyncMode`（同步记录模式）和 `LogPoolMode`(异步协程池记录模式)

当使用 `LogSyncMode` 的时候，所有日志会同步打印在控制台和输出文件中

当使用 `LogPoolMode` 时，需指定协程池中的协程数量`poolWorkerNum`(默认是100) 和 每个空闲协程的存活时间`poolExpiredTime`(默认是50秒)

主要逻辑的实现都在日志结构体中内嵌的那个 `Connector` 接口中，所有实现了该接口的结构体都可以通过 `Log` 结构体的 `LoadConnector` 方法加载进来，实现自定义自己的日志格式

### 效果

1. 使用 `BaseConnector`:

   将在程序运行目录下创建一个名为 `bingo.log` 的日志文件，格式为：`[日志级别][打印时间][日志内容xxx]`

   例如:

   ```
      [FATAL] [2018-10-20 00:28:32] fatal testing
      [ERROR] [2018-10-20 00:28:32] Error testing
      [WARNING][2018-10-20 00:28:32] Warning testing
      [DEBUG] [2018-10-20 00:28:32] Debug testing
      [INFO][2018-10-20 00:28:32] Info testing
      [FATAL] [2018-10-20 00:28:32] Fatal testing
      [ERROR] [2018-10-20 00:28:32] Error testing
      [WARNING][2018-10-20 00:28:32] Warning testing
      [DEBUG] [2018-10-20 00:28:32] Debug testing
      [INFO][2018-10-20 00:28:32] Info testing
   ```

2. 使用 `KirinConnector`:

   将在程序运行目录下（或者是在你传入的配置map中指定的root目录下）,根据你传入的日期格式（配置map中的 format 键对应的值，默认是 `2006_01_02`）

   生成一个名为 bingo_年_月_日.log 的日志文件

   输出的日志较详细，例如:

   ```
      [DEBUG]
      [2018-10-20 00:28:32]
      [FILE] /Users/silsuer/go/src/github.com/silsuer/bingo-log/kirin_test.go
      [LINE] 15
      [CONTENT] KirinConnector Debug testing
   ```

### 基本使用过程

1. 创建日志对象并指定写入模式

   ```go

      log := bingo_log.NewLog(bingo_log.LogSyncMode)  // 同步写入模式
      log := bingo_log.NewLog(bingo_log.LogPoolMode)  // 异步协程池模式

   ```

2. 如果指定模式为`LogPoolMode`，需要设置协程过期时间和协程池中的协程数量（当然，不设置使用默认值也可以）

  ```go
       log.SetPoolExpiredTime(100)  // 设置协程过期时间
       log.SetPoolWorkerNum(100)    // 设置协程池中协程数量
  ```

3. 创建日志连接器并做好初始化工作

  ```
    // 对于基础的连接器不需要做其他工作
    conn := new(bingo_log.BaseConnector)

    // 对于 KirinConnector连接器需要传入自定义的配置map
    config := make(map[string]string)
    config["root"] = "/set/any/path"
    config["format"] = "2006_01_02"
    conn := bingo_log.NewKirinConnector(config)

    // 对于你自己实现的连接器emmm你自己开心就好....
  ```

4. 将连接器载入日志对象内

  ```
     log.LoadConnector(conn)
  ```

5. 可以愉快的打印了...

  ```
    log.Info("testing")
    log.Debug("testing")
    log.Warning("testing")
    log.Error("testing")
    log.Fatal("testing")
  ```


### 进阶: 自定义自己的日志格式

此时我们就需要构建一个属于自己的连接器了，连接器接口 `bingo_log.Connector` 如下：

```
  type Connector interface {
  	Fatal(message ...interface{})
  	Error(message ...interface{})
  	Warning(message ...interface{})
  	Debug(message ...interface{})
  	Info(message ...interface{})                           // 打印
  	Output(message string)                                 // 将信息输出到文件中
  	GetMessage(degree int, message ...interface{}) string // 将输入的信息添加抬头（例如添加打印时间等）
  	GetFile(config map[string]string) *os.File             // 当前日志要输出到的文件位置,传入一个map 代表配置
  }
```

所以我就需要令一个结构体实现该接口的所有方法

例如，我想打印日志的时候，定义输出格式为 "[出错了...] xxxx",并输出在 "/Users/silsuer/go/error.log" 文件中

1. 那么新建一个结构体：

 ```go
   type ExampleConnector struct {

   }
 ```

2. 先实现`GetMessage`方法,所有的打印方法（`Fatal`、`Error`等）都会先调用这个方法将传入的参数进行包装

  ```
   func (e ExampleConnector) GetMessage(degree int, message ...interface{}) string {
	   return "[出错了...]" + fmt.Sprint(message...)
   }
  ```

3. 实现`GetFile` 方法，返回要写入到的文件指针

   ```
     func (e ExampleConnector) GetFile(config map[string]string) *os.File {
	   return bingo_log.OpenOrCreateFile("/Users/silsuer/go/error.log")
     }
   ```

4. 接下来就是实现其他接口中的其他方法了，基本上照着`BaseConnector`搬下来就可以:

   其中 `Fatal`,`Error`,`Warning`,`Debug`,`Info` 5种方法负责控制输出，在基础连接器中，就是根据不同的日志级别，包装成不同的日志信息，然后

   打印到控制台一份，输出到文件中一份，所以此处你可以对这5种方法进行不同的定制（如果你开心的话），这里我只写一个 `Fatal` 方法的实现

   ```go
    func (e ExampleConnector) Fatal(message ...interface{}) {
   	  m := e.GetMessage(bingo_log.FATAL, message...)
   	  fmt.Print(m)
   	  e.Output(m)
    }
   ```

   再实现 `Output` 方法，这个方法就是上面那5种方法中调用的要输出到文件中的方法：

   ```go
     func (e ExampleConnector) Output(message string) {
   	  // 获取到要输出的文件路径
   	  file := e.GetFile(make(map[string]string))
   	  defer file.Close()
   	  n, _ := file.Seek(0, os.SEEK_END)
   	  // 写入数据
   	  file.WriteAt([]byte(message), n)
   }
   ```

5. 这样我们就实现了一个完整的连接器了，当然你可以通过结构体组合的形式，减少代码的书写，可以使用了

```
   func main(){
     log := bingo_log.NewLog(bingo_log.LogSyncMode)  // 同步写入模式
     e := new(ExampleConnector)  // 新建刚刚写好的连接器
     log.LoadConnector(e)        // 加载连接器

     log.Info("成功了...")
     log.Fatal("成功了...")
     // ......
   }
```














