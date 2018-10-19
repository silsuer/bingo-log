package bingo_log

import (
	"testing"
	"github.com/ivpusic/grpool"
	"reflect"
	"fmt"
)

func TestLog_initialize(t *testing.T) {
	type fields struct {
		Connector       Connector
		initialized     bool
		mode            int
		pool            *grpool.Pool
		poolExpiredTime int
		poolWorkerNum   int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "case1",
			fields: struct {
				Connector       Connector
				initialized     bool
				mode            int
				pool            *grpool.Pool
				poolExpiredTime int
				poolWorkerNum   int
			}{Connector: BaseConnector{}, initialized: false, mode: LogSyncMode, pool: grpool.NewPool(100, 50), poolExpiredTime: 100, poolWorkerNum: 50},
		},
		{
			name: "case2",
			fields: struct {
				Connector       Connector
				initialized     bool
				mode            int
				pool            *grpool.Pool
				poolExpiredTime int
				poolWorkerNum   int
			}{Connector: BaseConnector{}, initialized: true, mode: LogSyncMode, pool: grpool.NewPool(100, 50), poolExpiredTime: 100, poolWorkerNum: 50},
		},
		{
			name: "case3",
			fields: struct {
				Connector       Connector
				initialized     bool
				mode            int
				pool            *grpool.Pool
				poolExpiredTime int
				poolWorkerNum   int
			}{Connector: BaseConnector{}, initialized: false, mode: LogPoolMode, pool: grpool.NewPool(100, 50), poolExpiredTime: 100, poolWorkerNum: 50},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				Connector:       tt.fields.Connector,
				initialized:     tt.fields.initialized,
				mode:            tt.fields.mode,
				pool:            tt.fields.pool,
				poolExpiredTime: tt.fields.poolExpiredTime,
				poolWorkerNum:   tt.fields.poolWorkerNum,
			}
			l.initialize()
		})
	}
}

func TestLog_SetPoolWorkerNum(t *testing.T) {
	type fields struct {
		Connector       Connector
		initialized     bool
		mode            int
		pool            *grpool.Pool
		poolExpiredTime int
		poolWorkerNum   int
	}
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case",
			fields: struct {
				Connector       Connector
				initialized     bool
				mode            int
				pool            *grpool.Pool
				poolExpiredTime int
				poolWorkerNum   int
			}{Connector: BaseConnector{}, initialized: false, mode: LogPoolMode, pool: grpool.NewPool(100, 50), poolExpiredTime: 100, poolWorkerNum: 50},
			args: struct{ num int }{num: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				Connector:       tt.fields.Connector,
				initialized:     tt.fields.initialized,
				mode:            tt.fields.mode,
				pool:            tt.fields.pool,
				poolExpiredTime: tt.fields.poolExpiredTime,
				poolWorkerNum:   tt.fields.poolWorkerNum,
			}
			l.SetPoolWorkerNum(tt.args.num)
		})
	}
}

func TestLog_SetPoolExpiredTime(t *testing.T) {
	type fields struct {
		Connector       Connector
		initialized     bool
		mode            int
		pool            *grpool.Pool
		poolExpiredTime int
		poolWorkerNum   int
	}
	type args struct {
		t int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case",
			fields: struct {
				Connector       Connector
				initialized     bool
				mode            int
				pool            *grpool.Pool
				poolExpiredTime int
				poolWorkerNum   int
			}{Connector: BaseConnector{}, initialized: false, mode: LogPoolMode, pool: grpool.NewPool(100, 50), poolExpiredTime: 100, poolWorkerNum: 50},
			args: struct{ t int }{t: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				Connector:       tt.fields.Connector,
				initialized:     tt.fields.initialized,
				mode:            tt.fields.mode,
				pool:            tt.fields.pool,
				poolExpiredTime: tt.fields.poolExpiredTime,
				poolWorkerNum:   tt.fields.poolWorkerNum,
			}
			l.SetPoolExpiredTime(tt.args.t)
		})
	}
}

func TestLog_GetPoolWorkerNum(t *testing.T) {
	type fields struct {
		Connector       Connector
		initialized     bool
		mode            int
		pool            *grpool.Pool
		poolExpiredTime int
		poolWorkerNum   int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "case",
			fields: struct {
				Connector       Connector
				initialized     bool
				mode            int
				pool            *grpool.Pool
				poolExpiredTime int
				poolWorkerNum   int
			}{Connector: BaseConnector{}, initialized: false, mode: LogPoolMode, pool: grpool.NewPool(100, 50), poolExpiredTime: 50, poolWorkerNum: 100},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				Connector:       tt.fields.Connector,
				initialized:     tt.fields.initialized,
				mode:            tt.fields.mode,
				pool:            tt.fields.pool,
				poolExpiredTime: tt.fields.poolExpiredTime,
				poolWorkerNum:   tt.fields.poolWorkerNum,
			}
			if got := l.GetPoolWorkerNum(); got != tt.want {
				t.Errorf("Log.GetPoolWorkerNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog_GetPoolExpiredTime(t *testing.T) {
	type fields struct {
		Connector       Connector
		initialized     bool
		mode            int
		pool            *grpool.Pool
		poolExpiredTime int
		poolWorkerNum   int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "case",
			fields: struct {
				Connector       Connector
				initialized     bool
				mode            int
				pool            *grpool.Pool
				poolExpiredTime int
				poolWorkerNum   int
			}{Connector: BaseConnector{}, initialized: false, mode: LogPoolMode, pool: grpool.NewPool(100, 50), poolExpiredTime: 50, poolWorkerNum: 100},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				Connector:       tt.fields.Connector,
				initialized:     tt.fields.initialized,
				mode:            tt.fields.mode,
				pool:            tt.fields.pool,
				poolExpiredTime: tt.fields.poolExpiredTime,
				poolWorkerNum:   tt.fields.poolWorkerNum,
			}
			if got := l.GetPoolExpiredTime(); got != tt.want {
				t.Errorf("Log.GetPoolExpiredTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLog(t *testing.T) {
	type args struct {
		mode int
	}
	tests := []struct {
		name string
		args args
		want *Log
	}{
		{
			name: "case",
			args: struct{ mode int }{mode: LogSyncMode},
			want: &Log{mode: LogSyncMode, initialized: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLog(tt.args.mode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog_IsInitialized(t *testing.T) {
	type fields struct {
		Connector       Connector
		initialized     bool
		mode            int
		pool            *grpool.Pool
		poolExpiredTime int
		poolWorkerNum   int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "case",
			fields: struct {
				Connector       Connector
				initialized     bool
				mode            int
				pool            *grpool.Pool
				poolExpiredTime int
				poolWorkerNum   int
			}{Connector: BaseConnector{}, initialized: false, mode: LogPoolMode, pool: grpool.NewPool(100, 50), poolExpiredTime: 50, poolWorkerNum: 100},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				Connector:       tt.fields.Connector,
				initialized:     tt.fields.initialized,
				mode:            tt.fields.mode,
				pool:            tt.fields.pool,
				poolExpiredTime: tt.fields.poolExpiredTime,
				poolWorkerNum:   tt.fields.poolWorkerNum,
			}
			if got := l.IsInitialized(); got != tt.want {
				t.Errorf("Log.IsInitialized() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog_SetMode(t *testing.T) {
	type fields struct {
		Connector       Connector
		initialized     bool
		mode            int
		pool            *grpool.Pool
		poolExpiredTime int
		poolWorkerNum   int
	}
	type args struct {
		m int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case",
			fields: struct {
				Connector       Connector
				initialized     bool
				mode            int
				pool            *grpool.Pool
				poolExpiredTime int
				poolWorkerNum   int
			}{Connector: BaseConnector{}, initialized: false, mode: LogPoolMode, pool: grpool.NewPool(100, 50), poolExpiredTime: 50, poolWorkerNum: 100},
			args: struct{ m int }{m: LogPoolMode},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				Connector:       tt.fields.Connector,
				initialized:     tt.fields.initialized,
				mode:            tt.fields.mode,
				pool:            tt.fields.pool,
				poolExpiredTime: tt.fields.poolExpiredTime,
				poolWorkerNum:   tt.fields.poolWorkerNum,
			}
			l.SetMode(tt.args.m)
		})
	}
}

func TestLog_LoadConnector(t *testing.T) {
	type fields struct {
		Connector       Connector
		initialized     bool
		mode            int
		pool            *grpool.Pool
		poolExpiredTime int
		poolWorkerNum   int
	}
	type args struct {
		conn Connector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case",
			fields: struct {
				Connector       Connector
				initialized     bool
				mode            int
				pool            *grpool.Pool
				poolExpiredTime int
				poolWorkerNum   int
			}{Connector: BaseConnector{}, initialized: false, mode: LogPoolMode, pool: grpool.NewPool(100, 50), poolExpiredTime: 50, poolWorkerNum: 100},
			args: struct{ conn Connector }{conn: BaseConnector{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				Connector:       tt.fields.Connector,
				initialized:     tt.fields.initialized,
				mode:            tt.fields.mode,
				pool:            tt.fields.pool,
				poolExpiredTime: tt.fields.poolExpiredTime,
				poolWorkerNum:   tt.fields.poolWorkerNum,
			}
			l.LoadConnector(tt.args.conn)
		})
	}
}

func TestLog_Fatal(t *testing.T) {
	type fields struct {
		Connector       Connector
		initialized     bool
		mode            int
		pool            *grpool.Pool
		poolExpiredTime int
		poolWorkerNum   int
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case",
			fields: struct {
				Connector       Connector
				initialized     bool
				mode            int
				pool            *grpool.Pool
				poolExpiredTime int
				poolWorkerNum   int
			}{Connector: BaseConnector{}, initialized: true, mode: LogSyncMode, pool: grpool.NewPool(100, 50), poolExpiredTime: 50, poolWorkerNum: 100},
			args: struct{ message string }{message: "Fatal testing"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				Connector:       tt.fields.Connector,
				initialized:     tt.fields.initialized,
				mode:            tt.fields.mode,
				pool:            tt.fields.pool,
				poolExpiredTime: tt.fields.poolExpiredTime,
				poolWorkerNum:   tt.fields.poolWorkerNum,
			}
			l.Fatal(tt.args.message)
		})
	}
}

func TestLog_Error(t *testing.T) {
	type fields struct {
		Connector       Connector
		initialized     bool
		mode            int
		pool            *grpool.Pool
		poolExpiredTime int
		poolWorkerNum   int
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case",
			fields: struct {
				Connector       Connector
				initialized     bool
				mode            int
				pool            *grpool.Pool
				poolExpiredTime int
				poolWorkerNum   int
			}{Connector: BaseConnector{}, initialized: true, mode: LogPoolMode, pool: grpool.NewPool(100, 50), poolExpiredTime: 50, poolWorkerNum: 100},
			args: struct{ message string }{message: "Error testing"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				Connector:       tt.fields.Connector,
				initialized:     tt.fields.initialized,
				mode:            tt.fields.mode,
				pool:            tt.fields.pool,
				poolExpiredTime: tt.fields.poolExpiredTime,
				poolWorkerNum:   tt.fields.poolWorkerNum,
			}
			l.Error(tt.args.message)
			l.pool.WaitAll()
		})
	}
}

func TestLog_Warning(t *testing.T) {
	type fields struct {
		Connector       Connector
		initialized     bool
		mode            int
		pool            *grpool.Pool
		poolExpiredTime int
		poolWorkerNum   int
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case",
			fields: struct {
				Connector       Connector
				initialized     bool
				mode            int
				pool            *grpool.Pool
				poolExpiredTime int
				poolWorkerNum   int
			}{Connector: BaseConnector{}, initialized: false, mode: LogPoolMode, pool: grpool.NewPool(100, 50), poolExpiredTime: 0, poolWorkerNum: 0},
			args: struct{ message string }{message: "Warning testing"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				Connector:       tt.fields.Connector,
				initialized:     tt.fields.initialized,
				mode:            tt.fields.mode,
				pool:            tt.fields.pool,
				poolExpiredTime: tt.fields.poolExpiredTime,
				poolWorkerNum:   tt.fields.poolWorkerNum,
			}
			l.Warning(tt.args.message)
			//l.WaitAll()
			l.WaitAll()
		})
	}
}

func TestLog_Debug(t *testing.T) {
	type fields struct {
		Connector       Connector
		initialized     bool
		mode            int
		pool            *grpool.Pool
		poolExpiredTime int
		poolWorkerNum   int
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case",
			fields: struct {
				Connector       Connector
				initialized     bool
				mode            int
				pool            *grpool.Pool
				poolExpiredTime int
				poolWorkerNum   int
			}{Connector: BaseConnector{}, initialized: false, mode: LogPoolMode, pool: grpool.NewPool(100, 50), poolExpiredTime: 50, poolWorkerNum: 100},
			args: struct{ message string }{message: "Debug testing"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				Connector:       tt.fields.Connector,
				initialized:     tt.fields.initialized,
				mode:            tt.fields.mode,
				pool:            tt.fields.pool,
				poolExpiredTime: tt.fields.poolExpiredTime,
				poolWorkerNum:   tt.fields.poolWorkerNum,
			}
			l.Debug(tt.args.message)
			l.WaitAll()
		})
	}
}

func TestLog_Info(t *testing.T) {
	type fields struct {
		Connector       Connector
		initialized     bool
		mode            int
		pool            *grpool.Pool
		poolExpiredTime int
		poolWorkerNum   int
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case",
			fields: struct {
				Connector       Connector
				initialized     bool
				mode            int
				pool            *grpool.Pool
				poolExpiredTime int
				poolWorkerNum   int
			}{Connector: BaseConnector{}, initialized: false, mode: LogPoolMode, pool: grpool.NewPool(100, 50), poolExpiredTime: 50, poolWorkerNum: 100},
			args: struct{ message string }{message: "Info testing"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				Connector:       tt.fields.Connector,
				initialized:     tt.fields.initialized,
				mode:            tt.fields.mode,
				pool:            tt.fields.pool,
				poolExpiredTime: tt.fields.poolExpiredTime,
				poolWorkerNum:   tt.fields.poolWorkerNum,
			}
			l.Info(tt.args.message)
			l.WaitAll()
		})
	}
}

func TestLog_exec(t *testing.T) {
	type fields struct {
		Connector       Connector
		initialized     bool
		mode            int
		pool            *grpool.Pool
		poolExpiredTime int
		poolWorkerNum   int
	}
	type args struct {
		f       func(message ...interface{})
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case",
			fields: struct {
				Connector       Connector
				initialized     bool
				mode            int
				pool            *grpool.Pool
				poolExpiredTime int
				poolWorkerNum   int
			}{Connector: BaseConnector{}, initialized: false, mode: LogPoolMode, pool: grpool.NewPool(100, 50), poolExpiredTime: 50, poolWorkerNum: 100},
			args: struct {
				f       func(message ...interface{})
				message string
			}{f: func(message ...interface{}) {
				fmt.Println(message)
			}, message: "Exec testing"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				Connector:       tt.fields.Connector,
				initialized:     tt.fields.initialized,
				mode:            tt.fields.mode,
				pool:            tt.fields.pool,
				poolExpiredTime: tt.fields.poolExpiredTime,
				poolWorkerNum:   tt.fields.poolWorkerNum,
			}
			l.exec(tt.args.f, tt.args.message)
			l.pool.WaitAll()
		})
	}
}
