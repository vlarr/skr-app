package main

import (
	"reflect"
	"testing"
)

func prepareIdSet(ids ...int) map[int]bool {
	result := map[int]bool{}
	for _, id := range ids {
		result[id] = true
	}
	return result
}

func Test_findEffectIdPair(t *testing.T) {
	type args struct {
		a [][4]int
	}
	tests := []struct {
		name string
		args args
		want map[int]bool
	}{
		{"", args{[][4]int{{0, 0, 0, 0}}}, prepareIdSet()},
		{"", args{[][4]int{{0, 0, 0, 0}, {0, 0, 0, 0}}}, prepareIdSet()},
		{"", args{[][4]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}}, prepareIdSet()},

		{"", args{[][4]int{{2, 0, 0, 0}, {0, 0, 0, 0}}}, prepareIdSet()},
		{"", args{[][4]int{{2, 0, 0, 0}, {2, 0, 0, 0}}}, prepareIdSet(2)},
		{"", args{[][4]int{{2, 3, 0, 0}, {2, 4, 0, 0}}}, prepareIdSet(2)},
		{"", args{[][4]int{{2, 3, 4, 0}, {2, 3, 5, 0}}}, prepareIdSet(2, 3)},
		{"", args{[][4]int{{2, 3, 4, 5}, {2, 3, 4, 5}}}, prepareIdSet(2, 3, 4, 5)},

		{"", args{[][4]int{{2, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}}, prepareIdSet()},
		{"", args{[][4]int{{2, 0, 0, 0}, {2, 0, 0, 0}, {0, 0, 0, 0}}}, prepareIdSet(2)},
		{"", args{[][4]int{{2, 0, 0, 0}, {2, 0, 0, 0}, {2, 0, 0, 0}}}, prepareIdSet(2)},
		{"", args{[][4]int{{2, 0, 0, 0}, {3, 0, 0, 0}, {2, 3, 0, 0}}}, prepareIdSet(2, 3)},
		{"", args{[][4]int{{2, 3, 0, 0}, {2, 4, 0, 0}, {4, 3, 0, 0}}}, prepareIdSet(2, 3, 4)},

		{"", args{[][4]int{{2, 3, 8, 9}, {2, 4, 5, 8}, {3, 4, 6, 7}, {5, 6, 7, 9}}}, prepareIdSet(2, 3, 4, 5, 6, 7, 8, 9)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findActiveEffectsByIngridEffects(tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findActiveEffectsByIngridEffects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateWorth(t *testing.T) {
	contextTest := context{
		effectIdToInfoMap: map[int]*effectInfo{
			3: {id: 3, name: "effect1", worth: 100},
			4: {id: 4, name: "effect2", worth: 200},
			5: {id: 5, name: "effect3", worth: 400},
			6: {id: 6, name: "effect4", worth: 800},
		},
		ingridIdToInfoMap: map[int]*ingridInfo{
			6:  {id: 6, name: "ingrid1", effectIds: [4]int{0, 0, 0, 0}},
			7:  {id: 7, name: "ingrid2", effectIds: [4]int{1, 1, 1, 1}},
			8:  {id: 8, name: "ingrid3", effectIds: [4]int{3, 4, 0, 1}},
			9:  {id: 9, name: "ingrid4", effectIds: [4]int{3, 5, 0, 1}},
			10: {id: 10, name: "ingrid5", effectIds: [4]int{3, 4, 5, 1}},
			11: {id: 11, name: "ingrid6", effectIds: [4]int{0, 3, 4, 5}},
			12: {id: 12, name: "ingrid7", effectIds: [4]int{5, 4, 0, 0}},
		},
	}

	type args struct {
		contextPtr *context
		ingridIds  []int
	}
	tests := []struct {
		name       string
		args       args
		wantExists bool
		wantWorth  float64
	}{
		{"", args{&contextTest, []int{6, 7}}, false, 0},
		{"", args{&contextTest, []int{6, 8}}, false, 0},
		{"", args{&contextTest, []int{6, 9}}, false, 0},
		{"", args{&contextTest, []int{7, 8}}, false, 0},
		{"", args{&contextTest, []int{7, 9}}, false, 0},
		{"", args{&contextTest, []int{8, 9}}, true, 100},
		{"", args{&contextTest, []int{8, 10}}, true, 300},
		{"", args{&contextTest, []int{8, 11}}, true, 300},
		{"", args{&contextTest, []int{9, 10}}, true, 500},
		{"", args{&contextTest, []int{9, 11}}, true, 500},
		{"", args{&contextTest, []int{10, 11}}, true, 700},
		{"", args{&contextTest, []int{8, 9, 11}}, true, 700},
		{"", args{&contextTest, []int{8, 9, 12}}, true, 700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExists, gotWorth := calculateWorth(tt.args.contextPtr, tt.args.ingridIds...)
			if gotExists != tt.wantExists {
				t.Errorf("calculateWorth() gotExists = %v, want %v", gotExists, tt.wantExists)
			}
			if gotWorth != tt.wantWorth {
				t.Errorf("calculateWorth() gotWorth = %v, want %v", gotWorth, tt.wantWorth)
			}
		})
	}
}
