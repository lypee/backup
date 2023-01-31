package main

import (
	"reflect"
	"testing"
)

func Test_sortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "t1",
			args: args{nums: []int{}},
		},
		{
			name: "t2",
			args: args{nums: []int{1}},
		},
		{
			name: "t3",
			args: args{nums: []int{1, 0}},
		},
		{
			name: "t4",
			args: args{nums: []int{0, -1, 3, 1241, 1, 23, 124, 12, 34, 12, 312}},
		},
		{
			name: "t5",
			args: args{nums: []int{0, -1, 3, 1241, 1, 23, 124, 12, 34, 12, 312, 311}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
