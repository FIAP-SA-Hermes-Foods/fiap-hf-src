package valueObject

type Cpf struct {
	Value string `json:"value,omitempty"`
}

func (c Cpf) Validate() bool {
	if len(c.Value) < 18 {
		return false
	}
	return true
}
