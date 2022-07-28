package elog

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	tw = &TimeWriter{
		Dir:      "./",
		IsStdout: true,
	}
	errorColor  = "\u001B[3;31m"
	debugColor  = "\u001B[4;33m"
	infoColor   = "\u001B[1;34m"
	errorPrefix = "[ERROR] "
	debugPrefix = "[DEBUG] "
	infoPrefix  = "[INFO] "
	Log         = log.New(tw, "", log.Ldate|log.Ltime|log.LstdFlags|log.Llongfile)
	callDepth   = 2
)

// TimeWriter 日志分割结构体
type TimeWriter struct {
	Dir         string // 日志所在目录
	IsStdout    bool   // 是否使用标准输出
	mu          sync.Mutex
	file        *os.File
	curFilename string
	color       string
}

func SetCallDepth(calldepth int) {
	mu.Lock()
	defer mu.Unlock()
	callDepth = calldepth
}

func Debug(v ...any) {
	if Level > InfoLevel {
		Log.SetPrefix(debugPrefix)
		tw.color = debugColor
		Log.Output(callDepth, fmt.Sprintln(v...))
	}
}

func DebugF(format string, v ...any) {
	Log.SetPrefix(debugPrefix)
	tw.color = debugColor
	if Level > InfoLevel {
		Log.Output(callDepth, fmt.Sprintf(format, v...))
	}
}

func Info(v ...any) {
	if Level > ErrorLevel {
		Log.SetPrefix(infoPrefix)
		tw.color = infoColor
		Log.Output(callDepth, fmt.Sprintln(v...))
	}
}

func InfoF(format string, v ...any) {
	if Level > ErrorLevel {
		Log.SetPrefix(infoPrefix)
		tw.color = infoColor
		Log.Output(callDepth, fmt.Sprintf(format, v...))
	}
}

func Error(v ...any) {
	if Level > Disabled {
		Log.SetPrefix(errorPrefix)
		tw.color = errorColor
		Log.Output(callDepth, fmt.Sprintln(v...))
	}
}

func ErrorF(format string, v ...any) {
	if Level > Disabled {
		Log.SetPrefix(errorPrefix)
		tw.color = errorColor
		Log.Output(callDepth, fmt.Sprintf(format, v...))
	}
}

func Panicln(v ...any) {
	Log.SetPrefix(errorPrefix)
	tw.color = errorColor
	s := fmt.Sprintln(v...)
	Log.Output(callDepth, s)
	panic(s)
	//Log.Output(2, fmt.Sprintln(v...))
}

func PanicF(format string, v ...any) {
	Log.SetPrefix(errorPrefix)
	tw.color = errorColor
	s := fmt.Sprintf(format, v...)
	Log.Output(callDepth, s)
	panic(s)
}
func SetErrColor(color string) {
	errorColor = color
}

func (t *TimeWriter) Write(p []byte) (n int, err error) {
	if t.IsStdout {
		p = bytes.TrimRight(p, "\n")
		p = append([]byte(t.color), p...)
		p = append(p, []byte("\u001B[0m\n")...)
		return os.Stdout.Write(p)
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.file == nil {
		if e := t.openFile(); e != nil {
			return 0, e
		}
	}

	if t.curFilename != t.fileName() {
		if e := t.rotate(); e != nil {
			return 0, e
		}
	}

	n, err = t.file.Write(p)
	return n, err
}

// 打开文件
func (t *TimeWriter) openFile() error {
	t.curFilename = t.fileName()
	if _, err := os.Stat(t.curFilename); os.IsNotExist(err) {
		err = os.MkdirAll(t.Dir, 0744)
		if err != nil {
			return err
		}
	}
	file, err := os.OpenFile(t.curFilename, os.O_CREATE|os.O_APPEND|os.O_SYNC|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	t.file = file
	return nil
}

// 获取当前文件名称
func (t *TimeWriter) fileName() string {
	name := fmt.Sprintf("%s.log", time.Now().Format("2006-01-02"))
	if t.Dir != "" {
		return filepath.Join(t.Dir, name)
	}
	return filepath.Join(os.TempDir(), name)
}

// 更换文件
func (t *TimeWriter) rotate() error {
	if e := t.file.Close(); e != nil {
		return e
	}
	return t.openFile()
}
