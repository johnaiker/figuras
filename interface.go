package figuras

import "fmt"

type Geometria interface {
	area() float64
	perimetro() float64
}


func Medidas(g Geometria) {
	fmt.Printf("\nFigura: %v  / Area: %v  / Perimetro: %v \n", g, g.area(),g.perimetro())
}

