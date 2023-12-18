package util

// Must is used to handle errors neatly.
//
// If err is not nil, this function panics with the error. Otherwise, it returns
// the first argument.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

// Assert panics with the given message if cond is false.
func Assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}
