package figuras

import "math"

type Circle struct {
	Radio float64 
}

// Para circulos
func (c *Circle) area() float64{
	return math.Pi * (c.Radio * c.Radio)
}

func (c *Circle) perimetro() float64{
	return math.Pi * c.Radio * 2
}
