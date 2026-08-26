package main

import "fmt"

type Carro struct {
	Nome string
}

func (c Carro) andar() {
	fmt.Println(c.Nome, "Andou")

}

func main() {

	carro := Carro{
		Nome: "Fiat",
	}

	carro.andar()

	resultado2 := func(n ...int) func() int {

		res := 0

		for _, v := range n {
			res += v
		}

		return func() int {
			return res * res
		}

	}

	fmt.Println("Resultado do metodo resultado 2:", resultado2(1, 2, 3, 45, 80)())

	// ==========================================

	resultado := sumAll(1, 2, 5, 10, 45, 80)

	fmt.Println("Resultado do metodo soma:", resultado)

}

func soma(a int, b int) (result int) {

	result = a + b
	return
}

func sumAll(n ...int) int {
	result := 0

	for _, v := range n {
		result += v
	}

	return result
}
