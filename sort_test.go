package stalinSort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
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
