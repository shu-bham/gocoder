package util

// QuickSort is the main function that initiates the QuickSort algorithm.
func QuickSort(arr []int) {
	// Call the helper function to sort the entire array.
	quickSortHelper(arr, 0, len(arr)-1)
}

// quickSortHelper is a recursive function that sorts the array by partitioning it.
func quickSortHelper(arr []int, lo int, hi int) {
	// Continue sorting if there are at least two elements.
	if lo < hi {
		// Partition the array and get the partition index.
		partitionIndex := partition(arr, lo, hi)
		// Recursively sort the elements before and after the partition.
		quickSortHelper(arr, lo, partitionIndex-1)
		quickSortHelper(arr, partitionIndex+1, hi)
	}
}

// partition rearranges elements around a pivot and returns the index of the pivot.
func partition(arr []int, lo int, hi int) int {
	// Choose the first element as the pivot.
	pivot := arr[lo]
	// Initialize pointers for the left and right ends.
	i, j := lo, hi

	// Continue swapping elements until the pointers meet.
	for i < j {
		// Move the left pointer to the right until finding an element greater than the pivot.
		for arr[i] <= pivot && i <= hi-1 {
			i++
		}

		// Move the right pointer to the left until finding an element less than the pivot.
		for arr[j] >= pivot && j >= lo+1 {
			j--
		}

		// Swap elements if the left pointer is still to the left of the right pointer.
		if i < j {
			// Swap arr[i] and arr[j].
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Place the pivot in its correct position.
	arr[lo] = arr[j]
	arr[j] = pivot

	// Return the index of the pivot.
	return j
}
