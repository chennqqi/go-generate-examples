//go:generate easyjson -all example.go

package example

import (
	"time"
)

type Person struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Birthday time.Time `json:"persion"`
}
