package main
import (
  "cliente-servidor-grpc-golang-algoritmo-spigot/estruturas/thread"
  "cliente-servidor-grpc-golang-algoritmo-spigot/api/gRpc"

  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "strconv"
  "bytes"
  "log"
  "os"
)

func calculaTermoDePi(nThread int, canal chan thread.Thread, clienteApi api.PingClient) {
  resultado, err := clienteApi.GetTermoPi(context.Background(), &api.PingRequest{TermoIndice: int32(nThread)})
  if err != nil {
    log.Fatalf("Error when calling GetTermoPi: %s", err)
  }
  canal <- thread.Thread{int32(nThread), resultado.TermoValor}
}

func retornaStringAPartirVetor(valor []int32) {
  var stringValor bytes.Buffer
  stringValor.WriteString(strconv.Itoa(valor[0]))
  stringValor.WriteString(",")
  for i := 1; i < len(valor); i++ {
    stringValor.WriteString(strconv.Itoa(valor[i]))
  }
  return stringValor.String()
}

func enviaValorAtualDePi(valorAtual []int32) {
	data := url.Values{}
	data.Set("valor_pi", retornaStringAPartirVetor(valorAtual))
	req, err := http.NewRequest("POST", "http://localhost:8080/status", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	cliente := &http.Client{}
	resp, err := cliente.Do(req)
    if err != nil {
        panic(err)
    }
	defer resp.Body.Close()
}

func main() {
  var conexaoRpc *grpc.ClientConn
  conexaoRpc, err := grpc.Dial(":7777", grpc.WithInsecure())


  canal := make(chan thread.Thread)

  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }

  defer conexaoRpc.Close()
  clienteApi := api.NewPingClient(conexaoRpc)
  numeroTermos, err := strconv.Atoi(os.Args[1])

  if err != nil {
	  os.Exit(1)
  }

  digitos := make([]string, numeroTermos + 1)
  digitos[0] = "3"

  for i := 1; i <= numeroTermos; i++ {
    go calculaTermoDePi(i, canal, clienteApi)
    resultado := <- canal 
		digitos[resultado.Indice] = strconv.Itoa(int(resultado.Valor))
    


    // log.Printf("Response from server: %d", resultado.Valor)
  }
  log.Println(digitos)
}