package leetcode

import (
	"testing"
)

func TestEvalRPN(t *testing.T) {
	got := evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	_ = got
}
