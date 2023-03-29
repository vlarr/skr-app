package main

import "testing"

func Test_validateIngridByActiveEffects(t *testing.T) {
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
		name string
		args args
		want bool
	}{
		{"", args{&contextTest, []int{6, 7}}, false},
		{"", args{&contextTest, []int{7, 8}}, false},
		{"", args{&contextTest, []int{8, 9}}, true},
		{"", args{&contextTest, []int{9, 10}}, true},
		{"", args{&contextTest, []int{10, 11}}, true},
		{"", args{&contextTest, []int{8, 9, 11}}, true},
		{"", args{&contextTest, []int{8, 9, 12}}, true},
		{"", args{&contextTest, []int{8, 10, 11}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateIngridByActiveEffects(tt.args.contextPtr, tt.args.ingridIds); got != tt.want {
				t.Errorf("validateIngridByActiveEffects() = %v, want %v", got, tt.want)
			}
		})
	}
}
