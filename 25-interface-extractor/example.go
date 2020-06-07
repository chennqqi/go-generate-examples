//go:generate interface-extractor -type example.Bar
package example

import (
	"context"
)

type Bar struct {
}

func (b Bar) Run(ctx context.Context) error {
	return nil
}

func (b Bar) Add() {

}

func (b Bar) Stop() error {
	return nil
}
