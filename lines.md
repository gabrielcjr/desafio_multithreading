
// func ConsultaCep(adapter AdapterInterface) {
//  var data struct{}
//  err = json.Unmarshal(res, &data)
//  if err != nil {
//   fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
//  }
// elapsed := time.Since(start)
// fmt.Printf("%s execution took %s\n", url, elapsed)
//  fmt.Print(data)
//  return data
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

// consigo injetar o viacep no busca cep pq ele implementa a interface

// func (b *BuscasCep) BuscaApiCep(url string) struct{} {
//  start := time.Now()
//  req, err := http.Get(url)
//  if err != nil {
//   fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
//  }
//  defer req.Body.Close()
//  res, err := io.ReadAll(req.Body)
//  if err != nil {
//   fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
//  }
//  var data struct{}
//  err = json.Unmarshal(res, &data)
//  if err != nil {
//   fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
//  }
//  elapsed := time.Since(start)
//  fmt.Printf("%s execution took %s\n", url, elapsed)
//  return data
// }
