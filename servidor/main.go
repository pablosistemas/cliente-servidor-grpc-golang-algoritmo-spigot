package main

import (
  "fmt"
  "log"
  "net"

  "cliente-servidor-grpc-golang-algoritmo-spigot/api/gRpc"
  "cliente-servidor-grpc-golang-algoritmo-spigot/config"
  "google.golang.org/grpc"
)

func main() {
  config := ambiente.CarregaConfiguracoesAmbiente("config/config.development.json")
  lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GRpc.Porta))
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := api.Server{}
  grpcServer := grpc.NewServer()
  api.RegisterPingServer(grpcServer, &s)
  if err := grpcServer.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %s", err)
  }
}