package main

import (
	"cliente-servidor-grpc-golang-algoritmo-spigot/src"
	"cliente-servidor-grpc-golang-algoritmo-spigot/estruturas/thread"
	"fmt"
	"strconv"
)


func calculaTermoDePi(nThread int, c chan thread.Thread) {
	resultado := calculatermopi.AlgoritmoSpigotCalculaEnesimoTermoDePi(int32(nThread))
	c <- thread.Thread {int32(nThread), resultado}
}

func main() {
	n := 25
	canal := make(chan thread.Thread)
	digitos := make([]string, n + 1)

	digitos[0] = "3"

	for nThread := 1; nThread <= n; nThread++ {
		go calculaTermoDePi(nThread, canal)
		resultado := <- canal 
		digitos[resultado.Indice] = strconv.Itoa(int(resultado.Valor))
	}

	fmt.Println(digitos)
}