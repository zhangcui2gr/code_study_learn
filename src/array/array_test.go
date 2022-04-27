package array

import "testing"

func TestSum(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	res := Sum(arr)
	expected := 15

	if res != expected {
		t.Errorf("res = %v, expected = %v, array : %v", res, expected, arr)
	}
}
