package main

import "fmt"

func main() {
	n := 7
	series := fibonacciSeries(n)
	fmt.Println("Fibonacci series:", series)

	primeNumber := primeNumber(n)
	fmt.Println("primeNumber", primeNumber)

}
