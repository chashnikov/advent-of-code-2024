package main

import "cmp"

func Map[S ~[]A, A any, B any](s S, mapping func(A) B) []B {
	result := make([]B, len(s))
	for i, a := range s {
		result[i] = mapping(a)
	}
	return result
}

func Flatten[A any](s [][]A) []A {
	result := make([]A, 0)
	for _, a := range s {
		result = append(result, a...)
	}
	return result
}

func Filter[S ~[]A, A any](s S, predicate func(A) bool) []A {
	result := make([]A, 0)
	for _, a := range s {
		if predicate(a) {
			result = append(result, a)
		}
	}
	return result
}

func FilterWithIndex[S ~[]A, A any](s S, predicate func(int, A) bool) []A {
	result := make([]A, 0)
	for i, a := range s {
		if predicate(i, a) {
			result = append(result, a)
		}
	}
	return result
}

func ZipMap[S ~[]A, A any, T ~[]B, B any, C any](s S, t T, mapping func(A, B) C) []C {
	result := make([]C, min(len(s), len(t)))
	for i := 0; i < len(result); i++ {
		result[i] = mapping(s[i], t[i])
	}
	return result
}

func All[S ~[]A, A any](s S, predicate func(A) bool) bool {
	for _, a := range s {
		if !predicate(a) {
			return false
		}
	}
	return true
}

func Any[S ~[]A, A any](s S, predicate func(A) bool) bool {
	for _, a := range s {
		if predicate(a) {
			return true
		}
	}
	return false
}

func Sum(s []int) int {
	result := 0
	for _, a := range s {
		result += a
	}
	return result
}

func Sum64(s []int64) int64 {
	var result int64 = 0
	for _, a := range s {
		result += a
	}
	return result
}

func MinIndex[S ~[]E, E cmp.Ordered](x S) int {
	if len(x) < 1 {
		panic("slices.Min: empty list")
	}
	m := 0
	for i := 1; i < len(x); i++ {
		if x[i] < x[m] {
			m = i
		}
	}
	return m
}

func Count[S ~[]A, A any](s S, predicate func(A) bool) int {
	result := 0
	for _, a := range s {
		if predicate(a) {
			result++
		}
	}
	return result
}

func CreateInt2dSlice(size1 int, size2 int, value int) [][]int {
	result := make([][]int, size1)
	for i := 0; i < size1; i++ {
		result[i] = make([]int, size2)
		for j := 0; j < size2; j++ {
			result[i][j] = value
		}
	}
	return result
}

func CreateBool2dSlice(size1 int, size2 int) [][]bool {
	result := make([][]bool, size1)
	for i := 0; i < size1; i++ {
		result[i] = make([]bool, size2)
	}
	return result
}

func CreateSlice[A any](size int, value A) []A {
	result := make([]A, size)
	for i := 0; i < size; i++ {
		result[i] = value
	}
	return result
}
