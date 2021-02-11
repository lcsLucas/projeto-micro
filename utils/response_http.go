package utils

type ResponseHTTP struct {
	Status bool        `json:"status"`
	Erro   string      `json:"erro,omitempty"`
	Dados  interface{} `json:"dados,omitempty"`
}
