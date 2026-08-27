package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome  string
	Email string
	CPF   int
}

func (c Cliente) Exibe() {
	fmt.Println("\nExibindo cliente pelo metodo:", c.Nome)

}

type ClienteInternacional struct {
	Cliente
	Pais string `json:"locale"`
}

func main() {

	cliente := Cliente{
		Nome:  "João",
		Email: "email@email.com",
		CPF:   12345678910,
	}

	fmt.Println(cliente)

	cliente2 := Cliente{"Marcio", "teste@email.com", 45678912302}
	fmt.Printf("\n Nome: %s. \n Email: %s. \n CPF: %d", cliente2.Nome, cliente2.Email, cliente2.CPF)

	fmt.Println("\n \n======================")
	fmt.Println("Cliente Internacional")
	fmt.Println("======================")

	cliente3 := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "Davi",
			Email: "email@teste.com",
			CPF:   25896314778,
		},
		Pais: "Holanda",
	}

	fmt.Printf("\n Nome: %s. \n Email: %s. \n CPF: %d \n Pais: %s. \n", cliente3.Nome, cliente3.Email, cliente3.CPF, cliente3.Pais)
	cliente.Exibe()
	cliente2.Exibe()
	cliente3.Exibe()

	fmt.Println("\n \n======================")
	fmt.Println("Convertendo struct em json")
	fmt.Println("======================")

	clienteJson, err := json.Marshal(cliente3)
	if err != nil {

		log.Fatal(err.Error())

	}

	fmt.Println(string(clienteJson))

	jsoncliente4 := `{"Nome":"Davi","Email":"email@teste.com","CPF":25896314778,"locale":"Holanda"}`
	cliente4 := ClienteInternacional{}

	json.Unmarshal([]byte(jsoncliente4), &cliente4)
	fmt.Println(cliente4)

}
