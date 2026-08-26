package main

import "fmt"

func main() {

	// fmt.Println("Ola mundo")

	var abc string
	abc = "Rodrigo"

	fmt.Printf("%s ", abc)
	fmt.Printf("%T \n", abc)

	var numero int
	var numero2 int64
	var flutuante float64
	var boleano bool
	var texto string

	fmt.Printf("%T \n", numero)
	fmt.Printf("%T \n", numero2)
	fmt.Printf("%T \n", flutuante)
	fmt.Printf("%T \n", boleano)
	fmt.Printf("%T \n", texto)

	a := "Marcos"
	a = "Sara"

	// fmt.Println(a)

	b := 10
	c := 3.144
	d := false
	e := `uooou
	legal`

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)

}
