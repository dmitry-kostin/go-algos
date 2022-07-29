package quicksort

import (
	"math/rand"
)

// Select k-th smallest element using quickselect algorithm
// using random pivot point the average complexity is O(N)
func Select(nums []int, k int) int {
	n := len(nums)
	return quickselect(0, n-1, k, nums)
}

func quickselect(l, r, k int, nums []int) int {
	if l == r {
		return nums[l]
	}
	pivot := rand.Intn(r-l) + l
	pivot = partition(l, r, pivot, nums)
	if pivot == k {
		return nums[k]
	}
	if k < pivot {
		return quickselect(l, pivot-1, k, nums)
	}
	return quickselect(pivot+1, r, k, nums)
}

func quickSelectPartition(l, r, p int, nums []int) int {
	pivot := nums[p]
	nums[p], nums[r] = nums[r], nums[p]
	ind := l
	for i := l; i <= r; i++ {
		if nums[i] < pivot {
			nums[i], nums[ind] = nums[ind], nums[i]
			ind++
		}
	}
	nums[r], nums[ind] = nums[ind], nums[r]
	return ind
}
