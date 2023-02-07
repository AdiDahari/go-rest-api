package main

import "fmt"

func Run() error {
	fmt.Println("Starting up application")
	return nil
}

func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}