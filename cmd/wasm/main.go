package main

//export Fibonacci
func Fibonacci(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return Fibonacci(i-1) + Fibonacci(i-2)
}

func main() {}
