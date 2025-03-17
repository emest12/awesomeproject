package main

import (
	"fmt"
	"golang.org/x/net/context"
	"testing"
	"time"
)

func TestContext1(t *testing.T) {
	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, 10*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(11 * time.Millisecond):
		fmt.Println("oversleep")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
