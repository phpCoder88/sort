// Package sort implements functions for slice sorting using different variation of algorithms
package sort

// Bubble implements bubble sort algorithm
func Bubble(slice []int) []int {
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)

	for i, isTouched := 0, true; isTouched; i++ {
		isTouched = false

		for j := 0; j < len(slice)-1-i; j++ {
			if sliceCopy[j] > sliceCopy[j+1] {
				isTouched = true
				sliceCopy[j], sliceCopy[j+1] = sliceCopy[j+1], sliceCopy[j]
			}
		}
	}

	return sliceCopy
}
