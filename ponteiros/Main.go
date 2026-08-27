package main

import "fmt"

type Carro struct {
	Name string
}

func (c *Carro) andou() {
	c.Name = "Volvo"
	fmt.Println(c.Name, "andou")
}

func main() {

	a := 10

	var ponteiro *int = &a

	fmt.Println(*ponteiro)
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Println("===========================")

	*ponteiro = 50
	fmt.Println(*ponteiro)
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Println("===========================")

	b := &a
	fmt.Println(*b)
	*b = 60
	fmt.Println(*b)
	fmt.Println(b)
	fmt.Println(&b)

	fmt.Println("===========================")
	*ponteiro = 50
	fmt.Println(*ponteiro)
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Println("===========================")
	variavel := 10
	abc(&variavel)
	fmt.Println(variavel)
	fmt.Println(&variavel)

	fmt.Println("===========================")

	carro := Carro{
		Name: "Fiat",
	}

	carro.andou()
	fmt.Println(carro.Name)
}

func abc(a *int) {
	*a = 200
}
