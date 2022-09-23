package main

func parseArray(array []int) []int {
	var (
		left     int
		right    int
		maxLeft  int
		maxRight int
		onlyOne  bool
	)

	onlyOne = true

	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			if left == 0 {
				left = i

				onlyOne = false
			}
		} else {
			right = i + 1

			if len(array[left:right]) > len(array[maxLeft:maxRight]) {
				maxLeft = left
				maxRight = right
			}

			left = 0
		}

	}
	if right == 0 {
		return array
	}

	if onlyOne {
		return array[0:1]
	}

	return array[maxLeft:maxRight]
}
