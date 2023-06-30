package main

import "fmt"

type Pessoa struct {
	nome     string
	idade    int
	endereço Endereço
}

type Endereço struct {
	rua    string
	numero int
	cidade string
	estado string
}

func main() {
	p := Pessoa{
		nome:  "pitbull",
		idade: 17,
		endereço: Endereço{
			rua:    "doideira",
			numero: 17,
			cidade: "Brasília",
			estado: "DF",
		},
	}
	fmt.Println("Seu nome é: ", p.nome)
	fmt.Println("Sua idade é: ", p.idade)
	fmt.Println("Sua rua é: Rua ", p.endereço.rua)
	fmt.Println("O número da sua casa é: ", p.endereço.numero)
	fmt.Println("Sua cidade é: ", p.endereço.cidade)
	fmt.Println("Seu estado é: ", p.endereço.estado)
}
