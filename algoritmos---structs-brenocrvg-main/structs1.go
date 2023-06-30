package main

import (
	"fmt"
	"math"
)

type Circuito struct {
	raio float64
}

func calcular(r Circuito) float64 {
	area := r.raio * r.raio * math.Pi
	return area
}

func main() {
	r := Circuito{raio: 5.7}
	s := calcular(r)
	fmt.Println(s)
}
