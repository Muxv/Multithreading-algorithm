package mulalg

var res = 0

func Fib(n int) int {
	if n <= 1 {
		return n
	}

	x := Fib(n - 1)
	y := Fib(n - 2)
	return x + y
}

func FibParallel(n int) int {
	if n <= 1 {
		return n
	}
	ch := make(chan int)

	go func(in int, out chan<- int) {
		out <- FibParallel(in)
	}(n-1, ch)
	return <-ch + Fib(n-2)
}
