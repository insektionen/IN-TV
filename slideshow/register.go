package slideshow

import (
	"github.com/insektionen/IN-TV/v1/viewmodels"
	"time"
)

var ConnectedClients = map[string]int64{}

func RegisterClient(reg *viewmodels.Register) {
	ConnectedClients[reg.Name] = time.Now().Unix()
}
