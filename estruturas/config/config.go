package config

type Configuracao struct {
	GRpc struct {
		Endereco string `json:"Endereco"`
		Porta int `json:"Porta"`
	} `json:"gRpc"`
	ApiRest struct {
		Endereco string `json:"Endereco"`
		Porta int `json:"Porta"`
	} `json:"apiRest"`
}
