package main

import (
	"context"
	"github.com/skyterra/clog"
)

type requestId struct{}

func ReadRequestID(ctx context.Context) string {
	value := ctx.Value(requestId{})
	requestID, _ := value.(string)
	return requestID
}

func main() {
	ctx := context.WithValue(context.Background(), requestId{}, "aaa")
	logger := clog.NewClog(ReadRequestID)

	logger.Info(ctx, "hello, welcome to clog world. %s", "you will have a nice trip")

}
