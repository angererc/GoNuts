package gonuts

//we define the same interface as the official sort package
type Comparable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}