package sort

func bubblesort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; i < len(array)-i-1; i++ {
			if array[j] > array[j+1] {
				prevVal := array[j]
				nextVal := array[j+1]
				array[j] = nextVal
				array[j+1] = prevVal
			}
		}
	}

	return array
}
