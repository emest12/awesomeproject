package channel

import (
	"fmt"
	"golang.org/x/net/context"
	"sync"
	"testing"
	"time"
)

// context.WithTimeout and select case
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

// waitgroup
func TestChannel2(t *testing.T) {
	ch1 := make(chan context.Context, 1)
	ch2 := make(chan context.Context, 1)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		ch1 <- context.Background()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		ch2 <- context.Background()
	}()

	// 等两个协程都执行完
	wg.Wait()
	for {
		select {
		case <-ch1:
			t.Log("received ch1")
		case <-ch2:
			t.Log("received ch2")
		case <-time.After(1 * time.Second):
			goto RETURN
		}
	}
RETURN:
	fmt.Println("done")
}

// for range
func TestChannel3(t *testing.T) {
	ch := make(chan context.Context)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			ctx := context.WithValue(context.Background(), "key", i)
			ch <- ctx
		}
		close(ch)
	}()

	// channel不close for range会一直阻塞尝试从channel中取数据
	for c := range ch {
		t.Log(c)
	}
	// channel close for range会退出
	fmt.Println("done")
}

// c, open := <-ch ch未close时，会阻塞，c为channel中数据，open为true； ch close时，不会阻塞，c为channel中数据的零值，open为false
func TestChannel4(t *testing.T) {
	ch := make(chan context.Context)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			ctx := context.WithValue(context.Background(), "key", i)
			ch <- ctx
		}
		close(ch)
	}()

	for {
		c, open := <-ch
		t.Log(c)
		t.Log(open)
	}
}
