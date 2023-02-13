package utils

import (
	"testing"
)

func Test_addUpper(t *testing.T) {
	res := AddUpper(10)
	if res != 45 {
		t.Fatalf("实行错误:%v", res)
	}

	t.Logf("执行正确")
}
