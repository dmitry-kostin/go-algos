package quicksort

import (
	"math/rand"
)

func Sort(nums []int) {
	quicksort(0, len(nums)-1, nums)
}

func quicksort(l, r int, nums []int) {
	if l < r {
		pivot := rand.Intn(r-l) + l
		pivot = partition(l, r, pivot, nums)
		quicksort(l, pivot-1, nums)
		quicksort(pivot+1, r, nums)
	}
}

func partition(l, r, p int, nums []int) int {
	pivot := nums[p]
	nums[r], nums[p] = nums[p], nums[r]
	ind := l
	for i := l; i <= r; i++ {
		if nums[i] < pivot {
			nums[ind], nums[i] = nums[i], nums[ind]
			ind++
		}
	}
	nums[r], nums[ind] = nums[ind], nums[r]
	return ind
}
