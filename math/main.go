package main

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

// greatest common divisor
func GCD(a, b int) int { return 0 }

// least common multiple
func LCM(a, b int) int     { return 0 }
func DivCeil(a, b int) int { return 0 }
func CountBits(x int) int  { return 0 }
