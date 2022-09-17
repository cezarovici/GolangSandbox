package reverse

import (
	"fmt"
)

func reverseByRune(s string) string {
	r := []rune(s)
	l := len(s) - 1
	for i := 0; i < (len(s)-1)/2; i++ {
		r[i], r[l-i] = r[l-i], r[i]
	}

	return string(r)
}
func reverseAppEnd(s string) string {
	var res string

	for i := len(s) - 1; i >= 0; i = i - 1 {
		res = res + s[i:i+1]
	}

	return res
}

func main() {

	str := "cezar"
	fmt.Println(reverseByRune(str))
	fmt.Println(reverseAppEnd(str))
}
