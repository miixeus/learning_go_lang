package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

func main() {
	fmt.Println("Arquivo structs")

	//Maneira de popular o struct
	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	//MANEIRA MAIS RÁPIDA DE POPULAR
	usuario2 := usuario{"Davi", 21}
	fmt.Println(usuario2)

	//Quando não tenho todos dados do usuário
	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)

}