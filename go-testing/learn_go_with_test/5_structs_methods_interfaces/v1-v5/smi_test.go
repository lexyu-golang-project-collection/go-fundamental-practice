package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := &Rectangle{10.0, 10.0}
	result := Perimeter(rectangle)
	expected := 40.0

	if result != expected {
		t.Errorf("result : %.2f, expected : %.2f", result, expected)
	}
}

func TestArea(t *testing.T) {

	// v3, v4
	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		result := shape.Area()
		if result != expected {
			t.Errorf("result : %.2f, expected : %.2f", result, expected)
		}
	}

	t.Run("calculate rectangle area", func(t *testing.T) {
		rectangle := &Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("calculate circle area", func(t *testing.T) {
		circle := &Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})

	// v2
	/*
		t.Run("calculate rectangle area", func(t *testing.T) {
			rectangle := &Rectangle{12.0, 6.0}
			result := rectangle.Area()
			expected := 72.0

			if result != expected {
				t.Errorf("result : %.2f, expected : %.2f", result, expected)
			}
		})

		t.Run("calculate circle area", func(t *testing.T) {
			circle := &Circle{10.0}
			result := circle.Area()
			expected := 314.1592653589793

			if result != expected {
				t.Errorf("result : %.2f, expected : %.2f", result, expected)
			}
		})
	*/
}
