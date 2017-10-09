package main

import "fmt"

func main() {
	sayHello("Mahesh")
}

func sayHello(name string, name2 ...interface{}) {
	fmt.Println("Hello, " + name + "!")
}
