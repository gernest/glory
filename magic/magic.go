package magic

import (
	"reflect"
	"sync"
)

// Book is a collection os spells, where you can learn about spells and pracice
// the art of magic
type Book struct {
	spells map[string]Spell
	mu     sync.RWMutex
}

// Spell is an interface for conjuring magic
type Spell interface {
	Name() string
	Effect() string
	NumIn() int
	NumOut() int
	Conjure(...interface{}) (Result, error)
}

type Result []interface{}

func ToResult(src []reflect.Value) Result {
	var v Result
	for _, val := range src {
		v = append(v, val.Interface())
	}
	return v
}

func (b *Book) AddSpell(s Spell) {
	b.mu.Lock()
	b.spells[s.Name()] = s
	b.mu.Unlock()
}
