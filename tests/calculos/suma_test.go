package test

import (
	"testing"

	"github.com/udistrital/api_beego_request/calculos"
)

func TestSuma(t *testing.T) {
	valor := calculos.Suma(2, 2)
	if valor != 4 {
		t.Error("Se espera 4 y es obtuvo", valor)
		t.Fail()
	} else {
		t.Log("TestSuma Finalizado Correctamente (OK)")
	}
}
