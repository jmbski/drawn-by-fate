package coords

// Coordinates interface represents any struct with a float X,Y value
type Coordinates interface {
	Value() (x, y *float64)
}

func AreEqual(c1, c2 Coordinates) bool {
	x1, y1 := c1.Value()
	x2, y2 := c2.Value()
	return x1 == x2 && y1 == y2
}

func Update(c1 Coordinates, x, y float64) {
	cX, cY := c1.Value()
	(*cX) += x
	(*cY) += y
}
