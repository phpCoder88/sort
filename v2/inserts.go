package sort

// Inserts implements insertion sort algorithm
func Inserts(slice []int) []int {
	copiedSlice := make([]int, len(slice))
	copy(copiedSlice, slice)

	for i := 1; i < len(copiedSlice); i++ {
		for j := i; j > 0 && copiedSlice[j-1] > copiedSlice[j]; j-- {
			copiedSlice[j], copiedSlice[j-1] = copiedSlice[j-1], copiedSlice[j]
		}
	}

	return copiedSlice
}
