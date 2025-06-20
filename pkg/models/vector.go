package models

import "math"

type Vec2 struct {
	X, Y float64
}

func (v Vec2) Len() float64 {
	return math.Hypot(v.X, v.Y)
}

func (v Vec2) Normalize() Vec2 {
	len := v.Len()
	if len == 0 {
		return Vec2{}
	}
	return Vec2{X: v.X / len, Y: v.Y / len}
}

func (v Vec2) Diff(other Vec2) Vec2 {
	return Vec2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vec2) Scale(s float64) Vec2 {
	return Vec2{X: v.X * s, Y: v.Y * s}
}

func NewVec2(x, y float64) Vec2 {
	return Vec2{X: x, Y: y}
}
