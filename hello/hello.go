package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) addAge() {
	p.age++
}

func getHelloMessage(name string) (msg string) {
	msg = "Hello " + name + "\n"
	return
}

type greeter interface {
	greet()
}

// Japanese is hoge.
type Japanese struct{}
type american struct{}

func (ja Japanese) greet() {
	fmt.Println("こんにちは")
}

func (us american) greet() {
	fmt.Println("Hello")
}

func main() {
	fmt.Printf("hello, world\n")

	p1 := person{"hoge", 10}
	p1.addAge()
	fmt.Printf("%#v\n", p1)

	fmt.Print(getHelloMessage("Hoge"))

	greeters := []greeter{Japanese{}, american{}}
	for _, greeter := range greeters {
		greeter.greet()
	}
}
