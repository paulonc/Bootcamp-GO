package main

import "testing"

func TestSoma(t *testing.T) {
	resultado := soma(1, 2, 3)
	esperado := 6
	if resultado != esperado {
		t.Errorf("Soma: esperado %d, mas obteve %d", esperado, resultado)
	}
}

func TestSubtracao(t *testing.T) {
	resultado := subtracao(40, 20, 10)
	esperado := 10
	if resultado != esperado {
		t.Errorf("Subtração: esperado %d, mas obteve %d", esperado, resultado)
	}
}

func TestMultiplicacao(t *testing.T) {
	resultado := multiplicacao(10, 3, 4)
	esperado := 120
	if resultado != esperado {
		t.Errorf("Multiplicação: esperado %d, mas obteve %d", esperado, resultado)
	}
}

func TestDivisao(t *testing.T) {
	resultado, err := divisao(100, 5, 2)
	esperado := float64(10)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if resultado != esperado {
		t.Errorf("Divisão: esperado %f, mas obteve %f", esperado, resultado)
	}
}

func TestDivisaoPorZero(t *testing.T) {
	resultado, err := divisao(100, 0)
	if err == nil {
		t.Error("Esperava um erro, mas nenhum ocorreu")
	}
	if resultado != 0.0 {
		t.Errorf("Resultado incorreto. Esperado: 0.0, Obtido: %f", resultado)
	}
}
