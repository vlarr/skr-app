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

func Test_calculateWorth(t *testing.T) {
	contextTest := context{
		effectIdToInfoMap: map[int]*effectInfo{
			3: {id: 3, name: "effect1", worth: 100},
			4: {id: 4, name: "effect2", worth: 200},
			5: {id: 5, name: "effect2", worth: 400},
		},
		ingridIdToInfoMap: map[int]*ingridInfo{
			6:  {id: 6, name: "ingrid1", effectIdArr: [4]int{0, 0, 0, 0}},
			7:  {id: 7, name: "ingrid2", effectIdArr: [4]int{1, 1, 1, 1}},
			8:  {id: 8, name: "ingrid3", effectIdArr: [4]int{3, 4, 0, 1}},
			9:  {id: 9, name: "ingrid4", effectIdArr: [4]int{3, 5, 0, 1}},
			10: {id: 10, name: "ingrid5", effectIdArr: [4]int{3, 4, 5, 1}},
			11: {id: 11, name: "ingrid5", effectIdArr: [4]int{0, 3, 4, 5}},
		},
	}

	type args struct {
		contextInst *context
		ingridId1   int
		ingridId2   int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"", args{&contextTest, 6, 7}, 0},
		{"", args{&contextTest, 6, 8}, 0},
		{"", args{&contextTest, 6, 9}, 0},
		{"", args{&contextTest, 7, 8}, 0},
		{"", args{&contextTest, 7, 9}, 0},
		{"", args{&contextTest, 8, 9}, 100},
		{"", args{&contextTest, 8, 10}, 300},
		{"", args{&contextTest, 8, 11}, 300},
		{"", args{&contextTest, 9, 10}, 500},
		{"", args{&contextTest, 9, 11}, 500},
		{"", args{&contextTest, 10, 11}, 700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateWorth(tt.args.contextInst, tt.args.ingridId1, tt.args.ingridId2); got != tt.want {
				t.Errorf("calculateWorth() = %v, want %v", got, tt.want)
			}
		})
	}
}
