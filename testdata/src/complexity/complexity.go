package complexity

func SimpleFunction(n int) {
	println(n)
}

func ComplexFunction(n int) { // want `function ComplexFunction complexity=10`
	if n == 0 {
		println("zero")
	}
	if n == 1 {
		println("One")
	}
	if n == 2 {
		println("two")
	}
	if n == 3 {
		println("three")
	}
	if n == 4 {
		println("four")
	}
	if n == 5 {
		println("five")
	}
	if n == 6 {
		println("six")
	}
	if n == 7 {
		println("seven")
	}
	if n == 8 {
		println("eight")
	}
	println("not between zero and eight")
}
