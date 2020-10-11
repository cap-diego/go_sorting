package sorting


// MergeSort given array of numbers returns the ordered array
func MergeSort(arr *[]int) []int{
	len := len(*arr)

	if (len == 2) {
		new := make([]int, 2)
		aux := (*arr)[0]
		new[0], new[1] = aux, (*arr)[1]
		if (aux > new[1]) {
			new[0], new[1] = new[1], aux
		}
		return new
	}else if len == 1 {
		return *arr
	}
	middleIdx := len/2
	arr1 := (*arr)[0:middleIdx]
	arr2 :=  (*arr)[middleIdx:len]
	ordArr1 := MergeSort(&arr1)
	ordArr2 := MergeSort(&arr2)
	return merge(&ordArr1, &ordArr2)

}

func merge(arr1, arr2 *[]int) []int{

	//mergedArray := [len(*arr1)+len(*arr2)]int{}
	var mergedArray []int = *arr1
	for _, v := range(*arr2){
		mergedArray = append(mergedArray, v)
	}
	return mergedArray
}