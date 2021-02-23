package sort

// Quick implements quick sort algorithm
func Quick(slice []int) []int {
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)

	sort(sliceCopy, 0, len(sliceCopy)-1)

	return sliceCopy
}

func sort(A []int, leftInd int, rightInd int) {
	if leftInd >= rightInd {
		return
	}

	middle := leftInd + (rightInd-leftInd)/2
	opora := A[middle]

	i, j := leftInd, rightInd

	for i <= j {
		for A[i] < opora {
			i++
		}
		for A[j] > opora {
			j--
		}

		if i <= j {
			A[i], A[j] = A[j], A[i]
			i++
			j--
		}
	}

	if leftInd < j {
		sort(A, leftInd, j)
	}

	if rightInd > i {
		sort(A, i, rightInd)
	}
}
