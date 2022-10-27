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
	Status     int    `json:"status"`
	Code       string `json:"code"`
	State      string `json:"state"`
	Localidade string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
}

type BuscasCep struct {
	ViaCEP
	ApiCep
}

func main() {

	c1 := make(chan interface{})
	c2 := make(chan interface{})
	cep := "44007-200"

	go func() {
		result := BuscasCep.BuscaCep("http://viacep.com.br/ws/" + cep + "/json/")
		c1 <- result
	}()

	go func() {
		result := BuscasCep.BuscaApiCep("https://cdn.apicep.com/file/apicep/" + cep + ".json")
		c2 <- result
	}()

	select {
	case msg1 := <-c1:
		println("received from Viacep\n Cidade:", msg1)
	case msg2 := <-c2:
		println("received from Apicep\n Cidade:", msg2)
	case <-time.After(time.Second):
		println("timeout")
	}

}

func (b *BuscasCep) BuscaViaCep(url string) interface{} {
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
	data := []byte{}
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
	}
	elapsed := time.Since(start)
	fmt.Printf("%s execution took %s\n", url, elapsed)
	return data
}

func (b *BuscasCep) BuscaApiCep(url string) interface{} {
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

// func (b *BuscasCep) findApiData(url string, res []byte) FindInterface {
// 	if strings.Contains(url, "viacep.com.br") {
// 		var data ViaCEP
// 		err := json.Unmarshal(res, &data)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
// 		}
// 		return data
// 	}
// 	if strings.Contains(url, "cdn.apicep.com") {
// 		var data ApiCep
// 		err := json.Unmarshal(res, &data)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
// 		}
// 		return data
// 	}
// 	return nil
// }

// type FindInterface interface {
// 	findApiData()
// }
