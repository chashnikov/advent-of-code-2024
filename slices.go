package main

func Map[S ~[]A, A any, B any](s S, mapping func(A) B) []B {
	result := make([]B, len(s))
	for i, a := range s {
		result[i] = mapping(a)
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
