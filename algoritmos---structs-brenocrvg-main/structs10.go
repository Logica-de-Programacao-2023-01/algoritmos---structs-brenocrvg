package main

import "fmt"

type Filme struct {
	titulo     string
	diretor    string
	ano        int
	avaliacoes []int
}

func adiciona(f *Filme, avaliacao int) {
	f.avaliacoes = append(f.avaliacoes, avaliacao)
}

func remove(f *Filme, indice int) {
	if indice >= 0 && indice < len(f.avaliacoes) {
		f.avaliacoes = append(f.avaliacoes[:indice], f.avaliacoes[indice+1:]...)
	}
}

func calcularm(f Filme) float64 {
	soma := 0
	for _, avaliacao := range f.avaliacoes {
		soma += avaliacao
	}
	media := float64(soma) / float64(len(f.avaliacoes))
	return media
}

func main() {
	filme := Filme{titulo: "Homem Aranha", diretor: "Sam Raimi", ano: 2021, avaliacoes: []int{8, 7, 9}}
	fmt.Println("avaliações do filme:", filme.avaliacoes)
	adiciona(&filme, 6)
	fmt.Println("avaliações do filme após adição:", filme.avaliacoes)
	remove(&filme, 1)
	fmt.Println("avaliações do filme após remoção:", filme.avaliacoes)
	media := calcularm(filme)
	fmt.Println("média das avaliações do filme:", media)
}
