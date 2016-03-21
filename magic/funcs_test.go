package magic

import (
	"fmt"
	"testing"
)

func TestFuncSpell(t *testing.T) {
	msg := "hello world"
	hello := func() string {
		return msg
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
