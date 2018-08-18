package main

import (
    "log"
    "fmt"
    "net/http"

    "cliente-servidor-grpc-golang-algoritmo-spigot/config"
)

func main() {
  config := ambiente.CarregaConfiguracoesAmbiente("config/config.development.json")
  router := NewRouter()

  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiRest.Porta), router))
}