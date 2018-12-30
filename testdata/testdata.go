package testdata

func SimpleFunction(n int) {
	println(n)
}

func ComplexFunction(n int) { // want `function ComplexFunction complexity=10`
	if n > 0 {
		println("more than zero")
		if n > 1 {
			println("more than One")
			if n > 2 {
				println("more than two")
				if n > 3 {
					println("more than three")
					if n > 4 {
						println("more than four")
					}
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			println(i * j)
		}
	}
	for k := 0; k < n; k++ {
		for l := k; l < n; l++ {
			println(k * l)
		}
	}
}
