package main

import "fmt"

func main() {
	fmt.Println(Greet("World"))
}

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
