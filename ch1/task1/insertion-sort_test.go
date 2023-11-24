package task1_test

import (
	"reflect"
	"testing"

	"algorithms/ch1/task1"
)

func TestInsertionSortInt(t *testing.T) {
	tcs := []struct {
		arg  []int
		want []int
	}{
		{
			arg:  []int{2, 3, 1, 5, 4},
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range tcs {
		task1.InsertionSortInt(tc.arg)
		ok := reflect.DeepEqual(tc.arg, tc.want)
		if !ok {
			t.Fatal("want: %w, got: %w", tc.want, tc.arg)
		}
	}
}
