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
	initialized     bool         // 该日志对象是否初始化
	mode            int          // 日志记录模式 0 同步记录 2 协程池记录
	pool            *grpool.Pool // 协程池
	poolExpiredTime int          // 协程池模式下，每个空闲协程的存活时间
	poolWorkerNum   int          // 协程池模式下，允许的最高协程数
}

// 初始化结构体,如果已经初始化过会再次初始化
func (l *Log) initialize() {
	if l.mode == LogPoolMode {
		// 创建协程池
		l.pool = grpool.NewPool(l.getPoolWorkerNum(), l.getPoolWorkerNum())
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

// 获取协程池中空闲协程的存活时间(秒)
func (l *Log) getPoolExpiredTime() int {
	if l.poolExpiredTime == 0 {
		l.poolExpiredTime = 50
	}
	return l.poolExpiredTime
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


