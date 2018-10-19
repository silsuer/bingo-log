package bingo_log

import (
	"testing"
)

func TestBaseConnector_Fatal(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		b    BaseConnector
		args args
	}{
		{
			name: "case",
			b:    BaseConnector{},
			args: struct{ message string }{message: "fatal testing"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BaseConnector{}
			b.Fatal(tt.args.message)
		})
	}
}

func TestBaseConnector_Error(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		b    BaseConnector
		args args
	}{
		{
			name: "case",
			b:    BaseConnector{},
			args: struct{ message string }{message: "Error testing"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BaseConnector{}
			b.Error(tt.args.message)
		})
	}
}

func TestBaseConnector_Warning(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		b    BaseConnector
		args args
	}{
		{
			name: "case",
			b:    BaseConnector{},
			args: struct{ message string }{message: "Warning testing"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BaseConnector{}
			b.Warning(tt.args.message)
		})
	}
}

func TestBaseConnector_Debug(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		b    BaseConnector
		args args
	}{
		{
			name: "case",
			b:    BaseConnector{},
			args: struct{ message string }{message: "Debug testing"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BaseConnector{}
			b.Debug(tt.args.message)
		})
	}
}

func TestBaseConnector_Info(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		b    BaseConnector
		args args
	}{
		{
			name: "case",
			b:    BaseConnector{},
			args: struct{ message string }{message: "Info testing"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BaseConnector{}
			b.Info(tt.args.message)
		})
	}
}
