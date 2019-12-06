// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)

// Log is a global variable used for loggig
var Log *logrus.Logger

// ContextHook ...
type ContextHook struct{}

// Levels ...
func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire ...
func (hook ContextHook) Fire(entry *logrus.Entry) error {
	if pc, file, line, ok := runtime.Caller(10); ok {
		funcName := runtime.FuncForPC(pc).Name()

		entry.Data["source"] = fmt.Sprintf("%s:%v:%s", path.Base(file), line, path.Base(funcName))
	}

	return nil
}
