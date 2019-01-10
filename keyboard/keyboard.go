package keyboard

import (
	"github.com/go-vgo/robotgo"
	log "github.com/sirupsen/logrus"
)

// KeyToggle sends a key toggle request to the system of type `key` and can configure holding
// Valid keys found here https://github.com/go-vgo/robotgo/blob/master/docs/keys.md
// Valid states are "up" or "down"
func KeyToggle(key string, state string) {
	log.WithFields(log.Fields{
		"key":   key,
		"state": state,
	}).Trace("A walrus appears")
	robotgo.KeyToggle(key, state)
}

// KeyTap sends a key tap request to the system of type `key`
// Valid keys found here https://github.com/go-vgo/robotgo/blob/master/docs/keys.md
func KeyTap(key string) {
	log.WithFields(log.Fields{
		"key": key,
	}).Trace("A walrus appears")
	robotgo.KeyTap(key)
}
