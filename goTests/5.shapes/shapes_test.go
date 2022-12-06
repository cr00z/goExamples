package shapes

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	want := 40.

	got := Perimeter(rectangle)

	assert.Equal(t, got, want)
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assert.Equal(t, got, want)
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		want := 72.
		checkArea(t, rectangle, want)
	})
	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		want := math.Pi * 100
		checkArea(t, circle, want)
	})
}

func TestArea2(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72},
		{name: "Circle", shape: Circle{10}, hasArea: math.Pi * 100},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			gotArea := tt.shape.Area()
			assert.Equal(t, gotArea, tt.hasArea)
		})
	}
}
