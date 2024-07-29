package main

import "testing"

// go test -bench=. -benchmem -memprofile=mem.out -cpuprofile=cpu.out
// go tool pprof -alloc_space mem.out
// go test -v -run . -bench=BenchmarkAppendEmpty -benchmem

func TestAppendNil(t *testing.T) {
	type StringSlice []string

	aSlice := StringSlice{"a", "b", "c"}
	result := append(StringSlice(nil), aSlice...)

	if len(result) != 3 {
		t.Errorf("expected len(result) == 3 got %d", len(result))
	}

	if result[0] != "a" || result[1] != "b" || result[2] != "c" {
		t.Errorf("expected result == [a b c], got %v", result)
	}
}

func TestAppendEmpty(t *testing.T) {
	type StringSlice []string

	aSlice := StringSlice{"a", "b", "c"}
	result := append(StringSlice{}, aSlice...)

	if len(result) != 3 {
		t.Errorf("expected len(result) == 3, got %d", len(result))
	}
	if result[0] != "a" || result[1] != "b" || result[2] != "c" {
		t.Errorf("expected result == [a b c], got %v", result)
	}
}

func TestAppendMultiple(t *testing.T) {
	type StringSlice []string

	aSlice := StringSlice{"a", "b", "c"}
	bSlice := StringSlice{"d", "e", "f"}
	cSlice := StringSlice{"g", "h", "i"}

	result := append(StringSlice{}, aSlice...)
	result = append(result, bSlice...)
	result = append(result, cSlice...)

	if len(result) != 9 {
		t.Errorf("expected len(result) == 9, got %d", len(result))
	}
	if result[0] != "a" || result[1] != "b" || result[2] != "c" ||
		result[3] != "d" || result[4] != "e" || result[5] != "f" ||
		result[6] != "g" || result[7] != "h" || result[8] != "i" {
		t.Errorf("expected result == [a b c d e f g h i], got %v", result)
	}
}

func TestAppendNilSlice(t *testing.T) {
	type StringSlice []string

	var aSlice StringSlice

	result := append(StringSlice{}, aSlice...)

	if len(result) != 0 {
		t.Errorf("expected len(result) == 0, got %d", len(result))
	}
}

type StringSlice []string

func BenchmarkAppendNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		aSlice := StringSlice{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
		result := append(StringSlice(nil), aSlice...)
		_ = result
	}
}

func BenchmarkAppendEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		aSlice := StringSlice{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
		result := append(StringSlice{}, aSlice...)
		_ = result
	}
}

func BenchmarkAppendMultipleNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		aSlice := StringSlice{"a", "b", "c"}
		bSlice := StringSlice{"d", "e", "f"}
		cSlice := StringSlice{"g", "h", "i"}
		result := append(StringSlice(nil), aSlice...)
		result = append(result, bSlice...)
		result = append(result, cSlice...)
	}
}

func BenchmarkAppendMultipleEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		aSlice := StringSlice{"a", "b", "c"}
		bSlice := StringSlice{"d", "e", "f"}
		cSlice := StringSlice{"g", "h", "i"}
		result := append(StringSlice{}, aSlice...)
		result = append(result, bSlice...)
		result = append(result, cSlice...)
	}
}
