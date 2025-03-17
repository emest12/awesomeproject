package channel

import (
	"fmt"
	"golang.org/x/net/context"
	"testing"
	"time"
)

func TestChannel1(t *testing.T) {
	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()

	ch := make(chan int)
	go func() {
		for {
			ch <- 1
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			goto RETURN
		case <-ch:
			fmt.Println("receive")
		}
	}

RETURN:
	fmt.Println("done")
}
