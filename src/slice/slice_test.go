package slice

import (
	"reflect"
	"testing"
)

//测试切片的创建方式 slice := array[1:] 到底长什么样
//func Testslicegen(t *testing.T) {
//	t.Run("测试切片生成")
//	arr := [5]int{1, 2, 3, 4, 5}
//	slic := slice_gen(arr)
//	expected_slice := []int{1, 2, 3, 4, 5}
//	if reflect.DeepEqual(slic, expected_slice) {
//		t.Errorf("slic : %v, expected_slice : %v", slic, expected_slice)
//	}
//}

func TestSumAgrs(t *testing.T) {
	t.Run("切片叠加", func(t *testing.T) {
		sli1 := []int{1, 2, 3}
		sli2 := []int{2, 3, 4}
		//res_sli := make([]int, 2)

		res_sli := SumAgrs(sli1, sli2)
		expected_sli := []int{6, 9}
		if !reflect.DeepEqual(res_sli, expected_sli) {
			t.Errorf("res_sli : %v, expected_sli: %v", res_sli, expected_sli)
		}
	})
	t.Run("测试切片生成", func(t *testing.T) {
		arr := [5]int{1, 2, 3, 4, 5}
		slic := slice_gen(arr)
		expected_slice := []int{2, 3, 4, 5}
		if !reflect.DeepEqual(slic, expected_slice) {
			t.Errorf("slic : %v, expected_slice : %v", slic, expected_slice)
		}
	})

}
