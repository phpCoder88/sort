package sort

import (
	"fmt"
	"reflect"
	"testing"
)

var globalResult []int

var tests = []struct {
	name  string
	slice []int
	want  []int
}{
	{
		name:  "Пустой слайс",
		slice: []int{},
		want:  []int{},
	},
	{
		name:  "Набор 1",
		slice: []int{11, 4, 0, 9, 34, 76, 17, -5, 23, 96, 45, -2, -2},
		want:  []int{-5, -2, -2, 0, 4, 9, 11, 17, 23, 34, 45, 76, 96},
	},
	{
		name:  "Набор 2",
		slice: []int{53, 21, 88, 48, 75, 97, 36, 35, 51, 56, 82, 43, 37, 58, 96, 44, 57, 46, 50, 66, 49, 87, 2, 84, 33, 0, 67, 18, 44, 69},
		want:  []int{0, 2, 18, 21, 33, 35, 36, 37, 43, 44, 44, 46, 48, 49, 50, 51, 53, 56, 57, 58, 66, 67, 69, 75, 82, 84, 87, 88, 96, 97},
	},
	{
		name:  "Набор 3",
		slice: []int{31, -75, 8, 24, -19, -37, 8, -26, 21, -70, -25, 37, 39, -75, 1, -3, -7, -32, -47, -16, -7, -18, -41, -9, -36, 45, 84, -63, 16, 9},
		want:  []int{-75, -75, -70, -63, -47, -41, -37, -36, -32, -26, -25, -19, -18, -16, -9, -7, -7, -3, 1, 8, 8, 9, 16, 21, 24, 31, 37, 39, 45, 84},
	},
}

var randomSlice = []int{53, 21, 88, 48, 75, 97, 36, 35, 51, 56, 82, 43, 37, 58, 96, 44, 57, 46, 50, 66, 49, 87, 2, 84, 33, 0, 67, 18, 44, 69}
var sortedSlice = []int{0, 2, 18, 21, 33, 35, 36, 37, 43, 44, 44, 46, 48, 49, 50, 51, 53, 56, 57, 58, 66, 67, 69, 75, 82, 84, 87, 88, 96, 97}
var reverseSortedSlice = []int{97, 96, 88, 87, 84, 82, 75, 69, 67, 66, 58, 57, 56, 53, 51, 50, 49, 48, 46, 44, 44, 43, 37, 36, 35, 33, 21, 18, 2, 0}
var almostSortedSlice = []int{97, 2, 18, 21, 33, 35, 36, 37, 43, 44, 44, 46, 48, 49, 50, 51, 53, 56, 57, 58, 66, 67, 69, 75, 82, 84, 87, 88, 96, 0}

func TestSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := Sort(tt.slice)

			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("test Failed - results not match\nGot:\n%v\nExpected:\n%v", res, tt.want)
			}
		})
	}
}

func BenchmarkSortRandomSlice(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = Sort(randomSlice)
	}
	globalResult = res
}

func BenchmarkSortSortedSlice(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = Sort(sortedSlice)
	}
	globalResult = res
}

func BenchmarkSortReverseSortedSlice(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = Sort(reverseSortedSlice)
	}
	globalResult = res
}

func BenchmarkSortAlmostSortedSlice(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = Sort(almostSortedSlice)
	}
	globalResult = res
}

func ExampleSort() {
	var slice = []int{11, 4, 0, 9, 34, 76, 17, -5, 23, 96, 45, -2}
	fmt.Println(Sort(slice))
	// Output: [-5 -2 0 4 9 11 17 23 34 45 76 96]
}
