package main

import (
	"fmt"
	cliente "golangoo/clientes"
	conta "golangoo/contas"
)

func main() {

	titularFilipe := cliente.Titular{Nome: "Filipe", Cpf: "123.456.789-00", Profissao: "Analista"}
	titularGabriel := cliente.Titular{Nome: "Gabriel", Cpf: "987.654.321-00", Profissao: "Desenvolvedor"}

	contaDoFilipe := conta.ContaCorrente{Titular: titularFilipe}
	ContaDoGabriel := conta.ContaCorrente{Titular: titularGabriel}

	ContaDoGabriel.Depositar(500)
	ContaDoGabriel.Transferir(200, &contaDoFilipe)

	fmt.Println(ContaDoGabriel.ObterSaldo())
	fmt.Println(contaDoFilipe.ObterSaldo())
}
