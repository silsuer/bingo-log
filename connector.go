package bingo_log

import (
	"fmt"
	"github.com/xcltapestry/xclpkg/clcolor"
	"time"
	"os"
)

// 日志连接器接口，所有实现该接口的结构体都可以作为配置参数传入Log中，用来切割文件和异步执行等
type Connector interface {
	Fatal(message string)
	Error(message string)
	Warning(message string)
	Debug(message string)
	Info(message string)               // 打印
	Output(message string)             // 将信息输出到文件中
	GetMessage(message string) string  // 将输入的信息添加抬头（例如添加打印时间等）
	GetFile(configs []string) *os.File // 当前日志要输出到的文件位置,传入一个数组代表配置
}

// 分为5种日志级别
// Fatal 发生严重的错误事件，导致程序退出
// Error 发生错误时间，但不影响系统运行
// Warning 警告，表明存在潜在错误
// Debug 调试信息
// Info 普通消息

// 基类连接器，实现简单的输出方法
type BaseConnector struct{}

func (b *BaseConnector) Fatal(message string) {
	// 红色输出
	message = "[FATAL] " + b.GetMessage(message)
	fmt.Println(clcolor.Red(message))
}
func (b *BaseConnector) Error(message string) {
	// 紫色输出
	message = "[ERROR] " + b.GetMessage(message)
	fmt.Println(clcolor.Magenta(message))
}
func (b *BaseConnector) Warning(message string) {
	// 黄色输出
	message = "[WARNING] " + b.GetMessage(message)
	fmt.Println(clcolor.Yellow(message))
}
func (b *BaseConnector) Debug(message string) {
	// 蓝色输出
	message = "[DEBUG] " + b.GetMessage(message)
	fmt.Println(clcolor.Blue(message))
}
func (b *BaseConnector) Info(message string) {
	// 绿色输出在控制台
	message = "[INFO] " + b.GetMessage(message)
	fmt.Println(clcolor.Green(message))
	// 输出在文件中
}

func (b *BaseConnector) GetMessage(message string) string {
	// 将传入的信息扩展一下
	// 默认添加当前时间
	return "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + message
}

func (b *BaseConnector) Output(message string) {
	// 获取到要输出的文件路径
}

func (b *BaseConnector) GetFile() *os.File {

}
