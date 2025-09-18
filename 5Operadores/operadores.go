package main

import "fmt"

func main() {
	//Aritmeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restDaDivisao )

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//Atribuições
	var variavel string = "String"
	variavel2 := "string"
	fmt.Println(variavel, variavel2)
	//Fim dos operadores de atribuição ("=", ":=")

	//Operadores relacionais SEMPRE retornam Booleanos (true/false), ex: (1 > 2)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	//Fim dos operacionais

	//Operadores Lógicos : &&(and ou igual), ||(ou), !(e a variavel muda o valor booleano)
	fmt.Println(true && true)
	fmt.Println(true || false)
	//FIM DOS OPERADORES LÓGICOS

	//OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero-=20
	fmt.Println(numero)
	//FIM DOS OPERADORES UNÁRIOS

	



	
}

