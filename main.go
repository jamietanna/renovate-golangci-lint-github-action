package main

import "fmt"

func foo() (string, error) {
	return "foo", fmt.Errorf("this won't be checked")
}

func main() {

	fmt.Println("vim-go")

	foo()
}
