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

type BuscasCep struct {
	ViaCep
	ApiCep
}

type AdapterInterface interface {
	buscaCep(cep string)
}

func main() {

	c1 := make(chan struct{})
	c2 := make(chan struct{})
	cep := "44007-200"

	buscaCep := new(BuscasCep)

	go func() {
		result := buscaCep.BuscaVia("http://viacep.com.br/ws/" + cep + "/json/")
		fmt.Print(result)
		c1 <- result
	}()

	go func() {
		result := buscaCep.BuscaApi("https://cdn.apicep.com/file/apicep/" + cep + ".json")
		fmt.Print(result)
		c2 <- result
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

func (b *BuscasCep) BuscaVia(adapter AdapterInterface) struct{} {
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
	var data struct{}
	err = json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
	}
	elapsed := time.Since(start)
	fmt.Printf("%s execution took %s\n", url, elapsed)
	fmt.Print(data)
	return data
}

func (v *ViaCep) buscaCep(cep string) {

}

// func (b *BuscasCep) BuscaApiCep(url string) struct{} {
// 	start := time.Now()
// 	req, err := http.Get(url)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
// 	}
// 	defer req.Body.Close()
// 	res, err := io.ReadAll(req.Body)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
// 	}
// 	var data struct{}
// 	err = json.Unmarshal(res, &data)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
// 	}
// 	elapsed := time.Since(start)
// 	fmt.Printf("%s execution took %s\n", url, elapsed)
// 	return data
// }

// cria uma interface q entra na funcao
// func consultaCep(adapter AdapterInterface) {.....
// type AdapterInterface interface {
//   buscaCep(cep string).....
// }
// type VIACEP struct {
// //....
// }

// func (v *VIACEP) buscaCep(cep string)

// viacep := NeViaCEP()
// buscaCep(viacep)
