package main

import (
	"fmt"
)

type IPoint interface {
	X() float64
	Y() float64
	SetX(float64)
	SetY(float64)
}

type IPoint3D interface {
	// IPoint is embedded, so IPoint3D inherits all its methods
	IPoint
	Z() float64
	SetZ(float64)
}

type Point struct {
	x float64
	y float64
}

type Point3D struct {
	// Point is embedded, so Point3D inherit all its methods
	Point
	z float64
}

func NewPoint(x float64, y float64) *Point {
	p := new(Point)

	p.SetX(x)
	p.SetY(y)

	return p
}

func (p *Point) X() float64 {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}

func (p *Point) SetX(x float64) {
	p.x = x
}

func (p *Point) SetY(y float64) {
	p.y = y
}

func NewPoint3D(x float64, y float64, z float64) *Point3D {
	p := new(Point3D)
	p.SetX(x)
	p.SetY(y)
	// Point3D's method
	p.SetZ(z)

	return p
}

func (p *Point3D) Z() float64 {
	return p.z
}

func (p *Point3D) SetZ(z float64) {
	p.z = z
}

func main() {
	// Make a slice of IPoint
	points := make([]IPoint, 0)

	p1 := NewPoint(3, 4)
	p2 := NewPoint3D(1, 2, 5)
	fmt.Printf("%.2f\n", p2.Z())
	// every type that implements IPoint's method is valid
	points = append(points, p1, p2)

	for _, p := range points {
		fmt.Println(fmt.Sprintf("(%.2f %.2f)", p.X(), p.Y()))
	}

	points3D := make([]IPoint3D, 0)
	p3 := NewPoint3D(2, 3, 5)
	points3D = append(points3D, p2, p3)
	for _, p := range points3D {
		fmt.Println(fmt.Sprintf("(%.2f %.2f %.2f)", p.X(), p.Y(), p.Z()))
	}
}
