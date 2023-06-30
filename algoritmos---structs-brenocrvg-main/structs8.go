package main

import (
	"fmt"
)

type Viagem struct {
	Origem  string
	Destino string
	Data    string
	Preco   float64
}

func cara(viagens []Viagem) (Viagem, error) {
	if len(viagens) == 0 {
		return Viagem{}, fmt.Errorf("slice de viagens vazio")
	}

	viagemMaisCara := viagens[0]

	for _, viagem := range viagens {
		if viagem.Preco > viagemMaisCara.Preco {
			viagemMaisCara = viagem
		}
	}

	return viagemMaisCara, nil
}

func main() {
	viagens := []Viagem{
		{Origem: "São Paulo", Destino: "Rio de Janeiro", Data: "2023-07-01", Preco: 5003.0},
		{Origem: "Rio de Janeiro", Destino: "Florianópolis", Data: "2023-07-05", Preco: 75045.0},
		{Origem: "São Paulo", Destino: "Salvador", Data: "2023-07-10", Preco: 10998.0},
	}

	viagemMaisCara, err := cara(viagens)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Viagem mais cara:")
		fmt.Println("Origem:", viagemMaisCara.Origem)
		fmt.Println("Destino:", viagemMaisCara.Destino)
		fmt.Println("Data:", viagemMaisCara.Data)
		fmt.Println("Preço:", viagemMaisCara.Preco)
	}
}
