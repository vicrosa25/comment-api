package main

import "fmt"

// Run - is going to be responsible for
// the instantiation and startup of our
// application
func Run() error {
	fmt.Println("Starting application...")
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
