## Quicksort & Quickselect

Quicksort is an in-place sorting algorithm. Implementation is based on a random pivot point, which makes it harder to create a data set that generates O(N^2) performance. In a result average of O(NlogN) is achieved more consistently.

Quickselect is a selection algorithm to find the kth smallest element in an unordered list. Implementation is also based on a random pivot point selection. This way it gives a more stable O(N) time complexity selection results. 

### Usage

```go
// FindKthLargest in array
// problem can be interpreted as "find n - k smallest" and quickselect is a good way to go
func FindKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	n := len(nums)
	return Select(nums, n-k)
}
```

