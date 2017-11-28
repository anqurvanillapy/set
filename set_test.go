package set

import "testing"

const NumBenchmem = 10000000

func BenchmarkBuiltinDelete(b *testing.B) {
	s := make(map[interface{}]bool)
	for i := 0; i < NumBenchmem; i++ {
		s[i] = true
	}
	for i := 0; i < NumBenchmem; i++ {
		delete(s, i)
	}
}

func BenchmarkAssignZeroValue(b *testing.B) {
	s := make(map[interface{}]bool)
	for i := 0; i < NumBenchmem; i++ {
		s[i] = true
	}
	for i := 0; i < NumBenchmem; i++ {
		s[i] = false
	}
}

func BenchmarkBoolMap(b *testing.B) {
	s := make(map[interface{}]bool)
	for i := 0; i < NumBenchmem; i++ {
		s[i] = true
	}
	for i := 0; i < NumBenchmem; i++ {
		delete(s, i)
	}
}

func BenchmarkStructMap(b *testing.B) {
	s := make(map[interface{}]struct{})
	for i := 0; i < NumBenchmem; i++ {
		s[i] = struct{}{}
	}
	for i := 0; i < NumBenchmem; i++ {
		delete(s, i)
	}
}

func TestUSetAdd(t *testing.T) {
	s := NewUSet()
	s.Add(42)
	if _, ok := s[42]; !ok {
		t.Error("USet Add error")
	}
	if exists := s.Add(42); !exists {
		t.Error("USet Add exists error")
	}
}

func TestUSetHas(t *testing.T) {
	s := NewUSet()
	s.Add(42)
	if !s.Has(42) {
		t.Error("USet Has error")
	}
}

func TestUSetDelete(t *testing.T) {
	s := NewUSet()
	s.Add(42)
	s.Delete(42)
	if s.Has(42) {
		t.Error("USet Delete error")
	}
}
