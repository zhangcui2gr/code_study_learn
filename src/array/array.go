package array

func Sum(arr [5]int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
