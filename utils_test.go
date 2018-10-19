package bingo_log

import (
	"testing"
	"os"
)

func TestKirinGetMessage(t *testing.T) {
	KirinGetMessage(INFO, "testing")
	KirinGetMessage(DEBUG, "testing")
	KirinGetMessage(WARNING, "testing")
	KirinGetMessage(ERROR, "testing")
	KirinGetMessage(FATAL, "testing")
}

func TestKirinGetFile(t *testing.T) {
	KirinGetFile(make(map[string]string))

	c := make(map[string]string)
	dir, _ := os.Getwd()
	c["root"] = dir + "/example/logs"
	c["format"] = "2006_01_02"
	KirinGetFile(c)
}
