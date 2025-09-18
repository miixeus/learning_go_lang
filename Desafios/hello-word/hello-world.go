package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)


func main () {

	//Entrada de dados
	reader := bufio.NewReader(os.Stdin)

    fmt.Print("Digite o seu nome: ")
    nome, _:= reader.ReadString('\n')
    nome = strings.TrimSpace(nome) // remove o ENTER no final

    fmt.Print("Digite sua idade: ")
    idadeStr, _ := reader.ReadString('\n')
    idadeStr = strings.TrimSpace(idadeStr)
    idade, _ := strconv.Atoi(idadeStr) // converte para int

    fmt.Print("Digite sua cidade: ")
    cidade, _ := reader.ReadString('\n')
    cidade = strings.TrimSpace(cidade)

    fmt.Printf("Olá, %s! Você tem %d anos e mora em %s.\n", nome, idade, cidade)

		// Um extra: verificar se é maior de idade
	if idade >= 18 {
		fmt.Println("Você já é maior de idade ✅")
	} else {
		fmt.Println("Você ainda é menor de idade 🚸")
	}
}
