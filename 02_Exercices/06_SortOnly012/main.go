package main

import "log"

// "You are given an array which contains only 0,1,2:
// (e.g. [2,0,1,1,0,0,2,2]). How would you sort it? What is the complexity of the solution you provided?"

func getArray() []int {
	return []int{2, 2, 2, 0, 0, 0, 1, 1, 1}
}

func swap(x, y *int) {
	aux := *x
	*x = *y
	*y = aux
}

func sort() []int {
	array := getArray()

	high := len(array) - 1
	low := 0
	mid := 0

	for mid <= high {
		switch array[mid] {
		case 0:
			swap(&array[low], &array[mid])
			low++
		case 1:
			swap(&array[mid], &array[high-low+1])
		case 2:
			swap(&array[mid], &array[high])
			high--
		}
		mid++
	}

	return array
}

func main() {
	log.Print(sort())
}
