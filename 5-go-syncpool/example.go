//go:generate go-syncpool -type "Pill<string>"
package example

import (
	"sync"
)

type Pill sync.Pool
