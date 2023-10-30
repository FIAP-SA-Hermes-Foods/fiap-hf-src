package valueObject

import "time"

type DeactivatedAt struct {
	Value time.Time `json:"value,omitempty"`
}

var deactivatedAtFormatLayout string = `02-01-2006 15:04:05`

func (d *DeactivatedAt) Format() string {
	f := d.Value.Format(deactivatedAtFormatLayout)
	if f == "01-01-0001 00:00:00" {
		return "null"
	}
	return f
}

func (d *DeactivatedAt) DurationFromString(du string) error {
	return nil
}
