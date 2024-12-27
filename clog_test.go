package clog

import (
	"bytes"
	"context"
	"log"
	"strings"
	"testing"
)

type requestId struct{}

func setupClog() (*Clog, *bytes.Buffer) {
	var buf bytes.Buffer
	logger := log.New(&buf, "", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	clog := &Clog{
		ins: logger,
		optional: []func(ctx context.Context) string{
			func(ctx context.Context) string {
				return ctx.Value(requestId{}).(string)
			},
		},
	}
	return clog, &buf
}

func TestClog_Debug(t *testing.T) {
	clog, buf := setupClog()
	ctx := context.WithValue(context.Background(), requestId{}, "test-request-id")

	clog.Debug(ctx, "debug message")
	logOutput := buf.String()

	if !strings.Contains(logOutput, DEBUG) {
		t.Errorf("expected log level %s, got %s", DEBUG, logOutput)
	}
	if !strings.Contains(logOutput, "test-request-id") {
		t.Errorf("expected request ID %s, got %s", "test-request-id", logOutput)
	}
	if !strings.Contains(logOutput, "debug message") {
		t.Errorf("expected message %s, got %s", "debug message", logOutput)
	}
}

func TestClog_Info(t *testing.T) {
	clog, buf := setupClog()
	ctx := context.WithValue(context.Background(), requestId{}, "test-request-id")

	clog.Info(ctx, "info message")
	logOutput := buf.String()

	if !strings.Contains(logOutput, INFO) {
		t.Errorf("expected log level %s, got %s", INFO, logOutput)
	}
	if !strings.Contains(logOutput, "test-request-id") {
		t.Errorf("expected request ID %s, got %s", "test-request-id", logOutput)
	}
	if !strings.Contains(logOutput, "info message") {
		t.Errorf("expected message %s, got %s", "info message", logOutput)
	}
}

func TestClog_Warn(t *testing.T) {
	clog, buf := setupClog()
	ctx := context.WithValue(context.Background(), requestId{}, "test-request-id")

	clog.Warn(ctx, "warn message")
	logOutput := buf.String()

	if !strings.Contains(logOutput, WARN) {
		t.Errorf("expected log level %s, got %s", WARN, logOutput)
	}
	if !strings.Contains(logOutput, "test-request-id") {
		t.Errorf("expected request ID %s, got %s", "test-request-id", logOutput)
	}
	if !strings.Contains(logOutput, "warn message") {
		t.Errorf("expected message %s, got %s", "warn message", logOutput)
	}
}

func TestClog_Error(t *testing.T) {
	clog, buf := setupClog()
	ctx := context.WithValue(context.Background(), requestId{}, "test-request-id")

	clog.Error(ctx, "error message")
	logOutput := buf.String()

	if !strings.Contains(logOutput, ERROR) {
		t.Errorf("expected log level %s, got %s", ERROR, logOutput)
	}
	if !strings.Contains(logOutput, "test-request-id") {
		t.Errorf("expected request ID %s, got %s", "test-request-id", logOutput)
	}
	if !strings.Contains(logOutput, "error message") {
		t.Errorf("expected message %s, got %s", "error message", logOutput)
	}
}
func TestNewClog(t *testing.T) {
	// Test with no options
	clog := NewClog()
	if clog == nil {
		t.Error("expected clog to be non-nil")
	}
	if len(clog.optional) != 0 {
		t.Errorf("expected no optional functions, got %d", len(clog.optional))
	}

	// Test with options
	opt := func(ctx context.Context) string {
		return "test-option"
	}
	clog = NewClog(opt)
	if len(clog.optional) != 1 {
		t.Errorf("expected 1 optional function, got %d", len(clog.optional))
	}
	if clog.optional[0](context.Background()) != "test-option" {
		t.Errorf("expected optional function to return 'test-option'")
	}
}

func TestSetDefaultOpts(t *testing.T) {
	// Test setting global options
	opt := func(ctx context.Context) string {
		return "global-option"
	}
	SetDefaultOpts(opt)

	clog := NewClog()
	if len(clog.optional) != 1 {
		t.Errorf("expected 1 global optional function, got %d", len(clog.optional))
	}
	if clog.optional[0](context.Background()) != "global-option" {
		t.Errorf("expected global optional function to return 'global-option'")
	}

	// Reset global options
	SetDefaultOpts()
	clog = NewClog()
	if len(clog.optional) != 0 {
		t.Errorf("expected no global optional functions, got %d", len(clog.optional))
	}
}
