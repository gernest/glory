package magic

import "testing"

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
	first := rst.First().(string)
	if first != msg {
		t.Errorf("expecetd %s got %s", msg, first)
	}
}
