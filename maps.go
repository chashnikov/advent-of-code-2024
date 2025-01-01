package main

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func Union[K comparable](s1 map[K]bool, s2 map[K]bool) map[K]bool {
	union := make(map[K]bool, len(s1)+len(s2))
	for key, value := range s1 {
		if value {
			union[key] = true
		}
	}
	for key, value := range s2 {
		if value {
			union[key] = true
		}
	}
	return union
}

func AddElement[K comparable](s map[K]bool, e K) map[K]bool {
	s[e] = true
	return s
}
