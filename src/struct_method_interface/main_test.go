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
