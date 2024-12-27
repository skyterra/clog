package main

import (
	"context"
	"github.com/skyterra/clog"
)

func main() {
	logger := clog.NewClog()
	logger.Info(context.Background(), "hello world:%s", "haha")

}
