//go:generate enumer -type=MemberState -json -text -yaml

// json
// text
// yaml
// sql
// enumer -type=MemberState -json -transform=snake
package example

type MemberState uint

const (
	STARUP     MemberState = iota
	PRIMARY
	SECONDARY
	RECOVERING       
	DOWN
	STARTUP2
	UNKNOWN
	ARBITER
	REMOVED
	ROLLBACK
	STALE
)
