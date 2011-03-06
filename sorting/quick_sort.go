package gonuts

func QuickSort(data Comparable) {
	quicksort(data, 0, data.Len()-1)
}

func quicksort(data Comparable, left, right int) {
	if left < right {
		pi := partition(data, left, right)
		quicksort(data, left, pi-1)
		quicksort(data, pi+1, right)
	}
}

func partition(data Comparable, left, right int) int {
	p := selectPivot(data, left, right)
	data.Swap(p, right)
	store := left
	for i := left; i < right; i++ {
		if data.Less(i, right) {
			data.Swap(i, store)
			store++
		}
	}
	data.Swap(store, right)
	return store
}

func selectPivot(data Comparable, left, right int) int {
	return right//(left+right)/2
}