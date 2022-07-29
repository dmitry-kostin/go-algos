package quicksort

import (
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test sort 1", args: args{nums: []int{7, 10, 4, 3, 20, 15}}},
		{name: "test sort 2", args: args{nums: []int{5, 1, 3, 9, 6, 4, 2}}},
		{name: "test sort 3", args: args{nums: []int{3, 3, 1, 1, 2, 2, 6, 4, 2}}},
		{name: "test sort 4", args: args{nums: []int{1, 2, 3, 4, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.args.nums)
			prev := tt.args.nums[0]
			for i := 1; i < len(tt.args.nums); i++ {
				if tt.args.nums[i] < prev {
					t.Errorf("array is not in sorted order: %v", tt.args.nums)
				}
			}
		})
	}
}
