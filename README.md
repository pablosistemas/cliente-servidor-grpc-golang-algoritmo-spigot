# Algoritmo Spigot 
O algoritmo [Spigot](https://www.maa.org/sites/default/files/pdf/pubs/amm_supplements/Monthly_Reference_12.pdf) permte calcular os valores das casas decimais de PI sem utilizar operações de ponto flutuante.

### Tutorial

O `Makefile` automatiza a instalação de dependências via `go get` e a construção dos binários. Todos os arquivos binários são gerados na pasta `bin`.

Constrói servidor gRPC, cliente gRPC e API REST:

```sh
make
```

Ou individualmente por:
```sh
make cliente
make servidor
make api
```

### Execução
Execute os comandos a partir do diretório do projeto `$GOPATH/src/cliente-servidor-grpc-golang-algoritmo-spigot`.

Servidor gRPC:
```sh
bin/servidor
```

API REST:
```sh
bin/apiRest
```

Cliente gRPC/REST para 50 casas decimais de PI:
```sh
bin/cliente -n 50
```
Ou
```sh
bin/cliente --help
Usage: cliente [--help] [-n value] [parameters ...]
     --help       Help
 -n, --num=value  Numero de casas decimais de PI

```

### Configuração

O arquivo `config/config.development.json` armazena o endereço e porta dos serviços gRPC e REST:
 
```json
{
    "gRpc": {
        "Endereco": "http://localhost",
        "Porta": 7777
    },
    "apiRest": {
        "Endereco": "http://localhost",
        "Porta": 8080
    }
}
```
