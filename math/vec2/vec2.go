package vec2

import (
	"math"

	mth "github.com/puoklam/collection/math"
)

type Vec2 [2]float64

func New(x, y float64) Vec2 {
	return Vec2{x, y}
}

func (v Vec2) Add(other Vec2) Vec2 {
	v[0] += other[0]
	v[1] += other[1]
	return v
}

func (v Vec2) Sub(other Vec2) Vec2 {
	v[0] -= other[0]
	v[1] -= other[1]
	return v
}

func (v Vec2) Mul(f float64) Vec2 {
	v[0] *= f
	v[1] *= f
	return v
}

func (v Vec2) MagnitudeSquared() float64 {
	return v[0]*v[0] + v[1]*v[1]
}

func (v Vec2) Magnitude() float64 {
	return math.Sqrt(v.MagnitudeSquared())
}

func (v Vec2) Normalize() Vec2 {
	m := v.Magnitude()
	v[0] /= m
	v[1] /= m
	return v
}

func (v Vec2) Zero() Vec2 {
	v[0], v[1] = 0, 0
	return v
}

func (v Vec2) Reverse() Vec2 {
	v[0] *= -1
	v[1] *= -1
	return v
}

func (v Vec2) Rotate(other Vec2, t float64) Vec2 {
	if t == 0 {
		return v
	}
	x, y := v[0]-other[0], v[1]-other[1]
	cos, sin := math.Cos(t), math.Sin(t)
	x, y = x*cos-y*sin, x*sin+y*cos
	v[0], v[1] = other[0]+x, other[1]+y
	return v
}

func Copy(v Vec2) Vec2 {
	return Vec2{v[0], v[1]}
}

func Add(v1, v2 Vec2) Vec2 {
	return Vec2{v1[0] + v2[0], v1[1] + v2[1]}
}

func Sub(v1, v2 Vec2) Vec2 {
	return Vec2{v1[0] - v2[0], v1[1] - v2[1]}
}

func Mul(v Vec2, f float64) Vec2 {
	return Vec2{v[0] * f, v[1] * f}
}

func Dot(v1, v2 Vec2) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1]
}

func Cross(v1, v2 Vec2) Vec2 {
	return Vec2{0, 0}
}

func Normalized(v Vec2) Vec2 {
	cp := Copy(v)
	return cp.Normalize()
}

func IsUnitVector(v Vec2) bool {
	m := v.Magnitude()
	return mth.FloatEqual(m, 1)
}

func IsZeroVector(v Vec2) bool {
	ms := v.MagnitudeSquared()
	return mth.FloatEqual(ms, 0)
}

func IsOrthogonal(v1, v2 Vec2) bool {
	return mth.FloatEqual(Dot(v1, v2), 0)
}

func Cos(v1, v2 Vec2) float64 {
	dp := Dot(v1, v2)
	m1, m2 := v1.Magnitude(), v2.Magnitude()
	if mth.FloatEqual(m1, 0) || mth.FloatEqual(m2, 0) {
		return math.NaN()
	}
	return dp / (m1 * m2)
}

func Sin(v1, v2 Vec2) float64 {
	cp := Cross(v1, v2)
	m1, m2 := v1.Magnitude(), v2.Magnitude()
	if mth.FloatEqual(m1, 0) || mth.FloatEqual(m2, 0) {
		return math.NaN()
	}
	return cp.Magnitude() / (m1 * m2)
}

func Angle(v1, v2 Vec2) float64 {
	return math.Acos(Cos(v1, v2))
}

func Rotate(v1, v2 Vec2, t float64) Vec2 {
	cp := Copy(v1)
	return cp.Rotate(v2, t)
}

func Projection(v1, v2, o Vec2) Vec2 {
	a := Sub(v1, o)
	b := Sub(v2, o)
	t := Dot(a, b) / Dot(b, b)
	return Add(o, Mul(v2, t))
}

// https://blog.gopheracademy.com/advent-2017/a-tale-of-two-rands/
