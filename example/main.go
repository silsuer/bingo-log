package main

import "github.com/silsuer/bingo-log"

func main() {

	// 判定数据
	c := new(bingo_log.KirinConnector)
	l := bingo_log.NewLog(bingo_log.LogSyncMode)
	l.LoadConnector(c)
	l.Info("KirinConnector testing")
	//bingo_log.KirinGetMessage("asd")
	//bingo_log.KirinGetFile(make(map[string]string))
	//bingo_log.CreateFile("/Users/silsuer/go/src/github.com/silsuer/bingo-log/aa/aaa/aaaa/bingo_2018_10_19.log/")
	//c:= bingo_log.NewLog(bingo_log.LogSyncMode)
	//c.SetPoolExpiredTime(100)
	//c.SetPoolWorkerNum(100)
	//
	//b := new(bingo_log.BaseConnector)
	//c.LoadConnector(b)
	//
	////s := sync.WaitGroup{}
	//
	//for i := 0; i < 1000; i++ {
	////s.Add()
	//c.Debug("这是第" + strconv.Itoa(i) + "次输入")
	//}
	//fmt.Println("Done!")
}
