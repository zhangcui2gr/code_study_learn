package slice

func Sum(sli []int) int {
	sum := 0
	for _, val := range sli {
		sum += val
	}
	return sum
}

func SumAgrs(numSlice ...[]int) []int {
	var sumSlice []int
	for _, member := range numSlice {
		sumSlice = append(sumSlice, Sum(member))
	}
	return sumSlice

}

func slice_gen(arr [5]int) []int {
	slic := arr[0:]
	return slic
}
