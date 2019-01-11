package anothergamepad

import (
	"time"
)

type Event struct {
	Key  string    `json:"key"`
	Type string    `json:"type"`
	Time time.Time `json:"time,omitempty"`
}

func (e Event) State() string {
	return e.Type
}
