package dto

type UserInput struct {
	CPF          string `json:"cpf,omitempty"`
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	WantRegister bool   `json:"wantRegister"`
}

type UserOutput struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Token      string `json:"token,omitempty"`
}
