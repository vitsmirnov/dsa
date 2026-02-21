package main

import (
	"fmt"
	"slices"
	"time"
)

type Integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

type Number interface{ Integer | float32 | float64 }

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

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Calculate(expression string) float64 { return 0 }

func Combinations(n int, k int) [][]int {
	// time: O(C(n,k)*k), space: O(n) / O(C(n,k)*k) ?

	if k > n {
		return nil
	}

	combinations := [][]int{}
	combination := make([]int, k)

	var combine func(pos int, start int)
	combine = func(pos int, start int) {
		if pos == k {
			combinations = append(combinations, slices.Clone(combination))
			return
		}

		for num := start; num < n; num++ {
			combination[pos] = num
			combine(pos+1, num+1)
		}
	}

	combine(0, 0)
	return combinations
}

func UniquePermutations(nums []int) [][]int {
	// k-permutations of N (partial permutation)

	numsLen := len(nums)
	permutations := [][]int{}

	var permute func(index int)
	permute = func(index int) {
		if index == numsLen {
			permutations = append(permutations, slices.Clone(nums))
			return
		}

		seen := map[int]bool{}
		for i := index; i < numsLen; i++ {
			if !seen[nums[i]] {
				nums[index], nums[i] = nums[i], nums[index]
				permute(index + 1)
				nums[index], nums[i] = nums[i], nums[index]
				seen[nums[i]] = true
			}
		}
	}

	permute(0)
	return permutations
}

func UniquePermutations0(nums []int) [][]int {
	// k-permutations of N (partial permutation)

	numsLen := len(nums)
	permutations := [][]int{}
	permutation := make([]int, numsLen)
	numsLeft := map[int]int{}
	for _, num := range nums {
		numsLeft[num]++
	}

	var permute func(index int)
	permute = func(index int) {
		if index == numsLen {
			permutations = append(permutations, slices.Clone(permutation))
			return
		}

		for num, left := range numsLeft {
			if left > 0 {
				permutation[index] = num
				numsLeft[num]--
				permute(index + 1)
				numsLeft[num]++
			}
		}
	}

	permute(0)
	return permutations
}

func Permutations(n int) [][]int {
	// time: O(n!*n), space: O(n) ?

	permutations := make([][]int, 0, Factorial(n))
	permutation := make([]int, n)
	for i := range permutation {
		permutation[i] = i
	}

	var permute func(pos int)
	permute = func(pos int) {
		if pos == n {
			permutations = append(permutations, slices.Clone(permutation))
			return
		}

		for i := pos; i < n; i++ {
			permutation[i], permutation[pos] = permutation[pos], permutation[i]
			permute(pos + 1)
			permutation[i], permutation[pos] = permutation[pos], permutation[i]
		}
	}

	permute(0)
	return permutations
}

func Permutations0(n int) [][]int {
	// time: O(n!*n), space: O(n) ?

	permutations := make([][]int, 0, Factorial(n))
	permutation := make([]int, n)
	usedIndices := make([]bool, n)

	var permute func(pos int)
	permute = func(pos int) {
		if pos == n {
			permutations = append(permutations, slices.Clone(permutation))
			return
		}

		for index, used := range usedIndices {
			if !used {
				permutation[pos] = index
				usedIndices[index] = true
				permute(pos + 1)
				usedIndices[index] = false
			}
		}
	}

	permute(0)
	return permutations
}

func NextPermutation(nums []int) {
	// time: O(n), space: O(1)

	numsLen := len(nums)
	i := numsLen - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		left, right := i+1, numsLen-1
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] > nums[i] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		nums[i], nums[right] = nums[right], nums[i]
	}
	for left, right := i+1, numsLen-1; left < right; left, right = left+1, right-1 {
		nums[left], nums[right] = nums[right], nums[left]
	}
}

func Factorial(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

func IsPowerOfTwo(x int) bool { return x > 0 && x&(x-1) == 0 }

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

func testCombinations() {
	strs := []string{"a", "b", "c"}
	for _, indices := range Combinations(len(strs), 2) {
		for _, index := range indices {
			fmt.Printf("%v ", strs[index])
		}
		fmt.Println()
	}

	return
	combinations := Combinations(10, 3)
	for _, combination := range combinations {
		fmt.Println(combination)
	}
}

func main() {
	// x := 3710
	// fmt.Printf("%b\n", x)
	// fmt.Println(CountBits(x))
	// fmt.Println(CountSetBits(x))
	// return

	n := 3
	for _, perm := range Permutations(n) {
		fmt.Println(perm)
	}
	fmt.Println()
	for _, perm := range Permutations0(n) {
		fmt.Println(perm)
	}
	// fmt.Println(len(Permutations(n)))
	// fmt.Println(len(Permutations0(n)))
	return

	// for i := range 10 {
	// 	fmt.Printf("%v: %v\n", i, Factorial(i))
	// }
	// return

	testCombinations()
	return

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
