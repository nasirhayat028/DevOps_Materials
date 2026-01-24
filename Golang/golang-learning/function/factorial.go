package main

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	// Example usage
	result := factorial(10)
	println("Factorial is:", result)
}
