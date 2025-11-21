package aco

type City struct {
	ID int
	X  float64
	Y  float64
}

func NewCity(id int, x, y float64) City {
	return City{ID: id, X: x, Y: y}
}

