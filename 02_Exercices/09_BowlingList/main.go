package main

func minMaxMed(list []int) []int {
	var max int
	var min int

	var highScore int
	var lowScore int

	var med int

	for i, element := range list {
		if i == 0 {
			max = element
			min = element
		}
		if element > max {
			max = element

			highScore++
		}
		if element < min {
			min = element

			lowScore++
		}
		med = med + element
	}

	return []int{highScore, lowScore, med / len(list)}
}

func main() {
	// go test -run TestMinMaxMed
}
