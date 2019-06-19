package golog

import (
	"context"
	"log"
)

// Verbose level
type Verbose int

const (
	// DebugPrefix text to prefix to each log entry generated by the Logger
	DebugPrefix = "DEBUG: "
	// InfoPrefix text to prefix to each log entry generated by the Logger
	InfoPrefix = "INFO: "
	// WarnPrefix text to prefix to each log entry generated by the Logger
	WarnPrefix = "WARN: "
	// ErrorPrefix text to prefix to each log entry generated by the Logger
	ErrorPrefix = "ERROR: "
	// FatalPrefix text to prefix to each log entry generated by the Logger
	FatalPrefix = "FATAL: "

	// DefaultFlags define default text to prefix to each log entry generated by the Logger
	DefaultFlags = log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC

	// Disabled verbose level disables logger messages
	Disabled Verbose = -1
	// Critical verbose level
	Critical Verbose = iota
	// Error verbose level
	Error
	// Warning verbose level
	Warning
	// Info verbose level
	Info
	// Debug verbose level
	Debug
)

type (
	// Logger interface
	Logger interface {
		Debug(ctx context.Context, format string, args ...interface{})
		Info(ctx context.Context, format string, args ...interface{})
		Warning(ctx context.Context, format string, args ...interface{})
		Error(ctx context.Context, format string, args ...interface{})
		Critical(ctx context.Context, format string, args ...interface{})

		// Override DefaultFlags value on a Logger instance
		SetFlags(flag int)
	}

	loggerFactory func(level Verbose) Logger
)

// New returns a Logger.
var New loggerFactory
