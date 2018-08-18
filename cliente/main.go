package main
import (
  "cliente-servidor-grpc-golang-algoritmo-spigot/estruturas/thread"
  "cliente-servidor-grpc-golang-algoritmo-spigot/api/gRpc"
  "cliente-servidor-grpc-golang-algoritmo-spigot/config"

  "github.com/pborman/getopt"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "net/http"
  "net/url"
  "strconv"
  "strings"
  "bytes"
  "log"
  "fmt"
  "os"
)

func calculaTermoDePi(nThread int, canal chan thread.Thread, clienteApi api.PingClient) {
  resultado, err := clienteApi.GetTermoPi(context.Background(), &api.PingRequest{TermoIndice: int32(nThread)})
  if err != nil {
    log.Fatalf("Error when calling GetTermoPi: %s", err)
  }
  canal <- thread.Thread{int32(nThread), resultado.TermoValor}
}

func retornaStringAPartirVetor(valor []string) (string) {
  var stringValor bytes.Buffer
  stringValor.WriteString(valor[0])
  stringValor.WriteString(",")
  for i := 1; i < len(valor); i++ {
    stringValor.WriteString(valor[i])
  }
  return stringValor.String()
}

func enviaValorAtualDePi(urlApi string, valorAtual []string) {
	data := url.Values{}
	data.Set("valor_pi", retornaStringAPartirVetor(valorAtual))
	req, err := http.NewRequest("POST", urlApi, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	cliente := &http.Client{}
	resp, err := cliente.Do(req)
  if err != nil {
      panic(err)
  }
	defer resp.Body.Close()
}


func main() {
  optNumeroTermos := getopt.StringLong("num", 'n', "", "Numero de casas decimais de PI")
  optHelp := getopt.BoolLong("help", 0, "Help")
  getopt.Parse()

  if *optHelp || *optNumeroTermos == "" {
    getopt.Usage()
    os.Exit(0)
  }

  config := ambiente.CarregaConfiguracoesAmbiente("config/config.development.json")
  urlApiRest := fmt.Sprintf("%s:%d/status", config.ApiRest.Endereco, config.ApiRest.Porta)

  var conexaoRpc *grpc.ClientConn
  conexaoRpc, err := grpc.Dial(fmt.Sprintf(":%d", config.GRpc.Porta), grpc.WithInsecure())

  canal := make(chan thread.Thread)

  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }

  defer conexaoRpc.Close()
  clienteApi := api.NewPingClient(conexaoRpc)
  numeroTermos, err := strconv.Atoi(*optNumeroTermos)

  if err != nil {
	  os.Exit(1)
  }

  digitos := make([]string, numeroTermos + 1)
  digitos[0] = "3"

  for i := 1; i <= numeroTermos; i++ {
    go calculaTermoDePi(i, canal, clienteApi)
    resultado := <- canal 
    digitos[resultado.Indice] = strconv.Itoa(int(resultado.Valor))
    enviaValorAtualDePi(urlApiRest, digitos)
  }
  log.Printf("Valor de PI: %s\n", retornaStringAPartirVetor(digitos))
}