package calculatermopi
import (
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

func AlgoritmoSpigotCalculaEnesimoTermoDePi(n int32) int32 {
	// começa a partir do primeiro decimal
	numeroTermos := n + 1
	limiteVetor := int32(math.Floor(10.0*float64(numeroTermos)/3.0))
	
	A := make([]int32, limiteVetor)
	num := make([]int32, limiteVetor)
	den := make([]int32, limiteVetor)
	digitos := make([]int32, numeroTermos)
	var idxInicializacao int32
	for idxInicializacao = 1; idxInicializacao <= limiteVetor; idxInicializacao++ {
		A[idxInicializacao - 1] = 2
		num[idxInicializacao - 1] = idxInicializacao - 1
		den[idxInicializacao - 1] = 2*idxInicializacao - 1
	}
	var digito, i, j int32
	for digito = 0; digito < int32(numeroTermos); digito++ {
		for j = 1; j <= limiteVetor; j++ {
			A[j - 1] = A[j - 1] * 10
		}

		var carryOut int32 = 0
		for i = limiteVetor - 1; i > 0; i-- {
			var ACarryOuted int32 = carryOut + A[i]
			A[i] = ACarryOuted % den[i]
			carryOut = int32(ACarryOuted / den[i]) * num[i]
		}

		ACarryOuted := carryOut + A[0]
		menorPotencia := math.Floor(math.Log10(float64(ACarryOuted)))
		digitos[digito] = int32(ACarryOuted / int32(math.Pow(10, menorPotencia)))
		A[0] = ACarryOuted % (digitos[digito] * int32(math.Pow(10, menorPotencia)))
	}

	return digitos[n]
}