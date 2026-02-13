package main

import (
	"fmt"
	"time"
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

func GetPrimeFactors(n int) []int {
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

// sieve of Eratosthenes
func Sieve(n int) []bool {
	// time: O(n log log n), space: O(n)

	primeTags := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primeTags[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if primeTags[i] {
			for j := i * i; j <= n; j += i {
				primeTags[j] = false
			}
		}
	}
	return primeTags
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

func testSieve() {
	const maxNum int = 1e5
	const loopCount = 1000

	countPrimes := func(n int) int {
		count := 0
		for i := range n + 1 {
			if IsPrime(i) {
				count++
			}
		}
		return count
	}
	countPrimesWithSieve := func(n int) int {
		count := 0
		for _, isPrime := range Sieve(n) {
			if isPrime {
				count++
			}
		}
		return count
	}

	t := time.Now()
	var d1, d2 time.Duration
	for range 1 { //loopCount {
		n := int(1e7) //rand.IntN(maxNum)
		_t := time.Now()
		count1 := countPrimes(n)
		d1 += time.Since(_t)
		_t = time.Now()
		count2 := countPrimesWithSieve(n)
		d2 += time.Since(_t)
		if count1 != count2 {
			fmt.Printf("Count primes failed for %v: %v != %v (sieve)\n", n, count1, count2)
		}
	}
	fmt.Printf("Count primes test time: %v (%v, %v (sieve))\n", time.Since(t), d1, d2)
}

func main() {
	testSieve()
	return

	// t := time.Now()
	// for i := range int(1e5) {
	// 	PrimeFactorization(i)
	// }
	// fmt.Println(time.Since(t))

	fmt.Println(GetPrimeFactors(2 * 3 * 5 * 7 * 11 * 13))
	fmt.Println(GetPrimeFactors(2))
	return

	// k := 0
	// for i := range int(1e6 + 1) {
	// 	if IsPrime(i) {
	// 		k++
	// 	}
	// }
	// fmt.Println(k)
}
