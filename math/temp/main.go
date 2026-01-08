package main

import "fmt"

func main() {
	// for n := range 50 {
	// 	if IsPrime(n) {
	// 		fmt.Println(n)
	// 	}
	// }
	// return

	n := 150
	fmt.Println(n)
	k := 1
	for _, factor := range GetFactors(n) {
		fmt.Println(factor)
		k *= factor
	}
	fmt.Println(k)

	// fmt.Println(IsPrime(1_000_000_007))
}

func GetFactors(n int) []int {
	// if n <= 1 {
	// 	return []int{}
	// }

	factors := []int{}
	for prime := 2; prime*prime <= n; prime++ {
		for ; n%prime == 0; n /= prime {
			factors = append(factors, prime)
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}

func IsPrime(n int) bool {
	// n = sqrt(n) * sqrt(n) = a * b
	// if a <= b then a <= sqrt(n) and b >= sqrt(n)
	// so it is enough to try deviders which <= sqrt(n)

	// if n <= 1 {
	// 	return false
	// }

	for d := 2; d*d <= n; d++ {
		if n%d == 0 {
			return false
		}
	}
	return n > 1
}
