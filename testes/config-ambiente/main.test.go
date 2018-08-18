package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
    "os"

	"cliente-servidor-grpc-golang-algoritmo-spigot/estruturas/config"
)

func main() {
	configuration := config.Configuracao{}
	dir, erro := os.Getwd()
	if erro != nil {
        fmt.Print(erro)
	}
	fmt.Println(dir)

	b, erro1 := ioutil.ReadFile("../../config/config.development.json") // just pass the file name
    if erro1 != nil {
        fmt.Print(erro1)
    }

	str := string(b) // convert content to a 'string'

    fmt.Println(str)
	
	file, _ := os.Open("../../config/config.development.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s:%d", configuration.GRpc.Endereco, configuration.GRpc.Porta)
	fmt.Printf("%s:%d", configuration.ApiRest.Endereco, configuration.ApiRest.Porta)
}