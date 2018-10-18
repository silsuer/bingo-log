package bingo_log

import (
	"log"
	"github.com/ivpusic/grpool"
	)

// log 包
// 1. 设定一个全局变量用来输出日志
// 2. 日志自动分割
// 3. 日志包含多个错误级别
// 4. 日志对象配置（日志文件的位置，自动分割规则等）
// 5. 协程池异步输出日志

const (
	LogSyncMode = iota
	LogPoolMode
)

var Logger Log // 全局log

type Log struct {
	log.Logger
	initialized        bool         // 该日志对象是否初始化
	mode               int          // 日志记录模式 0 同步记录 2 协程池记录
	pool               *grpool.Pool // 协程池
	poolWorkerNum      int          // 协程池模式下，允许的最高协程数
	poolJobQueueLength int          // 协程池中允许的job长度
}

// 初始化结构体,如果已经初始化过会再次初始化
func (l *Log) initialize() {
	if l.mode == LogPoolMode {
		// 创建协程池
		l.pool = grpool.NewPool(l.getPoolWorkerNum(), l.getPoolJobQueueLength())
	}
	l.initialized = true
}

// 获取协程池中允许的协程数量
func (l *Log) getPoolWorkerNum() int {
	if l.poolWorkerNum == 0 {
		l.poolWorkerNum = 100
	}
	return l.poolWorkerNum
}

// 获取协程池中允许的队列长度
func (l *Log) getPoolJobQueueLength() int {
	if l.poolJobQueueLength == 0 {
		l.poolJobQueueLength = 50
	}
	return l.poolJobQueueLength
}

// 判断该结构体是否被初始化
func (l *Log) isInitialized() bool {
	return l.initialized
}

// 设置模式
func (l *Log) SetMode(m int) {
	l.mode = m
	l.initialize()
}

func (l *Log) set() {

}
