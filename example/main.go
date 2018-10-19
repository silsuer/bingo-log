package main

import (
	"github.com/silsuer/bingo-log"
	"fmt"
	"os"
)

func main() {

	//log := bingo_log.NewLog(bingo_log.LogPoolMode)
	//log.SetPoolExpiredTime(100)
	//log.SetPoolWorkerNum(100)

	log := bingo_log.NewLog(bingo_log.LogSyncMode)

	//e := new(ExampleConnector)
	//log.LoadConnector(e)

	//config := make(map[string]string)
	//config["root"] = "/set/any/path"
	//config["format"] = "2006_01_02"

	conn := new(bingo_log.BaseConnector)
	//c := bingo_log.NewKirinConnector(make(map[string]string))

	log.LoadConnector(conn)
	//log.LoadConnector(c)

	log.Info("testing")
	log.Debug("testing")
	log.Warning("testing")
	log.Error("testing")
	log.Fatal("testing")

}

type ExampleConnector struct{}

func (e ExampleConnector) GetMessage(degree int, message ...interface{}) string {
	return "[出错了...]" + fmt.Sprint(message...)
}

func (e ExampleConnector) GetFile(config map[string]string) *os.File {
	return bingo_log.OpenOrCreateFile("/Users/silsuer/go/error.log")
}

func (e ExampleConnector) Output(message string) {
	// 获取到要输出的文件路径
	file := e.GetFile(make(map[string]string))
	defer file.Close()
	n, _ := file.Seek(0, os.SEEK_END)
	// 写入数据
	file.WriteAt([]byte(message), n)
}

func (e ExampleConnector) Fatal(message ...interface{}) {
	m := e.GetMessage(bingo_log.FATAL, message...)
	fmt.Print(m)
	e.Output(m)
}

func (e ExampleConnector) Error(message ...interface{}) {
	m := e.GetMessage(bingo_log.ERROR, message...)
	fmt.Print(m)
	e.Output(m)
}
func (e ExampleConnector) Warning(message ...interface{}) {
	m := e.GetMessage(bingo_log.WARNING, message...)
	fmt.Print(m)
	e.Output(m)
}
func (e ExampleConnector) Debug(message ...interface{}) {
	m := e.GetMessage(bingo_log.DEBUG, message...)
	fmt.Print(m)
	e.Output(m)
}
func (e ExampleConnector) Info(message ...interface{}) {
	m := e.GetMessage(bingo_log.INFO, message...)
	fmt.Print(m)
	e.Output(m)
}
