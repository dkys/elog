package elog

import (
	"sync"
)

const (
	Disabled = iota
	ErrorLevel
	InfoLevel
	DEBUGLevel
)

var mu sync.Mutex
var Level int = DEBUGLevel

func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()
	Level = level
}
