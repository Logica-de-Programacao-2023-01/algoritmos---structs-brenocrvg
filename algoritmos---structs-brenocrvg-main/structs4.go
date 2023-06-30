package main

import (
	"fmt"
	"time"
)

type Musica struct {
	Titulo  string
	Artista string
	Duracao time.Duration
}

type Playlist struct {
	Nome    string
	Musicas []Musica
}

func ImprimirPlaylist(p Playlist) {
	fmt.Println("Nome da Playlist:", p.Nome)

	tempoTotal := time.Duration(0)
	for _, musica := range p.Musicas {
		tempoTotal += musica.Duracao
		fmt.Println("Música:", musica.Titulo)
		fmt.Println("Artista:", musica.Artista)
		fmt.Println("Duração:", musica.Duracao)
		fmt.Println("----------------------")
	}

	fmt.Println("Tempo total da Playlist:", tempoTotal)
}

func main() {
	musica1 := Musica{
		Titulo:  "musica",
		Artista: "mc pipokinha",
		Duracao: time.Minute*2 + time.Second*53,
	}

	musica2 := Musica{
		Titulo:  "musica",
		Artista: "anitta",
		Duracao: time.Minute*4 + time.Second*29,
	}

	musica3 := Musica{
		Titulo:  "sacanagem",
		Artista: "sua prima",
		Duracao: time.Minute*2 + time.Second*33,
	}

	playlist := Playlist{
		Nome:    "Minha Playlist",
		Musicas: []Musica{musica1, musica2, musica3},
	}

	ImprimirPlaylist(playlist)
}
