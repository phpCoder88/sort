package sort

import (
	"fmt"
	"reflect"
	"testing"
)

var globalResult []int

func TestBubbleSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := Bubble(tt.slice)

			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("test Failed - results not match\nGot:\n%v\nExpected:\n%v", res, tt.want)
			}
		})
	}
}

func BenchmarkBubbleRandomSlice(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = Bubble(randomSlice)
	}
	globalResult = res
}

func BenchmarkBubbleSortedSlice(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = Bubble(sortedSlice)
	}
	globalResult = res
}

func BenchmarkBubbleReverseSortedSlice(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = Bubble(reverseSortedSlice)
	}
	globalResult = res
}

func BenchmarkBubbleAlmostSortedSlice(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = Bubble(almostSortedSlice)
	}
	globalResult = res
}

func ExampleBubble() {
	var slice = []int{11, 4, 0, 9, 34, 76, 17, -5, 23, 96, 45, -2}
	fmt.Println(Bubble(slice))
	// Output: [-5 -2 0 4 9 11 17 23 34 45 76 96]
}
