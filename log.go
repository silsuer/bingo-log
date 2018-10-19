package bingo_log

import (
	"github.com/ivpusic/grpool"
	"sync"
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

const (
	FATAL   = iota
	ERROR
	WARNING
	DEBUG
	INFO
)

var Logger Log // 全局log

type Log struct {
	Connector // 内嵌连接器，用来定制化功能
	sync.Mutex
	initialized     bool         // 该日志对象是否初始化
	mode            int          // 日志记录模式 0 同步记录 2 协程池记录
	pool            *grpool.Pool // 协程池
	poolExpiredTime int          // 协程池模式下，每个空闲协程的存活时间(秒)
	poolWorkerNum   int          // 协程池模式下，允许的最高协程数
}

// 初始化结构体,如果已经初始化过会再次初始化
func (l *Log) initialize() {
	// 已初始化直接返回
	if l.initialized == true {
		return
	}
	if l.mode == LogPoolMode {
		// 创建协程池
		l.pool = grpool.NewPool(l.GetPoolWorkerNum(), l.GetPoolExpiredTime())
	}
	l.initialized = true
}

// 设置协程数量
func (l *Log) SetPoolWorkerNum(num int) {
	l.poolWorkerNum = num
}

// 设置协程存活时间
func (l *Log) SetPoolExpiredTime(t int) {
	l.poolExpiredTime = t
}

// 获取协程池中允许的协程数量
func (l *Log) GetPoolWorkerNum() int {
	if l.poolWorkerNum == 0 {
		l.poolWorkerNum = 100
	}
	return l.poolWorkerNum
}

// 获取协程池中空闲协程的存活时间(秒)
func (l *Log) GetPoolExpiredTime() int {
	if l.poolExpiredTime == 0 {
		l.poolExpiredTime = 50
	}
	return l.poolExpiredTime
}

func NewLog(mode int) *Log {
	l := &Log{}
	l.SetMode(mode)
	l.initialize()
	return l
}

// 判断该结构体是否被初始化
func (l *Log) IsInitialized() bool {
	return l.initialized
}

// 设置模式
func (l *Log) SetMode(m int) {
	l.mode = m
	//l.initialize()
}

// 加载连接器
func (l *Log) LoadConnector(conn Connector) {
	l.Connector = conn
}

// 重写5种日志级别的打印函数
func (l *Log) Fatal(message string) {
	// 根据模式
	l.exec(l.Connector.Fatal, message)
}

func (l *Log) Error(message string) {
	l.exec(l.Connector.Error, message)
}

func (l *Log) Warning(message string) {
	l.exec(l.Connector.Warning, message)
}

func (l *Log) Debug(message string) {
	l.exec(l.Connector.Debug, message)
}

func (l *Log) Info(message string) {
	l.exec(l.Connector.Info, message)
}

func (l *Log) exec(f func(message ...interface{}), message string) {
	// 同步模式
	if l.mode == LogSyncMode {
		l.Lock()
		defer l.Unlock()
		f(message)
	} else if l.mode == LogPoolMode { // 协程池异步模式
		l.initialize() // 先初始化
		l.Lock()
		defer l.Unlock()
		l.AddWaitCount(1)
		l.pool.JobQueue <- func() {
			f(message)
			defer l.pool.JobDone()
		}
	}
}

// 等待所有job执行完毕
func (l *Log) WaitAll() {
	l.pool.WaitAll()
}

func (l *Log) AddWaitCount(count int) {
	l.pool.WaitCount(count)
}
