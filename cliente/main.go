package main
import (
  "tracksale.prova/estruturas/thread"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "tracksale.prova/api/gRpc"
  "strconv"
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