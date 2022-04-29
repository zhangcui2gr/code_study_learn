package struct_method_interface

import "testing"

func Test_Area(t *testing.T) {
	checkRes := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		res := shape.Area()
		if res != expected {
			t.Errorf("res : %.2f, expected : %.2f", res, expected)
		}
	}

	t.Run("计算矩形面积", func(t *testing.T) {
		rectangle := Rectangle{1, 2}
		checkRes(t, rectangle, 3)
	})

	t.Run("计算圆形面积", func(t *testing.T) {
		circle := Circle{10}
		checkRes(t, circle, 314)
	})
}

// 两个测试函数分开写
//func TestRectangle_Area(t *testing.T) {
//	res := Rectangle{1, 2}.Area()
//	expected := 2.00
//	if res != expected {
//		t.Errorf("res : %.2f, expected : %.2f", res, expected)
//	}
//}
//
//func TestCircle_Area(t *testing.T) {
//	res := Circle{100}.Area()
//	expected := 314.00
//	if res != expected {
//		t.Errorf("res : %.2f, expected : %.2f", res, expected)
//	}
//
//}
