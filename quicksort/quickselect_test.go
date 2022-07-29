package quicksort

import (
	"math/rand"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test quickselect 1",
			args: args{
				nums: []int{5, 1, 3, 9, 6, 4, 2},
				k:    6,
			},
			want: 9,
		},
		{
			name: "Test quickselect 2",
			args: args{
				nums: []int{7, 10, 4, 3, 20, 15},
				k:    2,
			},
			want: 7,
		},
		{
			name: "Test quickselect 3",
			args: args{
				nums: []int{7, 10, 4, 3, 20, 15},
				k:    3,
			},
			want: 10,
		},
	}
	rand.Seed(time.Now().UnixNano())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			copy1 := make([]int, len(tt.args.nums))
			copy2 := make([]int, len(tt.args.nums))
			copy(copy1, tt.args.nums)
			copy(copy2, tt.args.nums)
			if got := Select(copy1, tt.args.k); got != tt.want {
				t.Errorf("Select(%v, %v) = %v, want %v", copy2, tt.args.k, got, tt.want)
			}
		})
	}
}
