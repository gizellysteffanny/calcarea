package calcarea

import (
	"math"
)

/*
Pi é uma proporção numérica definida pela relação entre o perimetro de uma circunferência e seu diâmetro
*/
const Pi = 3.1416

// Circle é responsável por calcular a área da circunferência
func Circle(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsavel por calcular a área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível!
func _TriangleEq(base, altura float64) float64 {
	return base * altura / 2
}
