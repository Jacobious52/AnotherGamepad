package server // TODO(jacobious52): this should be project name folder in root. e.g. anothergamepad

import (
	"time"
)

type Event struct {
	Name string    `json:"name,omitempty"`
	Key  string    `json:"key"`
	Type string    `json:"type"`
	Time time.Time `json:"time,omitempty"`
}

func (e Event) State() string {
	switch e.Type {
	case "press":
		return "down"
	case "release":
		return "up"
	default:
		return "down"
	}
}
