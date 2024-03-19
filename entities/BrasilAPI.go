package entities

func NewBrasilAPI(cep string) *brasilAPI {
	return &brasilAPI{Cep: cep}
}

type brasilAPI struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func (e *brasilAPI) GetURL() string {
	return "https://brasilapi.com.br/api/cep/v1/" + e.Cep
}
