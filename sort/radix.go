package main

// todo: use generics

func RadixSort(nums []int)    { radixSort(nums, false) }
func PosRadixSort(nums []int) { radixSort(nums, true) }

func radixSort(nums []int, positive bool) {
	// time: O(d*(n+b)) -> O(d*n) ~> O(n), space: O(n+b) -> O(n)
	// n - len(nums)
	// d - max number of digits in nums
	// b - number base (10 in this case)

	var busketSize, minDigit int
	if positive {
		busketSize = 10
		minDigit = 0
	} else {
		busketSize = 9*2 + 1
		minDigit = -9
	}

	for e, done := 1, false; !done; e *= 10 {
		done = true
		buckets := make([][]int, busketSize)
		for _, num := range nums {
			digit := (num/e)%10 - minDigit
			buckets[digit] = append(buckets[digit], num)
			done = done && num/(e*10) == 0
		}
		i := 0
		for _, bucket := range buckets {
			for _, num := range bucket {
				nums[i] = num
				i++
			}
		}
	}
}

func radixSort1(nums []int) {
	// time: O(d*(n+b)) -> O(d*n) ~> O(n), space: O(n+b) -> O(n)
	// n - len(nums)
	// d - max number of digits in nums
	// b - number base (10 in this case)

	const bucketSize = 9*2 + 1
	const minDigit = -9

	for e, done := 1, false; !done; e *= 10 {
		done = true
		buckets := [bucketSize][]int{}
		for _, num := range nums {
			digit := (num/e)%10 - minDigit
			buckets[digit] = append(buckets[digit], num)
			done = done && num/(e*10) == 0
		}
		i := 0
		for _, bucket := range buckets {
			for _, num := range bucket {
				nums[i] = num
				i++
			}
		}
	}
}
