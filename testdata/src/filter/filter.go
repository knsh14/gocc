package filter

// Foo is not function declearation. so no warning.
var Foo = func(n int) int {
	return n * 10
}
