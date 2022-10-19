package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type ViaCEP struct {
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
	Status   int    `json:"status"`
	Code     string `json:"code"`
	State    string `json:"state"`
	City     string `json:"city"`
	District string `json:"district"`
	Address  string `json:"address"`
}

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		BuscaCep("viacep", "40301110")
		c1 <- 1
	}()

	go func() {
		BuscaCep("apicep", "40301-110")
		c2 <- 2
	}()

	select {
	case msg1 := <-c1:
		println("received from Viacep", msg1)
	case msg2 := <-c2:
		println("received from Apicep", msg2)
	case <-time.After(time.Second):
		println("timeout")
	}

}

func BuscaCep(site string, cep string) {
	start := time.Now()

	var req *http.Response
	var err error
	if site == "viacep" {
		req, err = http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
	}
	if site == "apicep" {
		req, err = http.Get("https://cdn.apicep.com/file/apicep/" + cep + ".json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
	}
	if site == "viacep" {
		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
		}
		fmt.Println("Cidade: ", data.Localidade)
		elapsed := time.Since(start)
		fmt.Printf("This execution took %s\n", elapsed)
	}
	if site == "apicep" {
		var data ApiCep
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
		}
		fmt.Println("Cidade: ", data.City)
		elapsed := time.Since(start)
		fmt.Printf("This execution took %s\n", elapsed)
	}
}
