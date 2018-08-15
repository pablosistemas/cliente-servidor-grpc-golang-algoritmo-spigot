package main

import (
    //"encoding/json"
    //"log"
    //"net/http"
	//"github.com/gorilla/mux"
	"fmt"
	"math"
)

func calcularPiBaileyBorweinPlouffeFormula(n int) (sum float64) {
	sum = 0
	for k := 0.0; k < float64(n); k+=1.0 {
		sum += 1/math.Pow(16, k) * (4/(8*k + 1) - 2/(8*k + 4) - 1/(8*k + 5) - 1/(8*k + 6))
	}
	return
}

func initializarVetor(n int) ([]int) {
	limiteVetor := int(10.0*float64(n)/3.0)
	v := make([]int, limiteVetor)
	for i := 0; i < limiteVetor; i++ {
		v[i] = 2
	}
	return v
}

func algoritmoSpigotCalculaEnesimoTermoDePi(n int) []int {
	
	limiteVetor := int(math.Floor(10.0*float64(n)/3.0))
	
	A := make([]int, limiteVetor)
	num := make([]int, limiteVetor)
	den := make([]int, limiteVetor)
	digitos := make([]int, n)

	for idxInicializacao := 1; idxInicializacao <= limiteVetor; idxInicializacao++ {
		A[idxInicializacao - 1] = 2
		num[idxInicializacao - 1] = idxInicializacao - 1
		den[idxInicializacao - 1] = 2*idxInicializacao - 1
	}

	for digito := 0; digito < n; digito++ {
		for j := 1; j <= limiteVetor; j++ {
			A[j - 1] = A[j - 1] * 10
		}

		carryOut := 0
		for i := limiteVetor - 1; i > 0; i-- {
			ACarryOuted := carryOut + A[i]
			A[i] = ACarryOuted % den[i]
			carryOut = int(ACarryOuted / den[i]) * num[i]
		}

		ACarryOuted := carryOut + A[0]
		menorPotencia := math.Floor(math.Log10(float64(ACarryOuted)))
		digitos[digito] = int(ACarryOuted / int(math.Pow(10, menorPotencia)))
		A[0] = ACarryOuted % (digitos[digito] * int(math.Pow(10, menorPotencia)))
	}
	
	return digitos
}

func main() {
	/*router := mux.NewRouter()
	router.HandleFunc("/calculaTermoDePi/{termo}", GetTermo).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))*/
	fmt.Println(algoritmoSpigotCalculaEnesimoTermoDePi(4))
}