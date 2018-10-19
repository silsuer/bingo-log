package bingo_log

import (
	"fmt"
	"github.com/xcltapestry/xclpkg/clcolor"
	"time"
	"os"
	"sync"
)

// 日志连接器接口，所有实现该接口的结构体都可以作为配置参数传入Log中，用来切割文件和异步执行等
type Connector interface {
	Fatal(message ...interface{})
	Error(message ...interface{})
	Warning(message ...interface{})
	Debug(message ...interface{})
	Info(message ...interface{})               // 打印
	Output(message string)             // 将信息输出到文件中
	GetMessage(message ...interface{}) string  // 将输入的信息添加抬头（例如添加打印时间等）
	GetFile(config map[string]string) *os.File // 当前日志要输出到的文件位置,传入一个map 代表配置
}

// 分为5种日志级别
// Fatal 发生严重的错误事件，导致程序退出
// Error 发生错误时间，但不影响系统运行
// Warning 警告，表明存在潜在错误
// Debug 调试信息
// Info 普通消息

// 基类连接器，实现简单的输出方法
type BaseConnector struct{
	sync.Mutex
}

func (b BaseConnector) Fatal(message ...interface{}) {
	// 红色输出
	m := "[FATAL] " + b.GetMessage(message)
	fmt.Print(clcolor.Red(m))
	b.Output(m)
}
func (b BaseConnector) Error(message ...interface{}) {
	// 紫色输出
	m := "[ERROR] " + b.GetMessage(message)
	fmt.Print(clcolor.Magenta(m))
	b.Output(m)
}
func (b BaseConnector) Warning(message ...interface{}) {
	// 黄色输出
	m := "[WARNING] " + b.GetMessage(message)
	fmt.Print(clcolor.Yellow(m))
	b.Output(m)
}
func (b BaseConnector) Debug(message ...interface{}) {
	// 蓝色输出
	m := "[DEBUG] " + b.GetMessage(message)
	fmt.Print(clcolor.Blue(m))
	b.Output(m)
}
func (b BaseConnector) Info(message ...interface{}) {
	// 绿色输出在控制台
	m := "[INFO] " + b.GetMessage(message)
	fmt.Print(clcolor.Green(m))
	// 输出在文件中
	b.Output(m)
}

func (b BaseConnector) GetMessage(message ...interface{}) string {
	// 将传入的信息扩展一下
	// 默认添加当前时间
	return "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + fmt.Sprint(message...) + "\n"
}

func (b BaseConnector) Output(message string) {
	// 获取到要输出的文件路径
	file := b.GetFile(make(map[string]string))
	defer file.Close()
	n, _ := file.Seek(0, os.SEEK_END)
	// 写入数据
	file.WriteAt([]byte(message), n)
}

// 返回一个文件句柄，用来写入数据
func (b BaseConnector) GetFile(config map[string]string) *os.File {
	// 默认情况下，输出到当前路径下的bingo.log文件中
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := dir + "/bingo.log" // 真实要保存的文件位置
	// 判断文件是否存在
	if _, err := os.Stat(path); err != nil {
		// 文件不存在,创建
		f, err := os.Create(path)
		//defer f.Close()  // 关闭操作要放在调用位置
		if err != nil {
			panic(err)
		}
		return f
	}
	// 打开该文件，追加模式
	f, err := os.OpenFile(path, os.O_WRONLY, os.ModeAppend)

	if err != nil {
		panic(err)
	}

	return f
}
