package main

import (
	"fmt"
	"time"
)

func CountSetBits(x int) int {
	// Hamming weight (pop count)
	// time: O(log n), space: O(1)

	count := 0
	for x != 0 {
		count++
		x &= x - 1
	}
	return count
}
func CountBits(x int) int {
	// time: O(log n), space: O(1)

	if x == 0 {
		return 1
	}

	count := 0
	for x != 0 {
		count++
		x >>= 1
	}
	return count
}

func SetBit(x int, pos int) int  { return x | (1 << pos) }
func OffBit(x int, pos int) int  { return x & ^(1 << pos) }
func FlipBit(x int, pos int) int { return x ^ (1 << pos) }

func OffLastSetBit(x int) int    { return x & (x - 1) }
func OffTrailingOnes(x int) int  { return x & (x + 1) }
func FlipLastUnsetBit(x int) int { return x | (x + 1) }
func SetKBits(k int) int         { return (1 << k) - 1 }

func RightmostSetBit(x int) int { return x & -x } // LowestSetBit

func PowerOfTwo(x int) int    { return 1 << x }
func IsPowerOfTwo(x int) bool { return x > 0 && x&(x-1) == 0 }
func DivByTwo(x int) int      { return x >> 1 }
func MulByTwo(x int) int      { return x << 1 }

// temp

func testCountBits() {
	t := time.Now()
	for i := range int(1e2) {
		bitWidth := CountBits(i)
		bitStr := fmt.Sprintf("%b", i)
		if bitWidth != len(bitStr) {
			fmt.Printf("CountBits failed: %v (%b) != %v (%s)\n", bitWidth, i, len(bitStr), bitStr)
		}
	}
	fmt.Printf("test CountBits time: %v\n", time.Since(t))
}

func main() {
	testCountBits()

	x := 0b1100011
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", SetBit(x, 3))
	fmt.Printf("%b\n", OffBit(x, 3))
	fmt.Printf("%b\n", 1<<3)
	fmt.Printf("%b\n", ^(1 << 3))

	x = FlipBit(x, 3)
	fmt.Printf("%b\n", x)
	x = FlipBit(x, 3)
	fmt.Printf("%b\n", x)
	x = FlipBit(x, 3)
	fmt.Printf("%b\n", x)

	fmt.Printf("%b\n", SetKBits(5))
	fmt.Println()

	y := int8(0b10100)
	_y := y
	fmt.Printf("%08b: %[1]v\n", y)
	// fmt.Printf("%08b\n", y)
	y = ^y
	// fmt.Printf("%08b: %[1]v\n", y)
	// fmt.Printf("%08b\n", y)
	fmt.Printf("%08b\n", uint8(y))
	fmt.Printf("%08b\n", uint8(y+1))
	fmt.Printf("%08b\n", _y&-_y)

	// fmt.Printf("%b\n", 5)  // output: 101
	// fmt.Printf("%b\n", -5) // output: -101
}
