package main

import (
	"fmt"
	"time"
)

type Funcionario struct {
	nome    string
	salario float64
	idade   int
}

func aumentar(f *Funcionario, porcentagem float64) {
	f.salario *= (1 + porcentagem/100)
}

func diminuir(f *Funcionario, porcentagem float64) {
	f.salario *= (1 - porcentagem/100)
}

func tempos(f Funcionario) int {
	idadeInicioTrabalho := 18
	tempoServico := time.Now().Year() - f.idade - idadeInicioTrabalho
	return tempoServico
}

func main() {
	funcionario := Funcionario{nome: "João", salario: 5000, idade: 30}
	aumentar(&funcionario, 10)
	fmt.Println("Salário após aumento:", funcionario.salario)
	diminuir(&funcionario, 5)
	fmt.Println("Salário após redução:", funcionario.salario)
	tempoServico := tempos(funcionario)
	fmt.Println("Tempo de serviço:", tempoServico, "anos")
}
