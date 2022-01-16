package quicksort

func partition(arr []int, leftIndex int, rightIndex int) int {
	// assign pivot rightmost value
	pivot := arr[rightIndex]

	// we have consumed the final value in the list as the pivot value. Move the
	// rightIndex down by one.
	// rightIndex--

	// position to place next element that is less than pivot.
	// the final step will be to place pivot in this slot, swapping
	// it with it's current position at the end of the list. For example,
	// if every element in the array is greater than pivot, this value will
	// remain at zero index, and pivot will be swapped with whatever is contained
	// in the beginning of the list
	//
	// in order to increment this value before using it, we subtract one. This allows
	// it to always be incremented.
	pivotPosition := leftIndex - 1

	for leftIndex < rightIndex {
		// verbose vars for verbosity, and debugging
		left := arr[leftIndex]

		// current left value is less than pivot. This means that the value should go before
		// pivot. We will will swap it with the current "pivot position", which is the index
		// pivot will be placed once the partition is finished. For example, if the first value
		// in the list is less than pivot, it will be swapped with itself, because it is already as
		// far left as it should be. following this pattern, if all values are less than pivot, the array
		// will remain unaltered, as each position, including pivot, is swapped with itself.
		if left < pivot {
			// increment pivot position.
			pivotPosition++

			// grab the current value, so we can assign it back after it is overwritten
			current := arr[pivotPosition]

			// assign left value to pivot position
			arr[pivotPosition] = left

			// assign the value that was overwritten where left was stored in list
			arr[leftIndex] = current
		}

		leftIndex++
	}

	// finally, place pivot in it's pivot position. follows same swap pattern as above.
	current := arr[pivotPosition+1]
	// we grab one greater than the pivot position, because the pivot should be placed
	// after the previous placements of smaller values.
	arr[pivotPosition+1] = pivot

	// finall, place value in pivot's previous slot.
	arr[rightIndex] = current

	// finally, return the pivot position and start the whole thing over.
	return pivotPosition + 1
}

func sort(arr []int, leftIndex int, rightIndex int) {
	// partition the list.
	pivotPosition := partition(arr, leftIndex, rightIndex)

	// great, we have a new pivot position. Next, we recurse into both
	// positions.

	// first, lets check if there's any stuff to the left that could be partitioned.
	if pivotPosition > leftIndex {
		sort(arr, leftIndex, pivotPosition-1)
	}

	// next, check for anything to the right of the current pivot
	if pivotPosition < rightIndex {
		sort(arr, pivotPosition+1, rightIndex)
	}
}

func Sort(arr []int) {
	rightIndex := len(arr) - 1
	leftIndex := 0

	sort(arr, leftIndex, rightIndex)
}
