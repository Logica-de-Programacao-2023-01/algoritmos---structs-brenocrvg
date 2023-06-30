package main

import "fmt"

type Musica struct {
	titulo  string
	artista string
	duracao int
}

type Playlist struct {
	nome    string
	musicas []Musica
}

func buscarPlaylistsPorTitulo(titulo string, playlists []Playlist) []Playlist {
	resultado := []Playlist{}
	for _, playlist := range playlists {
		for _, musica := range playlist.musicas {
			if musica.titulo == titulo {
				resultado = append(resultado, playlist)
				break
			}
		}
	}
	return resultado
}

func main() {
	musicas1 := []Musica{
		{titulo: "Música 1", artista: "Artista 1", duracao: 180},
		{titulo: "Música 2", artista: "Artista 2", duracao: 240},
	}
	musicas2 := []Musica{
		{titulo: "Música 3", artista: "Artista 3", duracao: 210},
		{titulo: "Música 4", artista: "Artista 4", duracao: 190},
	}
	playlist1 := Playlist{nome: "Playlist 1", musicas: musicas1}
	playlist2 := Playlist{nome: "Playlist 2", musicas: musicas2}
	playlists := []Playlist{playlist1, playlist2}

	tituloBuscado := "Música 3"
	resultado := buscarPlaylistsPorTitulo(tituloBuscado, playlists)
	fmt.Println("Playlists com a música", tituloBuscado+":")
	for _, playlist := range resultado {
		fmt.Println(playlist.nome)
	}
}
