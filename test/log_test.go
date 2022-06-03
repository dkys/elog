package test

import (
	"github.com/dkys/elog"
	"testing"
)

func TestLog(t *testing.T) {
	elog.SetErrColor("\033[41m")
	elog.SetLevel(elog.Disabled)
	elog.Error("error")
	elog.ErrorF("error : %s", "format")
	elog.Info("info")
	elog.InfoF("info : %s", "format")
	elog.Debug("debug")
	elog.DebugF("debug : %s", "format")
	//elog.Panicln("Panic")
	//elog.PanicF("Panic : %s", "format")
}
