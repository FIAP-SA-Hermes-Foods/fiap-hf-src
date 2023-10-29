package valueObject

import (
	"errors"
	"strings"
)

type Status struct {
	Value string `json:"value,omitempty"`
}

const (
	PaidStatus       = "Paid"
	CanceledStatus   = "Canceled"
	ReceivedStatus   = "Received"
	InProgressStatus = "In Progress"
	DoneStatus       = "Done"
	FinishedStatus   = "Finished"
)

var statusMap = map[string]string{
	"paid":        PaidStatus,
	"canceled":    CanceledStatus,
	"received":    ReceivedStatus,
	"in progress": InProgressStatus,
	"done":        DoneStatus,
	"finished":    FinishedStatus,
}

func (v *Status) Validate() error {

	status, ok := statusMap[strings.ToLower(v.Value)]

	if !ok {
		return errors.New("status is not valid")
	}

	v.Value = status

	return nil
}
