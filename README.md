# elog

---
A simple Golang log library.Split logs by day and support color-coded command line output.

# Installation

````
GO111MODULE=on
go get github.com/dkys/elog
````

# Using

```go
package main

import (
	"github.com/dkys/elog"
)

func main() {
	elog.Error("error")
	elog.SetErrColor("\033[41m")
	elog.ErrorF("error : %s", "format")
	elog.Info("info")
	elog.InfoF("info : %s", "format")
	elog.Debug("debug")
	elog.DebugF("debug : %s", "format")
	elog.Log.Println("44")
}

```
