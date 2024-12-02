package main

import "fmt"

// Run - is going to be responsive for
// the instantiation and startup of our
// go application
func Run() error {
	fmt.Println("Starting up the application")
	return nil
}

func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
