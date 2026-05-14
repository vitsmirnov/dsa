package main

import (
	"fmt"
	"time"
)

// Knuth-Morris-Pratt
func KMP(text, target string) []int       { return nil }
func RabinKarp(text, target string) []int { return nil }

const mod int = 1e9 + 7
const mod2 int = 1e9 + 9
const primeBase = 257 // 31

func CustomHash(s string, base int, mod int) int {
	// time: O(n), space: O(1)

	hash := 0
	for i := range s {
		hash = (hash*base + int(s[i]) + 1) % mod
	}
	return hash
}
func CustomHash2(s string, base int, mod int, ord func(c byte) int) int {
	// time: O(n), space: O(1)

	hash := 0
	for i := range s {
		hash = (hash*base + ord(s[i])) % mod
	}
	return hash
}
func Hash(s string) int {
	// time: O(n), space: O(1)

	return CustomHash(s, primeBase, mod)
}

func IntToDecStr(n int) string { return IntToStr(n, 10, 0) }
func IntToBinStr(n int) string { return IntToStr(n, 2, 0) }
func IntToOctStr(n int, upper bool) string {
	if upper {
		return IntToStr(n, 8, 'A')
	}
	return IntToStr(n, 8, 'a')
}
func IntToHexStr(n int, upper bool) string {
	if upper {
		return IntToStr(n, 16, 'A')
	}
	return IntToStr(n, 16, 'a')
}
func IntToStr(n int, base int, startLetter byte) string {
	// time: O(log n), space: O(log n)

	if n == 0 {
		return "0"
	}

	res := []byte{}
	for ; n != 0; n /= base {
		digit := n % base
		if digit < 10 {
			res = append(res, byte(digit)+'0')
		} else {
			res = append(res, byte(digit-10)+startLetter)
		}
	}
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return string(res)
}

func testIntToStr() {
	test := func(n int, base int, format string) {
		r1 := IntToStr(n, base, 'a')
		r2 := fmt.Sprintf(format, n)
		if r1 != r2 {
			fmt.Printf("IntToStr failed: (n = %v, base = %v) %q != %q\n", n, base, r1, r2)
		}
	}

	t := time.Now()
	for n := range int(1e5) {
		for _, baseFormat := range [][]any{{10, "%v"}, {2, "%b"}, {8, "%o"}, {16, "%x"}} {
			test(n, baseFormat[0].(int), baseFormat[1].(string))
		}

		// fmt.Printf("%d\n%v\n", n, IntToStr(n, 10))
		// fmt.Printf("%b\n%v\n", n, IntToStr(n, 2))
		// fmt.Printf("%o\n%v\n", n, IntToStr(n, 8))
		// fmt.Printf("%x\n%v\n", n, IntToStr(n, 16))
	}
	fmt.Printf("testIntToStr time: %v\n", time.Since(t))
}

func main() {
	testIntToStr()
}
