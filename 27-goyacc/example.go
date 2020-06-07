//go:generate goyacc -p "expr" example.y
package example

//&& 6g y.go && 6l -o goexpr y.6

func example() {
	main()
}
