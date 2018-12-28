package goarea

import "math"

//Pi é bla bla
const Pi = 3.1416

//Circ calcuça a area do circulo
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect calcula a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
