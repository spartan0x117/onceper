package onceper

import (
	"sync"
)

type OncePer[T comparable] struct {
	mu   sync.Mutex
	done map[T]interface{}
}

// New returns a new OncePer[T] instance.
func New[T comparable]() *OncePer[T] {
	return &OncePer[T]{done: make(map[T]interface{})}
}

// Do calls the function f if and only if Do is being called for the
// first time with this key.
func (o *OncePer[T]) Do(key T, f func()) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if _, ok := o.done[key]; !ok {
		f()
		o.done[key] = nil
	}
}

// DoWith calls the function f if and only if DoWith is being called for the
// first time with this key, passing the key as an argument to f.
func (o *OncePer[T]) DoWith(key T, f func(T)) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if _, ok := o.done[key]; !ok {
		f(key)
		o.done[key] = nil
	}
}