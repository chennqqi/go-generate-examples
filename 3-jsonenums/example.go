//go:generate jsonenums -type=Pill
package example

type Pill int //type def

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
