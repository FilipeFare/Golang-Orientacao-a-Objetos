package contas

import Clientes "golangoo/clientes"

type ContaCorrente struct {
    Titular       Clientes.Titular
    NumeroAgencia int
    NumeroConta   int
    saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
    podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
    if podeSacar {
        c.saldo -= valorDoSaque
        return "Saque realizado com sucesso"
    } else {
        return "saldo insuficiente"
    }
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
    if valorDoDeposito > 0 {
        c.saldo += valorDoDeposito
        return "Deposito realizado com sucesso", c.saldo
    } else {
        return "Valor do deposito menor que zero", c.saldo
    }
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
    if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
        c.saldo -= valorDaTransferencia
        contaDestino.Depositar(valorDaTransferencia)
        return true
    } else {
        return false
    }
}

func (c *ContaCorrente) ObterSaldo() float64 {
    return c.saldo
}