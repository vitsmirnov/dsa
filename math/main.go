package main

import (
	"fmt"
)

func Pow(base, exp int) int         { return 0 }
func PowMod(base, exp, mod int) int { return 0 }
func IsPrime(n int) bool {
	if n <= 2 || n%2 == 0 {
		return n == 2
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func getPrimeFactors(n int) []int {
	// time: O(sqrt(n)), space: O(1)

	primes := []int{}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			primes = append(primes, i)
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n > 1 {
		primes = append(primes, n)
	}
	return primes
}

// greatest common divisor
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func GCD2(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD2(b, a%b)
}

// least common multiple
func LCM(a, b int) int     { return Abs(a*b) / GCD(a, b) }
func DivCeil(a, b int) int { return (a + b - 1) / b }

type Integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

type Number interface{ Integer | float32 | float64 }

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Calculate(expression string) float64 { return 0 }

// bits
func CountSetBits(x int) int    { return 0 }
func CountBits(x int) int       { return 0 }
func IsPowerOfTwo(x int) bool   { return false }
func SetBit(x int, pos int) int { return x }
func OffBit(x int, pos int) int { return x }

func main() {
	// t := time.Now()
	// for i := range int(1e5) {
	// 	PrimeFactorization(i)
	// }
	// fmt.Println(time.Since(t))

	fmt.Println(getPrimeFactors(2 * 3 * 5 * 7 * 11 * 13))
	fmt.Println(getPrimeFactors(2))
	return

	// k := 0
	// for i := range int(1e6 + 1) {
	// 	if IsPrime(i) {
	// 		k++
	// 	}
	// }
	// fmt.Println(k)
}
