package magic

import (
	"fmt"
	"testing"
)

func TestFuncSpell(t *testing.T) {
	hello := func() string {
		fmt.Println("hello")
		return "hello world"
	}
	h, err := NewFuncSpell(hello, "helo", "hello world")
	if err != nil {
		t.Error(err)
	}
	rst, err := h.Conjure()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rst)
}
