package main

import (
	"reflect"
	"testing"
)

func TestIter_incIndex(t *testing.T) {
	values := &[]int{1, 3, 5}

	type fields struct {
		values  *[]int
		indexes []int
	}
	type args struct {
		indexNum int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"", fields{values: values, indexes: []int{0, 0, 0}}, args{indexNum: 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Iter{
				allValues: tt.fields.values,
				indexes:   tt.fields.indexes,
			}
			if got := p.incIndex(tt.args.indexNum); got != tt.want {
				t.Errorf("incIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter_next(t *testing.T) {
	values := &[]int{1, 3, 5, 7, 9}

	type fields struct {
		values  *[]int
		indexes []int
	}
	tests := []struct {
		name         string
		fields       fields
		wantExists   bool
		wantIterInst Iter
	}{
		{"", fields{values: values, indexes: []int{0, 0, 0}}, true, Iter{allValues: values, indexes: []int{1, 0, 0}}},
		{"", fields{values: values, indexes: []int{4, 0, 0}}, true, Iter{allValues: values, indexes: []int{0, 1, 0}}},
		{"", fields{values: values, indexes: []int{4, 4, 0}}, true, Iter{allValues: values, indexes: []int{0, 0, 1}}},
		{"", fields{values: values, indexes: []int{4, 4, 4}}, false, Iter{allValues: values, indexes: []int{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Iter{
				allValues: tt.fields.values,
				indexes:   tt.fields.indexes,
			}
			gotExists, gotIterPtr := p.next()
			if gotExists != tt.wantExists {
				t.Errorf("next() gotExists = %v, want %v", gotExists, tt.wantExists)
			}
			if !reflect.DeepEqual(gotIterPtr, tt.wantIterInst) {
				t.Errorf("next() gotIterPtr = %v, want %v", gotIterPtr, tt.wantIterInst)
			}
		})
	}
}

func TestIter_getValues(t *testing.T) {
	values := &[]int{1, 3, 5, 7, 9}

	type fields struct {
		allValues *[]int
		indexes   []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{"", fields{allValues: values, indexes: []int{0, 0}}, []int{1, 1}},
		{"", fields{allValues: values, indexes: []int{2, 3}}, []int{5, 7}},
		{"", fields{allValues: values, indexes: []int{4, 4}}, []int{9, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Iter{
				allValues: tt.fields.allValues,
				indexes:   tt.fields.indexes,
			}
			if got := p.getValues(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
