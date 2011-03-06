package gonuts

func SelectionSort(data Comparable) {
	for i := data.Len()-1; i >= 1; i-- {
		max := selectMax(data, i)
		data.Swap(i, max)
	}
}

func selectMax(data Comparable, right int) int {
	max := 0
	
	for i := 0; i <= right; i++ {
		if data.Less(max, i) {
			max = i
		}
	}
	
	return max
}