//go:generate go-atomicvalue -type "Pill<time.Time>"
package example

import (
	"sync/atomic"
)

type Pill atomic.Value
