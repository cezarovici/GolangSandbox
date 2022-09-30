package main

// "An ordered set of numbers is given. Re-sort the set so that the first
// maximum goes first, then the first minimum, then the second maximum, the
// second minimum, etc. Example:
// input [2,3,5,7,8,9] -[9,2,8,3,7,5]
// Ideally, make an =E2=80=9Cin place=E2=80=9D, i.e. without creating an add=
// itional array"

func sort(slice []int) []int {
	sorted := false

	for !sorted {

		sorted = true

		for i := range slice {
			if i != len(slice)-1 && slice[i] < slice[i+1] {

				slice[i], slice[i+1] = slice[i+1], slice[i]

				sorted = false
			}
		}
	}
	return slice
}

func sortMaxMin(slice []int) []int {
	lenght := len(slice) - 1
	sort(slice)

	for i := 0; i < lenght/2+1; i++ {
		if slice[i] != slice[lenght-i] {
			slice = append(slice, slice[i], slice[lenght-i])
		} else {
			slice = append(slice, slice[i])
		}
	}

	return slice[lenght+1:]
}

func main() {

}
