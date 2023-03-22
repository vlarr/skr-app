package main

import (
	"reflect"
	"testing"
)

func Test_findEffectIdPair(t *testing.T) {
	type args struct {
		a [][4]int
	}
	tests := []struct {
		name string
		args args
		want *[]int
	}{
		// TODO: Add test cases.
		{"test", args{[][4]int{{0, 0, 0, 0}, {0, 0, 0, 0}}}, &[]int{}},
		{"test", args{[][4]int{{2, 0, 0, 0}, {0, 0, 0, 0}}}, &[]int{}},
		{"test", args{[][4]int{{2, 0, 0, 0}, {2, 0, 0, 0}}}, &[]int{2}},
		{"test", args{[][4]int{{2, 3, 0, 0}, {2, 4, 0, 0}}}, &[]int{2}},
		{"test", args{[][4]int{{2, 3, 4, 0}, {2, 3, 5, 0}}}, &[]int{2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findEffectIdPair(tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findEffectIdPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
