package main

import (
	"github.com/silsuer/bingo-log"
	"strconv"
	"fmt"
)

func main() {
	//fmt.Println(os.Getwd())
	//c:= bingo_log.NewLog(bingo_log.LogSyncMode)
	//b := new(bingo_log.BaseConnector)
	//c.LoadConnector(b) // 加载连接器
	//c.Warning("ddd")
	//c := bingo_log.NewLog(bingo_log.LogPoolMode)
	c:= bingo_log.NewLog(bingo_log.LogSyncMode)
	c.SetPoolExpiredTime(100)
	c.SetPoolWorkerNum(100)

	b := new(bingo_log.BaseConnector)
	c.LoadConnector(b)

	//s := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		//s.Add()
		c.Debug("这是第" + strconv.Itoa(i) + "次输入")
	}
    fmt.Println("Done!")
}
