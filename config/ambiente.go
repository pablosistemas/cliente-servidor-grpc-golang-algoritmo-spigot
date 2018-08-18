package ambiente

import (
	"encoding/json"
	"fmt"
    "os"

	"cliente-servidor-grpc-golang-algoritmo-spigot/estruturas/config"
)

func CarregaConfiguracoesAmbiente(caminhoArquivo string) config.Configuracao {
	configuracao := config.Configuracao{}
	arquivoConfig, _ := os.Open(caminhoArquivo)
	defer arquivoConfig.Close()
	decoder := json.NewDecoder(arquivoConfig)
	err := decoder.Decode(&configuracao)
	if err != nil {
		fmt.Println(err)
	}
	return configuracao
}