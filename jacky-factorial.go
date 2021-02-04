package main

import "fmt"

func factorial(number int) (results []int) {
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			results = append(results, i)
		}
	}

	return results
}

func main() {
	num := 262144

	var total int
	for i := 1; i <= num; i++ {
		results := factorial(i)
		if len(results) == 6 {
			total++
		}
	}

	fmt.Println(total)
}
