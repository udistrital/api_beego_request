package calculos

import (
	"flag"
	"os"
	"testing"

	"github.com/udistrital/api_beego_request/calculos"
)

var parameters struct {
	Endpoint_1 string
}

func TestMain(m *testing.M) {

	parameters.Endpoint_1 = os.Getenv("Endpoint1")
	flag.Parse()
	os.Exit(m.Run())
}

func TestSuma(t *testing.T) {
	valor := calculos.Suma(2, 2)
	if valor != 4 {
		t.Error("Se espera 4 y es obtuvo", valor)
		t.Fail()
	} else {
		t.Log("TestSuma Finalizado Correctamente (OK)")
	}
}

func TestEndPointSuma(t *testing.T) {
	t.Log("AAAAAAAAAAAAAAAA")
	t.Log(parameters.Endpoint_1)
	t.Log("AAAAAAAAAAAAAAAA")
}
