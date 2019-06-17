package controllers

import (
	"flag"
	"os"
	"testing"

	"github.com/udistrital/api_beego_request/controllers"
)

var parameters struct {
	Endpoint2 string
}

func TestMain(m *testing.M) {
	parameters.Endpoint2 = os.Getenv("Endpoint2")
	flag.Parse()
	os.Exit(m.Run())
}

func TestResta(t *testing.T) {
	valor := controllers.Resta(4, 2)
	if valor != 2 {
		t.Error("Se espera 4 y se obtuvo", valor)
		t.Fail()
	} else {
		t.Log("TestResta Finalizado Correctamente (OK)")
	}
}

func TestEndPointResta(t *testing.T) {
	t.Log("AAAAAAAAAAAAAAAA")
	t.Log(parameters.Endpoint2)
	t.Log("AAAAAAAAAAAAAAAA")
}
