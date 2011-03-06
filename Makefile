include $(GOROOT)/src/Make.inc

TARG=gonuts
GOFILES=\
	sorting/sort.go\
	sorting/insertion_sort.go\
	sorting/quick_sort.go\
	sorting/selection_sort.go\

include $(GOROOT)/src/Make.pkg