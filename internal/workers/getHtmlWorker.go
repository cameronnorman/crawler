package workers

import (
	"context"
	"fmt"
	"os"
	"time"
)

func GetHtmlWorker(ctx context.Context, args ...interface{}) error {
	time.Sleep(5 * time.Second)
	name, _ := os.Hostname()
	fmt.Printf("[%s] Hello from get html worker: %s\n", name, args[0].(string))

	return nil
}
