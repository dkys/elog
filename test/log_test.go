package test

import (
	"errors"
	"github.com/dkys/elog"
	"testing"
)

func TestLog(t *testing.T) {
	//elog.Exit("退出")
	//elog.ExitF("退出2%s", "ttt")
	elog.SetErrColor("\033[41m")
	//elog.SetLevel(elog.Disabled)
	elog.Error("error")
	elog.ErrorF("error : %s", "format")
	elog.Info("info")
	elog.InfoF("info : %s", "format")
	elog.Debug("debug")
	elog.DebugF("debug : %s", "format")
	//elog.Panicln("Panic")
	//elog.PanicF("Panic : %s", "format")
	elog.SetCallDepth(3)
	newInfo()
	newError(errors.New("test...."))
	elog.SetCallDepth(2)
	elog.IsStdout(false)
	elog.SetOutPath("./logs")
	newInfo()
}

func newError(e error) {
	if e != nil {
		elog.Error(e.Error())
	}
}
func newInfo() {
	elog.Info("info")
}
