//go:generate go-syncmap -type "Pill<int, string>"
//go:generate go-syncmap -type "Pill<int, time.Time>"
//go:generate go-syncmap -type "Pill<int, encoding/json.Token>"
package example

import (
	"sync"
)

type Pill sync.Map //type def
