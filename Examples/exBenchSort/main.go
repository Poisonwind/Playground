package main

import "fmt"

func mergeSort(sl []int) []int {

	if len(sl) == 1 {
		return sl
	}

	mid := len(sl) / 2

	return merge(mergeSort(sl[:mid]), mergeSort(sl[mid:]))

}

func merge(left, right []int) []int {

	newLen, i, j := len(left)+len(right), 0, 0
	newSlice := make([]int, newLen, newLen)

	for k := 0; k < newLen; k++ {
		switch {
		case i > len(left)-1 && j < len(right):
			newSlice[k] = right[j]
			j++
		case j > len(right)-1 && i < len(left):
			newSlice[k] = left[i]
			i++
		case left[i] < right[j]:
			newSlice[k] = left[i]
			i++
		default:
			newSlice[k] = right[j]
			j++
		}
	}

	return newSlice

}

func bubleSort(sl []int) {

	for i := 1; i < len(sl); i++ {
		for j := 1; j <= len(sl)-i; j++ {
			if sl[j] < sl[j-1] {
				sl[j-1], sl[j] = sl[j], sl[j-1]
			}
		}
	}

}

func main() {

	ar := []int{1, 5, 8, 4, 9, 3, 0, 7, 2, 6}

	//ar = mergeSort(ar)
	bubleSort(ar)

	fmt.Println(ar)

}
