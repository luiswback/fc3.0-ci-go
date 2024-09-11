package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)
	if total != 30 {
		t.Errorf("Resultado errado: Resultado obtido %d. Valor esperado: %d", total, 30)
	}
}
