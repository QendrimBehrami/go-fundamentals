package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name     string
		shape    Shape
		hasPerim float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasPerim: 36.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasPerim: 62.83185307179586},
		{name: "Triangle", shape: Triangle{SideA: 3, SideB: 4, SideC: 5}, hasPerim: 12.0},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasPerim {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPerim)
			}
		})
	}
}
func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{SideA: 3, SideB: 4, SideC: 5}, hasArea: 6.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}

func BenchmarkRectangleArea(b *testing.B) {
	rectangle := Rectangle{Width: 12, Height: 6}
	for b.Loop() {
		rectangle.Area()
	}
}

func BenchmarkCircleArea(b *testing.B) {
	circle := Circle{Radius: 10}
	for b.Loop() {
		circle.Area()
	}
}

func BenchmarkTriangleArea(b *testing.B) {
	triangle := Triangle{SideA: 3, SideB: 4, SideC: 5}
	for b.Loop() {
		triangle.Area()
	}
}
