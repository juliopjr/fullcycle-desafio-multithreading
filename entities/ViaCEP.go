package entities

func NewViaCEP(cep string) *viaCEP {
	return &viaCEP{Cep: cep}
}

type viaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (e *viaCEP) GetURL() string {
	return "http://viacep.com.br/ws/" + e.Cep + "/json/"
}
