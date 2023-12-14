package valueObject

type PaymentPrice struct {
	Value float64 `json:"value,omitempty"`
}

func (p PaymentPrice) Validate() error {

	return nil
}
