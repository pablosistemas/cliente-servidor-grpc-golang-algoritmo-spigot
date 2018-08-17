package api
import (
  "log"
  "golang.org/x/net/context"
  "cliente-servidor-grpc-golang-algoritmo-spigot/src"
)

type Server struct {
}

func (s *Server) GetTermoPi(ctx context.Context, in *PingRequest) (*PingResponse, error) {
  log.Printf("Receive message %d", in.TermoIndice)
  var termo int32 = calculatermopi.AlgoritmoSpigotCalculaEnesimoTermoDePi(in.TermoIndice)
  return &PingResponse{ TermoValor: termo }, nil
}