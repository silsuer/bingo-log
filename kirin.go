package bingo_log

import (
	"os"
	"fmt"
	"github.com/xcltapestry/xclpkg/clcolor"
)

// 默认数据
//Gluttonous

// 内置的第一个连接器

// 抬头应该是  [等级][日期时间] 调用文件+调用行号 + 打印信息
// 对于5种方式正常打印
// 分片规则：按照日期分类，给定一个根目录，按照当前日期创建文件

// 内置连接器
type KirinConnector struct {
	BaseConnector
	config map[string]string
}

func (k KirinConnector) Fatal(message ...interface{}) {
	// 红色输出
	m := k.GetMessage(FATAL, message...)
	fmt.Print(clcolor.Red(m))
	k.Output(m)
}

func (k KirinConnector) Error(message ...interface{}) {
	// 紫色输出
	m := k.GetMessage(ERROR, message...)
	fmt.Print(clcolor.Magenta(m))
	k.Output(m)
}
func (k KirinConnector) Warning(message ...interface{}) {
	// 黄色输出
	m := k.GetMessage(WARNING, message...)
	fmt.Print(clcolor.Yellow(m))
	k.Output(m)
}
func (k KirinConnector) Debug(message ...interface{}) {
	// 蓝色输出
	m := k.GetMessage(DEBUG, message...)
	fmt.Print(clcolor.Blue(m))
	k.Output(m)
}
func (k KirinConnector) Info(message ...interface{}) {
	// 绿色输出在控制台
	m := k.GetMessage(INFO, message...)
	fmt.Print(clcolor.Green(m))
	// 输出在文件中
	k.Output(m)
}

func NewKirinConnector(config map[string]string) *KirinConnector {
	k := new(KirinConnector)
	k.config = config
	return k
}

func (k KirinConnector) Output(message string) {
	// 获取到要输出的文件路径
	file := k.GetFile(k.config)
	defer file.Close()
	n, _ := file.Seek(0, os.SEEK_END)
	// 写入数据
	file.WriteAt([]byte(message), n)
}

func (k KirinConnector) GetMessage(degree int, message ...interface{}) string {
	return KirinGetMessage(degree, message...)
}

func (k KirinConnector) GetFile(config map[string]string) *os.File {
	return KirinGetFile(config)
}
