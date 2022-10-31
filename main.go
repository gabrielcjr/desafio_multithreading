package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type ApiCep struct {
	Status     int    `json:"status"`
	Code       string `json:"code"`
	State      string `json:"state"`
	Localidade string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
}

type AdapterInterface interface {
	buscaCep(url string) interface{}
}

func main() {

	c1 := make(chan interface{})
	c2 := make(chan interface{})
	cep := "44007-200"

	viaApi := new(ApiCep)
	apiCep := new(ViaCep)

	go func() {
		consultaCep(viaApi, "http://viacep.com.br/ws/"+cep+"/json/")
		c1 <- 1
	}()

	go func() {
		consultaCep(apiCep, "https://cdn.apicep.com/file/apicep/"+cep+".json")
		c2 <- 2
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("received from Viacep\n Cidade:", msg1)
	case msg2 := <-c2:
		fmt.Println("received from Apicep\n Cidade:", msg2)
	case <-time.After(time.Second):
		println("timeout")
	}
}

func (v *ViaCep) buscaCep(url string) interface{} {
	start := time.Now()
	req, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
	}
	var data ViaCep
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
	}
	elapsed := time.Since(start)
	fmt.Printf("%s execution took %s\n", url, elapsed)
	return data
}

func (a *ApiCep) buscaCep(url string) interface{} {
	start := time.Now()
	req, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
	}
	var data ApiCep
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
	}
	elapsed := time.Since(start)
	fmt.Printf("%s execution took %s\n", url, elapsed)
	return data
}

func consultaCep(adapter AdapterInterface, url string) {
	adapter.buscaCep(url)
}
