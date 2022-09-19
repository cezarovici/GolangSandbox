package main

const IN = 1
const OUT = 0

const words = "Hello ! How are you ?\nFine,how about you ?\n- Ok"

func countWords(file string) int {
	var (
		state int
		wc    int
	)

	words := []byte(file)

	for _, char := range words {
		if split(char) {
			state = OUT
		} else if state == OUT {
			state = IN
			wc++
		}
	}

	return wc
}

func split(s byte) bool {
	return s <= 'A' || s >= 'z'
}

func main() {
	// go test -run TestCountWords
}
