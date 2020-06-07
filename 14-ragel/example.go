//go:generate ragel -Z -T0 -o atoi.go atoi.rl
package example

import (
	"fmt"
)

func examle() {
	var atoiTests = []atoiTest{
		atoiTest{"7", 7},
		atoiTest{"666", 666},
		atoiTest{"-666", -666},
		atoiTest{"+666", 666},
		atoiTest{"1234567890", 1234567890},
		atoiTest{"+1234567890\n", 1234567890},
		// atoiTest{"+ 1234567890", 1234567890}, // i will fail
	}

	for _, test := range atoiTests {
		res := atoi(test.s)
		if res != test.v {
			fmt.Printf("FAIL atoi(%#v) != %#v\n", test.s, test.v)
		}
	}
}
