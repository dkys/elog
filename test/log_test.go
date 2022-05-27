package test

import (
	"elog"
	"log"
	"os"
	"runtime"
	"testing"
)

func TestLog(t *testing.T) {
	elog.Error("this is error")
	elog.SetErrColor("\033[41m")
	elog.ErrorF("error : %s\n", "format")
	elog.Error("this is error2")
	elog.Info("this is info")
	elog.InfoF("info : %s\n", "format")
	elog.Debug("debug")
	elog.DebugF("debug : %s\n", "format")
	elog.Info("this is info")
	elog.Log.Println("44")
	log.New(os.Stdout, "[test]", log.Ldate|log.Ltime|log.LstdFlags|log.Llongfile).Println("ni hao ")
	_, file, line, ok := runtime.Caller(0)
	println(file, line, ok)
}
