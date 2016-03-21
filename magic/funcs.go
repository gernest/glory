package magic

import (
	"errors"
	"fmt"
	"reflect"
)

type FuncSpell struct {
	v    reflect.Value
	n    string
	desc string
	in   int
	out  int
}

func NewFuncSpell(v interface{}, name, desc string) (Spell, error) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Func {
		return nil, errors.New("non func argument")
	}
	return &FuncSpell{
		v:    val,
		n:    "",
		desc: "",
		in:   val.Type().NumIn(),
		out:  val.Type().NumOut(),
	}, nil
}

func (f *FuncSpell) Name() string {
	return f.n
}

func (f *FuncSpell) Effect() string {
	return f.desc
}
func (f *FuncSpell) NumIn() int {
	return f.in
}

func (f *FuncSpell) NumOut() int {
	return f.in
}

func (f *FuncSpell) Conjure(a ...interface{}) (Result, error) {
	sixe := len(a)
	if sixe == 0 {
		if f.NumIn() == 0 && f.NumOut() == 0 {
			f.v.Call(nil)
			return nil, nil
		}
		if f.NumIn() == 0 && f.NumOut() > 0 {
			return ToResult(f.v.Call(nil)), nil
		}
		return nil, fmt.Errorf("wrong number of arguments, expect %d got 0", f.in)
	}
	var args []reflect.Value
	for _, v := range a {
		val := reflect.ValueOf(v)
		if !val.IsValid() {
			return nil, errors.New("wrong arguments")
		}
		args = append(args, val)
	}
	if f.NumOut() > 0 {
		return ToResult(f.v.Call(args)), nil
	}
	f.v.Call(args)
	return nil, nil
}
