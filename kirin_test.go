package bingo_log

import (
	"testing"
)

func TestKirinConnector_GetMessage(t *testing.T) {
	// 测试新的连接器
	c := NewKirinConnector(make(map[string]string))
	l := NewLog(LogSyncMode)
	l.LoadConnector(c)
	l.Info("KirinConnector Info testing")
	l.Fatal("KirinConnector Fatal testing")
	l.Warning("KirinConnector Warning testing")
	l.Debug("KirinConnector Debug testing")
	l.Error("KirinConnector Error testing")
}

//
//func TestKirinConnector_GetFile(t *testing.T) {
//	type fields struct {
//		BaseConnector BaseConnector
//	}
//	type args struct {
//		config map[string]string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   *os.File
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			k := KirinConnector{
//				BaseConnector: tt.fields.BaseConnector,
//			}
//			if got := k.GetFile(tt.args.config); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("KirinConnector.GetFile() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
