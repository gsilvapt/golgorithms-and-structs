package sort

func selectsort(array []int) []int {
	for i := 0; i < len(array); i++ {
		minVal := array[i]
		minPos := i
		for j := i+1; j < len(array); j++ {
			if array[j] < minVal {
				minVal = array[j]
				minPos = j
			}
		}
		array[minPos] = array[i]
		array[i] = minVal
	}

	return array
}
