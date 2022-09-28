package main

// "You are given an array which contains only 0,1,2:
// (e.g. [2,0,1,1,0,0,2,2]). How would you sort it? What is the complexity of the solution you provided?"

func sortBySwapping(array []int) []int {
	high := len(array) - 1
	low := 0
	mid := 0

	for mid <= high {
		switch array[mid] {
		case 0:
			array[low], array[mid] = array[mid], array[low]as
			low++
		case 2:
			array[high], array[mid] = array[mid], array[high]
			high--
			mid--
		}
		mid++
	}

	return array
}

func main() {
}
