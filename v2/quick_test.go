package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuick(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := Quick(tt.slice)

			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("test Failed - results not match\nGot:\n%v\nExpected:\n%v", res, tt.want)
			}
		})
	}
}

func BenchmarkQuickRandomSlice(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = Quick(randomSlice)
	}
	globalResult = res
}

func BenchmarkQuickSortedSlice(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = Quick(sortedSlice)
	}
	globalResult = res
}

func BenchmarkQuickReverseSortedSlice(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = Quick(reverseSortedSlice)
	}
	globalResult = res
}

func BenchmarkQuickAlmostSortedSlice(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = Quick(almostSortedSlice)
	}
	globalResult = res
}

func ExampleQuick() {
	var slice = []int{11, 4, 0, 9, 34, 76, 17, -5, 23, 96, 45, -2}
	fmt.Println(Quick(slice))
	// Output: [-5 -2 0 4 9 11 17 23 34 45 76 96]
}
