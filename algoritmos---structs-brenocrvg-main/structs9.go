package main

import "fmt"

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func adicionar(a *Aluno, nota float64) {
	a.notas = append(a.notas, nota)
}

func remover(a *Aluno, indice int) {
	if indice >= 0 && indice < len(a.notas) {
		a.notas = append(a.notas[:indice], a.notas[indice+1:]...)
	}
}

func media(a Aluno) float64 {
	soma := 0.0
	for _, nota := range a.notas {
		soma += nota
	}
	media := soma / float64(len(a.notas))
	return media
}

func main() {
	aluno := Aluno{nome: "João", idade: 20, notas: []float64{8.5, 9.0, 7.5, 6.9}}
	fmt.Println("notas do aluno:", aluno.notas)
	adicionar(&aluno, 8.0)
	fmt.Println("notas após adição:", aluno.notas)
	remover(&aluno, 1)
	fmt.Println("notas após remoção:", aluno.notas)
	media := media(aluno)
	fmt.Println("média do aluno:", media)
}
