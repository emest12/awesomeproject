package panic

import (
	"fmt"
	"runtime/debug"
	"testing"
	"time"
)

// panic 会导致程序退出
func TestPanic1(t *testing.T) {
	panic("test")
	fmt.Println("done")
}

// panic 可以在defer中捕获，程序不会异常退出，但panic所在函数的代码不会继续执行
func TestPanic2(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("test")

	fmt.Println("done")
}

// panic 可以在defer中捕获，程序不会异常退出，panic所在函数的代码不会继续执行，但函数外可以继续执行
func TestPanic5(t *testing.T) {
	func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		panic("test")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("done")
}

// 协程panic会导致主协程退出
func TestPanic3(t *testing.T) {
	go func() {
		panic("test")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("done")
}

// 在协程中捕获panic，主协程不会退出
func TestPanic4(t *testing.T) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		panic("test")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("done")
}

// 因此一定要在协程中捕获panic，但是每次写defer都很麻烦
// WithRecover封装了defer的逻辑，如果panic，会将panic信息通过error返回出来
func WithRecover(f func() error) func() error {
	return func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				// 打印关键日志
				err = fmt.Errorf("panic in position: %s", string(debug.Stack()))
			}
		}()

		return f()
	}
}
func TestPanic6(t *testing.T) {
	go WithRecover(func() error {
		panic("test")
	})

	time.Sleep(1 * time.Second)
	fmt.Println("done")
}
