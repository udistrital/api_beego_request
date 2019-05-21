package test

import (
	"testing"

	"github.com/udistrital/api_beego_request/controllers"
)

func TestResta(t *testing.T) {
	valor := controllers.Resta(4, 2)
	if valor != 2 {
		t.Error("Se espera 4 y es obtuvo", valor)
		t.Fail()
	} else {
		t.Log("TestResta Finalizado Correctamente (OK)")
	}
}
