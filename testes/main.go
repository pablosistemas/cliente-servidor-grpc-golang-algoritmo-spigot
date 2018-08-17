package main

import (
	"tracksale.prova/src"
	"tracksale.prova/estruturas"
	"fmt"
	"strconv"
)


func calculaTermoDePi(nThread int, c chan estruturas.Thread) {
	resultado := calculatermopi.AlgoritmoSpigotCalculaEnesimoTermoDePi(int32(nThread))
	c <- estruturas.Thread {int32(nThread), resultado}
}

func main() {
	n := 25
	canal := make(chan estruturas.Thread)
	digitos := make([]string, n + 1)

	digitos[0] = "3"

	for nThread := 1; nThread <= n; nThread++ {
		go calculaTermoDePi(nThread, canal)
		resultado := <- canal 
		digitos[resultado.Indice] = strconv.Itoa(int(resultado.Valor))
	}

	fmt.Println(digitos)
}