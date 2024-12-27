package clog

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

type Clog struct {
	ins      *log.Logger
	optional []func(ctx context.Context) string
}

func (cl *Clog) Debug(ctx context.Context, format string, v ...any) {
	cl.output(ctx, DEBUG, format, v...)
}

func (cl *Clog) Info(ctx context.Context, format string, v ...any) {
	cl.output(ctx, INFO, format, v...)
}

func (cl *Clog) Warn(ctx context.Context, format string, v ...any) {
	cl.output(ctx, WARN, format, v...)
}

func (cl *Clog) Error(ctx context.Context, format string, v ...any) {
	cl.output(ctx, ERROR, format, v...)
}

func (cl *Clog) output(ctx context.Context, level string, format string, v ...any) {
	var prefix []string
	for _, opt := range cl.optional {
		prefix = append(prefix, opt(ctx))
	}
	prefix = append(prefix, level)
	content := fmt.Sprintf(format, v...)
	cl.ins.Output(2, fmt.Sprintf("|%s| %s", strings.Join(prefix, "|"), content))
}

func NewClog(opts ...func(ctx context.Context) string) *Clog {
	logger := &Clog{
		ins: log.New(os.Stdout, "", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
	}

	for _, opt := range opts {
		if opt != nil {
			logger.optional = append(logger.optional, opt)
		}
	}

	return logger
}
