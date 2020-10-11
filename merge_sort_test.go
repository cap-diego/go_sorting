package sorting_test


import (
	"testing"
	"sort"
	sorting "github.com/cap-diego/sorting"
)

func Test01ArrayWith2ElemsReturnsOrdered(t *testing.T) {
	arr := []int{10, 1}
	ordered := sorting.MergeSort(&arr)
	if ordered[0] != arr[1] || ordered[1] != arr[0] {
		t.Error("Error test 01")
	}
}

func Test02SingleElement(t *testing.T) {
	arr := []int{100}
	ordered := sorting.MergeSort(&arr)
	if !(ordered[0] == arr[0] && len(arr) == len(ordered)) {
		t.Error("Error test 02")
	}
}

func Test03(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ordered := sorting.MergeSort(&arr)
	if !sort.IntsAreSorted(ordered) {
		t.Error("Error test 03")
	}
}

func Test04(t *testing.T) { 
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	ordered := sorting.MergeSort(&arr)
	if sort.IntsAreSorted(ordered) {
		t.Error("Error test 04")
	}
}