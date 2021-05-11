package main

import "math"

// Iniciando com letra maiúscula é PÚBLICO (visivel dentro e fora do pacote)
// Iniciando com letra minúscula é PRIVADO (visivel apenas dentro do pacote)

// Por exemplo...
// Point - gerará algo público
// point ou _Point - gerará algo privado

// Pointer represent a point on plane cartesian
type Pointer struct {
	x, y float64
}

func catetos(a, b Pointer) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// distance is responsible for calculating the linear distance between two points
func Distance(a, b Pointer) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
