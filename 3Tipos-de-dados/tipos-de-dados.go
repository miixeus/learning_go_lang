package main

import (
	"errors"
	"fmt"
)

func main() {

	var numero int = 100000000000
	fmt.Println(numero)

	//uint // unsygned int é um int sem sinal

	var numero2 uint = 10000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 12345
	fmt.Println(numero3)

	//BYTE = UNT8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)

	//FIM NUMEROS REAIS

	var str string = "Texto" //Sempre aspas duplas
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'M'
	fmt.Println(char)

	//FIM STRINGS

	var texto string 
	fmt.Println(texto)
	
	//bOOLEANO

	var booleano bool
	fmt.Println(booleano)

	//ERROR

	var erro error
	fmt.Println(erro)

	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)
}