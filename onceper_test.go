package onceper_test

import (
	"github.com/spartan0x117/onceper"
	"testing"
)

func TestOncePerString(t *testing.T) {
	o := onceper.New[string]()
	var count int
	f := func() {
		count++
	}
	o.Do("foo", f)
	o.Do("foo", f)
	o.Do("bar", f)
	o.Do("bar", f)
	if count != 2 {
		t.Errorf("count = %d, want 2", count)
	}
}

func TestOncePerInt(t *testing.T) {
	o := onceper.New[int]()
	var count int
	f := func() {
		count++
	}
	o.Do(1, f)
	o.Do(1, f)
	o.Do(2, f)
	o.Do(2, f)
	if count != 2 {
		t.Errorf("count = %d, want 2", count)
	}
}

func TestOncePerMixed(t *testing.T) {
	o := onceper.New[interface{}]()

	var count int
	f := func() {
		count++
	}
	o.Do("foo", f)
	o.Do("foo", f)
	o.Do(1, f)
	o.Do(1, f)
	if count != 2 {
		t.Errorf("count = %d, want 2", count)
	}
}

func TestOncePerStruct(t *testing.T) {
	type S struct {
		a int
		b string
	}
	o := onceper.New[S]()
	var count int
	f := func() {
		count++
	}
	o.Do(S{1, "foo"}, f)
	o.Do(S{1, "foo"}, f)
	o.Do(S{2, "bar"}, f)
	o.Do(S{2, "bar"}, f)
	if count != 2 {
		t.Errorf("count = %d, want 2", count)
	}

	o.Do(S{1, "bar"}, f)
	o.Do(S{1, "bar"}, f)
	if count != 3 {
		t.Errorf("count = %d, want 3", count)
	}
}

func TestOncePerDoWith(t *testing.T) {
	o := onceper.New[string]()
	seen := make(map[string]int)
	f := func(s string) {
		seen[s] = seen[s] + 1
	}
	o.DoWith("foo", f)
	o.DoWith("foo", f)
	o.DoWith("bar", f)
	o.DoWith("bar", f)
	o.DoWith("baz", f)
	o.DoWith("baz", f)

	for k, v := range seen {
		if v != 1 {
			t.Errorf("count[%s] = %d, want 1", k, v)
		}
	}
}

func BenchmarkOncePer(b *testing.B) {
	o := onceper.New[int]()
	var count int
	f := func() {
		count++
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			o.Do(1, f)
		}
	})
}