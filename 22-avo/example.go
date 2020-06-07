//go:generate go run asm.go -out add.s -stubs stub.go -pkg example

package example

func runCallAdd(x, y uint64) uint64 {
	return Add(x, y)
}
