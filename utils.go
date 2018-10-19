package bingo_log

import (
	"time"
	"fmt"
	"os"
	"strings"
	"runtime"
	"strconv"
)

// 内置一些函数，需要的时候，直接实现一个连接器，组合数据即可

// 抬头  [时间]+文件+行数+消息
func KirinGetMessage(degree int, message ...interface{}) string {

	var title string
	switch degree {
	case FATAL:
		title = "[FATAL] "
	case ERROR:
		title = "[ERROR] "
	case WARNING:
		title = "[WARNING]"
	case DEBUG:
		title = "[DEBUG] "
	case INFO:
		title = "[INFO]"
	default:
		title = "[UNKNOWN]"
	}

	_, file, line, ok := runtime.Caller(7)
	if !ok {
		file = "???"
		line = 0
	}

	return `
` + title + `
[` + time.Now().Format("2006-01-02 15:04:05") + `] 
[FILE] ` + file + `
[LINE] ` + strconv.Itoa(line) + `
[CONTENT] ` + fmt.Sprint(message...) + `
`
}

// 配置一部分根目录
// 根目录的方式：给定一个根目录，按照当前日期创建文件
// 得到要输入的文件句柄
func KirinGetFile(config map[string]string) *os.File {
	// 根目录
	format := "2006_01_02"
	root := "/"
	if _, ok := config["root"]; !ok {
		root, _ = os.Getwd()
	} else {
		root = config["root"]
	}
	// 格式
	if _, ok := config["format"]; ok {
		format = config["format"]
	}
	// 在根目录下生成文件
	t := time.Now().Format(format)
	// 生成文件完整路径
	path := root + "/bingo_" + t + ".log"

	// 打开或创建文件，得到文件句柄
	return OpenOrCreateFile(path)
}

func OpenOrCreateFile(path string) *os.File {
	if _, err := os.Stat(path); os.IsNotExist(err) { // 文件不存在
		return CreateFile(path)
	} else {
		f, err := os.OpenFile(path, os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		return f
	}
}

// 创建文件
// 创建目录，如果以 / 号结尾，直接级联建立该目录
// 如果没有 / 号结尾，将最后一个 / 号后面的字符串作为文件忽略掉
func CreateFile(path string) *os.File {
	// 判断文件是否存在,如果不存在就创建
	if _, err := os.Stat(path); os.IsNotExist(err) { // 不存在,创建
		if strings.HasSuffix(path, "/") { //  创建目录
			// 建立目录
			os.MkdirAll(path, 0755)
			return nil
		} else {
			// 最后一个 / 之后的数据建立文件，前面的字符串建立文件夹
			dirPath := path[0:strings.LastIndex(path, "/")]
			os.MkdirAll(dirPath, 0755) // 创建文件
			f, err := os.Create(path)
			if err != nil {
				panic(err)
			}
			return f
		}
	} else {
		f, err := os.OpenFile(path, os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		return f
	}

}
