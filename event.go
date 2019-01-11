package anothergamepad

type Event struct {
	Key  string `json:"key"`
	Type string `json:"type"`
	Time int    `json:"time,omitempty"`
}

func (e Event) State() string {
	return e.Type
}
