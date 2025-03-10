package log

/*
	CIPHERC2 Implant Framework
	Copyright (C) 2019  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	envVarName = "CIPHERC2_ROOT_DIR"
)

var (
	// RootLoggerName - Root logger name, contains all log data
	RootLoggerName = "root"
	// RootLogger - Root Logger
	RootLogger = rootLogger()
)

// NamedLogger - Returns a logger wrapped with pkg/stream fields
func NamedLogger(pkg, stream string) *logrus.Entry {
	return RootLogger.WithFields(logrus.Fields{
		"pkg":    pkg,
		"stream": stream,
	})
}

// GetRootAppDir - Get the CIPHERC2 app dir, default is: ~/.CIPHERC2/
func GetRootAppDir() string {

	value := os.Getenv(envVarName)

	var dir string
	if len(value) == 0 {
		user, _ := user.Current()
		dir = filepath.Join(user.HomeDir, ".CIPHERC2")
	} else {
		dir = value
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0700)
		if err != nil {
			panic("Cannot write to CIPHERC2 root dir")
		}
	}
	return dir
}

// GetLogDir - Return the log dir
func GetLogDir() string {
	rootDir := GetRootAppDir()
	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		err = os.MkdirAll(rootDir, 0700)
		if err != nil {
			panic(err)
		}
	}
	logDir := path.Join(rootDir, "logs")
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0700)
		if err != nil {
			panic(err)
		}
	}
	return logDir
}

// RootLogger - Returns the root logger
func rootLogger() *logrus.Logger {
	rootLogger := logrus.New()
	rootLogger.Formatter = &logrus.JSONFormatter{}
	jsonFilePath := filepath.Join(GetLogDir(), "CIPHERC2.json")
	jsonFile, err := os.OpenFile(jsonFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file %v", err))
	}
	rootLogger.Out = jsonFile
	rootLogger.SetLevel(logrus.DebugLevel)
	rootLogger.SetReportCaller(true)
	rootLogger.AddHook(NewTxtHook("root"))
	return rootLogger
}

// RootLogger - Returns the root logger
func txtLogger() *logrus.Logger {
	txtLogger := logrus.New()
	txtLogger.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}
	txtFilePath := filepath.Join(GetLogDir(), "CIPHERC2.log")
	txtFile, err := os.OpenFile(txtFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file %v", err))
	}
	txtLogger.Out = txtFile
	txtLogger.SetLevel(logrus.DebugLevel)
	return txtLogger
}

// TxtHook - Hook in a textual version of the logs
type TxtHook struct {
	Name   string
	logger *logrus.Logger
}

// NewTxtHook - returns a new txt hook
func NewTxtHook(name string) *TxtHook {
	hook := &TxtHook{
		Name:   name,
		logger: txtLogger(),
	}
	return hook
}

// Fire - Implements the fire method of the Logrus hook
func (hook *TxtHook) Fire(entry *logrus.Entry) error {
	if hook.logger == nil {
		return errors.New("no txt logger")
	}

	// Determine the caller (filename/line number)
	srcFile := "<no caller>"
	if entry.HasCaller() {
		CIPHERC2Index := strings.Index(entry.Caller.File, "CIPHERC2")
		srcFile = entry.Caller.File
		if CIPHERC2Index != -1 {
			srcFile = srcFile[CIPHERC2Index:]
		}
	}

	switch entry.Level {
	case logrus.PanicLevel:
		hook.logger.Panicf("[%s:%d] %s", srcFile, entry.Caller.Line, entry.Message)
	case logrus.FatalLevel:
		hook.logger.Fatalf("[%s:%d] %s", srcFile, entry.Caller.Line, entry.Message)
	case logrus.ErrorLevel:
		hook.logger.Errorf("[%s:%d] %s", srcFile, entry.Caller.Line, entry.Message)
	case logrus.WarnLevel:
		hook.logger.Warnf("[%s:%d] %s", srcFile, entry.Caller.Line, entry.Message)
	case logrus.InfoLevel:
		hook.logger.Infof("[%s:%d] %s", srcFile, entry.Caller.Line, entry.Message)
	case logrus.DebugLevel, logrus.TraceLevel:
		hook.logger.Debugf("[%s:%d] %s", srcFile, entry.Caller.Line, entry.Message)
	}

	return nil
}

// Levels - Hook all levels
func (hook *TxtHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// RootLogger - Returns the root logger
func stdoutLogger() *logrus.Logger {
	txtLogger := logrus.New()
	txtLogger.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}
	txtLogger.Out = os.Stdout
	txtLogger.SetLevel(logrus.DebugLevel)
	return txtLogger
}

// TxtHook - Hook in a textual version of the logs
type StdoutHook struct {
	Name   string
	logger *logrus.Logger
}

// NewTxtHook - returns a new txt hook
func NewStdoutHook(name string) *StdoutHook {
	hook := &StdoutHook{
		Name:   name,
		logger: stdoutLogger(),
	}
	return hook
}

// Fire - Implements the fire method of the Logrus hook
func (hook *StdoutHook) Fire(entry *logrus.Entry) error {
	if hook.logger == nil {
		return errors.New("no txt logger")
	}

	// Determine the caller (filename/line number)
	srcFile := "<no caller>"
	if entry.HasCaller() {
		CIPHERC2Index := strings.Index(entry.Caller.File, "CIPHERC2")
		srcFile = entry.Caller.File
		if CIPHERC2Index != -1 {
			srcFile = srcFile[CIPHERC2Index:]
		}
	}

	switch entry.Level {
	case logrus.PanicLevel:
		hook.logger.Panicf("[%s:%d] %s", srcFile, entry.Caller.Line, entry.Message)
	case logrus.FatalLevel:
		hook.logger.Fatalf("[%s:%d] %s", srcFile, entry.Caller.Line, entry.Message)
	case logrus.ErrorLevel:
		hook.logger.Errorf("[%s:%d] %s", srcFile, entry.Caller.Line, entry.Message)
	case logrus.WarnLevel:
		hook.logger.Warnf("[%s:%d] %s", srcFile, entry.Caller.Line, entry.Message)
	case logrus.InfoLevel:
		hook.logger.Infof("[%s:%d] %s", srcFile, entry.Caller.Line, entry.Message)
	case logrus.DebugLevel, logrus.TraceLevel:
		hook.logger.Debugf("[%s:%d] %s", srcFile, entry.Caller.Line, entry.Message)
	}

	return nil
}

// Levels - Hook all levels
func (hook *StdoutHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// LevelFrom - returns level from int
func LevelFrom(level int) logrus.Level {
	switch level {
	case 0:
		return logrus.PanicLevel
	case 1:
		return logrus.FatalLevel
	case 2:
		return logrus.ErrorLevel
	case 3:
		return logrus.WarnLevel
	case 4:
		return logrus.InfoLevel
	case 5:
		return logrus.DebugLevel
	case 6:
		return logrus.TraceLevel
	}
	return logrus.DebugLevel
}
