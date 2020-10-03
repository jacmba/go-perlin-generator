package noise

import "math"

type vector2 struct {
	x, y float64
}

// Linear interpolation between a0 and a1 with w weight value for smoothing
func interpolate(a0, a1, w float64) float64 {
	return (1.0-w)*a0 + w*a1
}

// Create random direction vector
func randomGradient(x, y int) vector2 {
	xf := float64(x)
	yf := float64(y)
	random := 290.0 * math.Sin(xf*21942.0+yf*171324.0+8912.0) * math.Cos(xf*23157.0*yf*217832.0+9758.0)
	return vector2{math.Cos(random), math.Sin(random)}
}

// Compute dot product of distance and gradient vectors
func dotGridGradient(xi, yi int, x, y float64) float64 {
	gradient := randomGradient(xi, yi)

	dx := x - float64(xi)
	dy := y - float64(yi)

	return (dx*gradient.x + dy*gradient.y)
}

// Perlin compute perlin at coordinates x, y
func Perlin(x, y float64) float64 {
	x0 := int(x)
	x1 := x0 + 1
	y0 := int(y)
	y1 := y0 + 1

	sx := x - float64(x0)
	sy := y - float64(y0)

	n0 := dotGridGradient(x0, y0, x, y)
	n1 := dotGridGradient(x1, y0, x, y)
	i0 := interpolate(n0, n1, sx)

	n0 = dotGridGradient(x0, y1, x, y)
	n1 = dotGridGradient(x1, x1, x, y)
	i1 := interpolate(n0, n1, sx)

	return interpolate(i0, i1, sy)
}
