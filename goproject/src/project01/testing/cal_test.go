package test

import (
	"testing"
)

func TestAddupper(t *testing.T){
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("执行错误， 期望值=%v 实际值=%v\n", 55, res)
	}
	t.Logf("执行正确")
}

func TestSub(t *testing.T){
	res := getSub(10, 3)
	if res != 7 {
		t.Fatalf("执行错误， 期望值=%v 实际值=%v\n", 7, res)
	}
	t.Logf("执行正确")
}
