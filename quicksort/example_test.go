package quicksort

import (
	"fmt"
	"math/rand"
	"time"
)

// findKthLargest in array
// problem can be interpreted as "find n - k smallest" and quickselect is a good way to go
func FindKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	n := len(nums)
	return Select(nums, n-k)
}

func SelectPrintTest() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len([]int{7, 5, 8, 1, 3, 9, 6, 4, 2}); i++ {
		fmt.Println(Select([]int{7, 5, 8, 1, 3, 9, 6, 4, 2}, i))
	}
}
