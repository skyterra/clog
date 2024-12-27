# Clog - Custom Logger for Go
Clog is a custom logging package for Go that provides flexible logging functionality with different log levels and optional contextual data. It supports logging messages at different levels such as DEBUG, INFO, WARN, and ERROR. You can customize the logger with optional functions to append extra information, such as context or custom metadata, to the log messages.

# Features
- Log Levels: Support for DEBUG, INFO, WARN, and ERROR log levels.
- Customizable: Allow customization with optional functions that can append extra information (e.g., context data).
- Formatted Output: Support for formatted log messages.
- Microseconds & File Info: Display log messages with timestamps (including microseconds) and source file details.
  
# Installation
To install the Clog package, you can run:

```bash
go get github.com/skyterra/clog
```

# Usage
**Create a New Logger**  
You can create a new Clog instance using the NewClog function. Optionally, you can pass custom functions that add context to the logs.

```go
package main

import (
	"context"
	"fmt"
	"github.com/skyterra/clog"
)

func main() {
	logger := clog.NewClog()

	// Log messages at different levels
	logger.Debug(context.Background(), "This is a debug message")
	logger.Info(context.Background(), "This is an info message")
	logger.Warn(context.Background(), "This is a warning message")
	logger.Error(context.Background(), "This is an error message")
}
```

**Custom Optional Context** 
You can pass optional functions to customize the logger’s output. For example, you can add a function that adds a custom context.

```go

type requestId struct{}

func ReadRequestID(ctx context.Context) string {
    value := ctx.Value(requestId{})
    requestID, _ := value.(string)
    return requestID
}

func main() {
    ctx := context.WithValue(context.Background(), requestId{}, "xxx-request-id-xxx")
    logger := clog.NewClog(ReadRequestID)

    logger.Info(ctx, "hello, welcome to clog world. %s", "you will have a nice trip")
}
```

**Set Default Options**
You can set global options that will be used by all instances of Clog by calling SetDefaultOpts.
```go
type requestId struct{}

func ReadRequestID(ctx context.Context) string {
    value := ctx.Value(requestId{})
    requestID, _ := value.(string)
    return requestID
}

func main() {
    ctx := context.WithValue(context.Background(), requestId{}, "xxx-request-id-xxx")
    clog.SetDefaultOpts(ReadRequestID)

    // All loggers will use the default options
    logger := clog.NewClog()

    // Log message will include custom context
    logger.Debug(ctx, "Debug message with default context")
}

```
