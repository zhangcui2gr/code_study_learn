package array

func Sum(arr [5]int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func sumAll(sli1, slic2 []int) int {
	sum := 0

	for _, val := range sli1 {
		sum += val
	}
	for _, val := range slic2 {
		sum += val
	}

	return sum
}
