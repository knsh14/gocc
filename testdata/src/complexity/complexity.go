package complexity

func SimpleFunction(n int) {
	println(n)
}

func ComplexFunction(n int) { // want `func ComplexFunction complexity=10`
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

type Foo struct {
}

func (f Foo) Bar(n int) { // want `func \(f Foo\) Bar complexity=10`
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

func (f *Foo) Buzz(n int) { // want `func \(f \*Foo\) Buzz complexity=10`
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
