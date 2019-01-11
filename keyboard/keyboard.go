package keyboard

import (
	"strings"

	"github.com/go-vgo/robotgo"
	log "github.com/sirupsen/logrus"
)

// KeyToggle sends a key toggle request to the system of type `key` and can configure holding
// Valid keys found here https://github.com/go-vgo/robotgo/blob/master/docs/keys.md
// Valid states are "up" or "down"
func KeyToggle(key string, state string) {
	keys := strings.Split(key, " ")
	for _, k := range keys {
		log.WithFields(log.Fields{
			"key":   k,
			"state": state,
		}).Infoln("key toggle")
		robotgo.KeyToggle(k, state)
	}
}
