package util

import "reflect"

// SliceND creates and initializes a N dimensional slice.
//
// Return value should have type [][]...[]T. Call a type assertion on it to convert.
func SliceND[T any](size0 int, sizeRest ...int) any {
	if len(sizeRest) == 0 {
		return make([]T, size0)
	}

	t := reflect.SliceOf(reflect.TypeOf(*new(T)))
	for i := 0; i < len(sizeRest); i++ {
		t = reflect.SliceOf(t)
	}
	slice := reflect.MakeSlice(t, size0, size0)

	for i := 0; i < size0; i++ {
		slice.Index(i).Set(reflect.ValueOf(SliceND[T](sizeRest[0], sizeRest[1:]...)))
	}

	return slice.Interface()
}

// Diff returns the number of different elements between two slices, or -1 if they
// have different lengths.
//
// Comparisons are done with !=, not deep compare.
func Diff[S ~[]E, E comparable](l, r S) int {
	if len(l) != len(r) {
		return -1
	}
	diff := 0
	for i := 0; i < len(l); i++ {
		if l[i] != r[i] {
			diff++
		}
	}
	return diff
}

// Reduce a slice to a single value using a given function.
//
// The reduce function f is called on each element from slice s.
// It receives two arguments: the first is the current reduced value, the second is the next element from slice.
// For the first element, f(init, s[0]) is called.
func Reduce[S ~[]E, E any, R any](s S, f func(R, E) R, init R) R {
	for _, e := range s {
		init = f(init, e)
	}
	return init
}

// ReduceIndex reduces a slice to a single value using a given function.
//
// The 2nd argument to the reduce function f is the index of the current element.
func ReduceIndex[S ~[]E, E any, R any](s S, f func(R, int, E) R, init R) R {
	for i, e := range s {
		init = f(init, i, e)
	}
	return init
}

// ReduceMap reduces a map to a single value using a given function.
func ReduceMap[K comparable, V any, R any](m map[K]V, f func(R, K, V) R, init R) R {
	for k, v := range m {
		init = f(init, k, v)
	}
	return init
}
