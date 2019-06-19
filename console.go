// +build !appengine,!appenginevm

package golog

import (
	"context"
	"log"
	"os"
)

type consoleLogger struct {
	verboseLevel Verbose

	debug, info, warning, error, critical *log.Logger
}

// Terminal colours.
const (
	CLR_0 = "\x1b[30;1m"
	CLR_R = "\x1b[31;1m"
	CLR_G = "\x1b[32;1m"
	CLR_Y = "\x1b[33;1m"
	CLR_B = "\x1b[34;1m"
	CLR_M = "\x1b[35;1m"
	CLR_C = "\x1b[36;1m"
	CLR_W = "\x1b[37;1m"
	CLR_N = "\x1b[0m"
)

// NewConsoleLogger returns a Logger that writes to the console.
func NewConsoleLogger(level Verbose) Logger {
	return &consoleLogger{
		level,
		log.New(os.Stdout, CLR_W+DebugPrefix, DefaultFlags),
		log.New(os.Stdout, CLR_C+InfoPrefix, DefaultFlags),
		log.New(os.Stdout, CLR_Y+WarnPrefix, DefaultFlags),
		log.New(os.Stdout, CLR_R+ErrorPrefix, DefaultFlags),
		log.New(os.Stdout, CLR_R+FatalPrefix, DefaultFlags),
	}
}

func (logger *consoleLogger) SetFlags(flag int) {
	logger.debug.SetFlags(flag)
	logger.info.SetFlags(flag)
	logger.warning.SetFlags(flag)
	logger.error.SetFlags(flag)
	logger.critical.SetFlags(flag)
}

func (logger *consoleLogger) Debug(ctx context.Context, format string, args ...interface{}) {
	if logger.verboseLevel >= Debug {
		logger.debug.Printf(format, args...)
	}
}

func (logger *consoleLogger) Info(ctx context.Context, format string, args ...interface{}) {
	if logger.verboseLevel >= Info {
		logger.info.Printf(format, args...)
	}
}

func (logger *consoleLogger) Warning(ctx context.Context, format string, args ...interface{}) {
	if logger.verboseLevel >= Warning {
		logger.warning.Printf(format, args...)
	}
}

func (logger *consoleLogger) Error(ctx context.Context, format string, args ...interface{}) {
	if logger.verboseLevel >= Error {
		logger.error.Printf(format, args...)
	}
}

func (logger *consoleLogger) Critical(ctx context.Context, format string, args ...interface{}) {
	if logger.verboseLevel >= Critical {
		logger.critical.Printf(format, args...)
	}
}

func init() {
	New = NewConsoleLogger
}
