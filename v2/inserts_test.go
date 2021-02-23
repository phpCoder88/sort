package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertsSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := Inserts(tt.slice)

			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("test Failed - results not match\nGot:\n%v\nExpected:\n%v", res, tt.want)
			}
		})
	}
}

func BenchmarkInsertsRandomSlice(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = Inserts(randomSlice)
	}
	globalResult = res
}

func BenchmarkInsertsSortedSlice(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = Inserts(sortedSlice)
	}
	globalResult = res
}

func BenchmarkInsertsReverseSortedSlice(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = Inserts(reverseSortedSlice)
	}
	globalResult = res
}

func BenchmarkInsertsAlmostSortedSlice(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = Inserts(almostSortedSlice)
	}
	globalResult = res
}

func ExampleInserts() {
	var slice = []int{11, 4, 0, 9, 34, 76, 17, -5, 23, 96, 45, -2}
	fmt.Println(Inserts(slice))
	// Output: [-5 -2 0 4 9 11 17 23 34 45 76 96]
}
