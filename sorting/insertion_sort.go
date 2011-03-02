package gonuts

func InsertionSort(data Comparable) {
	for i := 1; i < data.Len(); i++ {
		insert(data, i)
	}
}

func insert(data Comparable, pos int) {
	for i := pos-1; i >= 0 && data.Less(i+1, i); i-- {
		data.Swap(i+1, i)
	}
}