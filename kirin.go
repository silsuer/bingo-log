package bingo_log

import "os"

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

func NewKirinConnector(config map[string]string)  {

}

func (k KirinConnector) Output(message string) {
	// 获取到要输出的文件路径
	config := make(map[string]string)
	k.Lock()
	config["root"], _ = os.Getwd()
	config["format"] = "2006-01-02"
	k.Unlock()
	file := k.GetFile(config)
	defer file.Close()
	n, _ := file.Seek(0, os.SEEK_END)
	// 写入数据
	file.WriteAt([]byte(message), n)
}

func (k KirinConnector) GetMessage(message ...interface{}) string {
	return KirinGetMessage(message)
}

func (k KirinConnector) GetFile(config map[string]string) *os.File {
	return KirinGetFile(config)
}
