package workers

import (
	"context"
	"fmt"
	"time"
)

func GetHtmlWorker(ctx context.Context, args ...interface{}) error {
	time.Sleep(5 * time.Second)
	fmt.Printf("Hello from get html worker: %s\n", args[0].(string))

	return nil
}
