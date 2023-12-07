package stalinSort

import (
	"reflect"
	"testing"
)

func TestSortInts(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{name: "empty", args: []int{}, want: []int{}},
		{name: "simple", args: []int{1, 2, 1}, want: []int{1, 2}},
		{name: "negs", args: []int{1, -2, 1}, want: []int{1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortDescInts(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{name: "empty", args: []int{}, want: []int{}},
		{name: "simple", args: []int{2, 4, 1}, want: []int{2, 1}},
		{name: "negs", args: []int{1, -2, 1}, want: []int{1, -2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortDesc(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortFloats(t *testing.T) {
	tests := []struct {
		name string
		args []float32
		want []float32
	}{
		{name: "empty", args: []float32{}, want: []float32{}},
		{name: "simple", args: []float32{1.0, 2.0, 1.1}, want: []float32{1.0, 2.0}},
		{name: "negs", args: []float32{1.1, -2.5, 1.1}, want: []float32{1.1, 1.1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
