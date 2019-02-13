package rc

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

func init() {
	log.SetHandler(text.Default)
	log.SetLevel(log.DebugLevel)
}
