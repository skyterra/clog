package clog

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

const (
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

var (
	once          sync.Once
	globalOpts    []func(ctx context.Context) string
	defaultLogger *Clog
)

type Clog struct {
	ins      *log.Logger
	optional []func(ctx context.Context) string
}

// Debug logs a message at the DEBUG level.
func (cl *Clog) Debug(ctx context.Context, format string, v ...any) {
	cl.output(ctx, DEBUG, format, v...)
}

// Info logs a message at the INFO level.
func (cl *Clog) Info(ctx context.Context, format string, v ...any) {
	cl.output(ctx, INFO, format, v...)
}

// Warn logs a message at the WARN level.
func (cl *Clog) Warn(ctx context.Context, format string, v ...any) {
	cl.output(ctx, WARN, format, v...)
}

// Error logs a message at the ERROR level.
func (cl *Clog) Error(ctx context.Context, format string, v ...any) {
	cl.output(ctx, ERROR, format, v...)
}

// output formats and logs a message with the given level and context.
func (cl *Clog) output(ctx context.Context, level string, format string, v ...any) {
	var prefix = []string{level}
	for _, opt := range cl.optional {
		prefix = append(prefix, opt(ctx))
	}

	content := fmt.Sprintf(format, v...)
	cl.ins.Output(2, fmt.Sprintf("|%s|$ %s", strings.Join(prefix, "|"), content))
}

// initDefaultLogger initializes the default logger if it is not already initialized.
func initDefaultLogger() {
	if defaultLogger == nil {
		defaultLogger = NewClog()
	}
}

// Debug logs a message at the DEBUG level using the default logger.
func Debug(ctx context.Context, format string, v ...any) {
	once.Do(initDefaultLogger)
	defaultLogger.Debug(ctx, format, v...)
}

// Info logs a message at the INFO level using the default logger.
func Info(ctx context.Context, format string, v ...any) {
	once.Do(initDefaultLogger)
	defaultLogger.Info(ctx, format, v...)
}

// Warn logs a message at the WARN level using the default logger.
func Warn(ctx context.Context, format string, v ...any) {
	once.Do(initDefaultLogger)
	defaultLogger.Warn(ctx, format, v...)
}

// Error logs a message at the ERROR level using the default logger.
func Error(ctx context.Context, format string, v ...any) {
	once.Do(initDefaultLogger)
	defaultLogger.Error(ctx, format, v...)
}

// NewClog creates a new Clog instance with optional context functions.
func NewClog(opts ...func(ctx context.Context) string) *Clog {
	logger := &Clog{
		ins: log.New(os.Stdout, "", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
	}

	if len(opts) == 0 {
		opts = globalOpts
	}

	for _, opt := range opts {
		if opt != nil {
			logger.optional = append(logger.optional, opt)
		}
	}

	return logger
}

// SetDefaultOpts sets the global options for all Clog instances.
func SetDefaultOpts(opts ...func(ctx context.Context) string) {
	globalOpts = opts
}
