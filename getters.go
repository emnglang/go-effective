/*
*/
package main

import "fmt"

type Person struct {
	name string
}

func (p Person) Name() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func main() {
	me := Person{"Roger"}
	fmt.Println("My name is: ", me.Name())

	me.SetName("Rafal")
	fmt.Println("My new name is: ", me.Name())
}
