package util

// Associative converts a slice to a map using a given function.
//
// The convert function f is called on each element of slice s.
// The key and value returned by f are used to populate the map.
//
// Should two slice item returns the same key, the last one is used.
func Associative[E any, S ~[]E, K comparable, V any](s S, f func(E) (K, V)) map[K]V {
	m := make(map[K]V, len(s))

	for _, e := range s {
		k, v := f(e)
		m[k] = v
	}

	return m
}

// ToVis converts a slice to a map, where each slice element is mapped to true.
func ToVis[K comparable, S ~[]K](s S) map[K]bool {
	m := make(map[K]bool, len(s))

	for _, e := range s {
		m[e] = true
	}

	return m
}
