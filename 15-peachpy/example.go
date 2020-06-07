//go:generate python -m peachpy.x86_64 dot_product.py -S -o dot_product_amd64.s -mabi=goasm
package example

import (
	"fmt"
)

func DotProduct(x *float32, y *float32, length uint) float32

func main() {
	x := make([]float32, 2048)
	y := make([]float32, len(x))
	for i := 0; i < len(x); i++ {
		x[i] = 2.0
		y[i] = 3.0
	}

	z := DotProduct(&x[0], &y[0], uint(len(x)))

	fmt.Println("hello world")
	fmt.Println("z =", z)
}
