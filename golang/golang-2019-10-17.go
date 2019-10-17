//------------------------------------------------
//Meus estudos da linguagem golang
//
//------------------------------------------------
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	funVarConst()
	funcPausa()

	fmt.Println(funcSimples(10, 10))
	funcPausa()

	fmt.Println(funcQuadradoCubo(4))
	funcPausa()

	fmt.Println(funcQuadradoCuboNome(5))
	funcPausa()

	funcOperMat()
	funcPausa()

	funcRelat()
	funcPausa()

	funcLogico()
	funcPausa()

	funcRetFor()
	funcPausa()

	funcWhileSQN()
	funcPausa()

	funcSwitch()
	funcPausa()

	funcSwitchTrue()
	funcPausa()

	funcVetor()
	funcPausa()

	funcSlice()
	funcPausa()

	funcArrayString()

}

//---------------------------------------------------------------------------------------------
// Variaveis e constantes
func funcPausa() {
	fmt.Print("Pressione 'ENTER' para continuar")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

//---------------------------------------------------------------------------------------------
// Variaveis e constantes
func funVarConst() {

	var idade int = 35
	fmt.Println(idade)

	idade1 := 36
	fmt.Println(idade1)

	var idade2 = 35
	fmt.Println(idade2)

	nome := "Rodrigo"
	fmt.Println("O meu nome e", nome)

	//Declaração de varias
	var peso, altura int = 76, 168

	fmt.Println("Minha altura e", altura, "e o meu peso e", peso)

	//Variavel em block
	var (
		carro             = "Fusca"
		km                = 1000000
		marca      string = "VW"
		fabricacao int    = 1980
		//nao pode usar := dentro do var
	)

	fmt.Println("Tenho um", carro, "ano", fabricacao, "com", km, "km, da marca", marca)

	variavelInt, variavelFloat, variaveisString := 30, 10.22, "Rodrigo"
	fmt.Println("Variavel Int", variavelInt)
	fmt.Println("Variavel Float", variavelFloat)
	fmt.Println("Variavel String", variaveisString)

	//Tipo de Variaveis
	var altura2 = 1.34
	aberto := true
	fmt.Printf("Tipo: %T Valor: %v\n", altura2, altura2)

	fmt.Printf("Tipo: %T Valor: %v\n", aberto, aberto)

	//Concatenar string
	ola := "Ola"
	mundo := "Mundo"
	fmt.Println(ola + " " + mundo)

	//Conversao de variaveis primitivas
	var varInteira int = 100
	var varFloat float64 = 5.3

	var newVarInt int = int(varFloat)
	var newVarFloat float64 = float64(varInteira)

	fmt.Println("inteiro", varInteira, "Float", varFloat, "Convertida Int", newVarInt, "Convertida float", newVarFloat)

	//constantes
	const Pi = 3.14

	fmt.Println("O valor de Pi e", Pi)
}

//---------------------------------------------------------------------------------------------
// Funcão com um retorno
func funcSimples(valorA int, valorB int) int {
	var resultado = valorA + valorB
	return resultado
}

//---------------------------------------------------------------------------------------------
// Funcão com varios retornos
func funcQuadradoCubo(a int) (int, int) {
	var quadrado int = a * a
	var cubo int = a * a * a
	return quadrado, cubo
}

//---------------------------------------------------------------------------------------------
// Funcão com varios retornos nomeados
func funcQuadradoCuboNome(a int) (quadrado int, cubo int) {
	quadrado = a * a
	cubo = a * a * a
	return quadrado, cubo
}

//---------------------------------------------------------------------------------------------
// Operadores matematicos
func funcOperMat() {
	var a int = 25
	var b int = 5
	fmt.Println("---------------------------------------------")
	fmt.Println("Valor de A:", a)
	fmt.Println("Valor da B:", b)
	fmt.Println("---------------------------------------------")
	fmt.Println("A soma de A+B.........:", a+b)
	fmt.Println("A subtração de A-B....:", a-b)
	fmt.Println("A divisao de A/B......:", a/b)
	fmt.Println("A Multiplicação de A*B:", a*b)
	fmt.Println("A Resto divisao de A%B:", a/b)
	fmt.Println("---------------------------------------------")

}

//---------------------------------------------------------------------------------------------
//Operadores de relacionamentos
func funcRelat() {
	var a int = 25
	var b int = 5
	fmt.Println("---------------------------------------------")
	fmt.Println("Valor de A:", a)
	fmt.Println("Valor da B:", b)
	fmt.Println("---------------------------------------------")
	fmt.Println("Maior A>B...........:", a > b)
	fmt.Println("Maior igual A>=B....:", a >= b)
	fmt.Println("Menor A<B...........:", a < b)
	fmt.Println("Menor igual A<=B....:", a <= b)
	fmt.Println("Igual A==B..........:", a == b)
	fmt.Println("Diferente A!=B......:", a != b)
	fmt.Println("---------------------------------------------")

}

//---------------------------------------------------------------------------------------------
//Operadores logicos
func funcLogico() {
	bTrue := true
	bFalse := false

	fmt.Println("-----------|Operadores logicos|---------------")
	fmt.Println("-----------|         &&  E     |---------------")
	fmt.Println(bTrue, "&&", bTrue, "=", bTrue && bTrue)
	fmt.Println(bTrue, "&&", bFalse, "=", bTrue && bFalse)
	fmt.Println(bFalse, "&&", bTrue, "=", bFalse && bTrue)
	fmt.Println(bFalse, "&&", bFalse, "=", bFalse && bFalse)
	fmt.Println("---------------------------------------------")
	fmt.Println("-----------|         || ou      |---------------")
	fmt.Println(bTrue, "||", bTrue, "=", bTrue || bTrue)
	fmt.Println(bTrue, "||", bFalse, "=", bTrue || bFalse)
	fmt.Println(bFalse, "||", bTrue, "=", bFalse || bTrue)
	fmt.Println(bFalse, "||", bFalse, "=", bFalse || bFalse)
	fmt.Println("---------------------------------------------")

	fmt.Println("------- ----|         ! Não       |---------------")
	fmt.Println("!", bTrue, "=", !bTrue)
	fmt.Println("!", bFalse, "=", !bFalse)
	fmt.Println("---------------------------------------------")

}

//---------------------------------------------------------------------------------------------
// Repetições: for
func funcRetFor() {
	fmt.Println("-----------|Repeticoes|---------------")
	fmt.Println("for ini;condicao;passo")

	for i := 0; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

//---------------------------------------------------------------------------------------------
//Repetição: for (substituindo o while)
func funcWhileSQN() {
	fmt.Println("-----------|Repeticoes|---------------")
	fmt.Println("while:nao tem = for condição")

	var soma int = 2
	for soma < 600 {
		fmt.Println(soma)
		soma = soma + soma

	}
	soma = soma + soma
}

//---------------------------------------------------------------------------------------------
// switch
func funcSwitch() {

	var nome string = "Ana"

	fmt.Println("-----------|switch|---------------")

	switch nome {
	case "Ana":
		fmt.Println("Eh a Ana")
	case "Jose":
		fmt.Println("Eh o Jose")
	default:
		fmt.Println("Não conheço")
	}
}

//---------------------------------------------------------------------------------------------
// switch true
func funcSwitchTrue() {

	var nota int = 6

	fmt.Println("-----------|switch|---------------")
	fmt.Print("A nota do aluno foi ", nota, ". Esta ")
	switch {
	case nota >= 0 && nota <= 3:
		fmt.Println("Péssimo")
	case nota > 3 && nota <= 5:
		fmt.Println("Ruim")
	case nota > 5 && nota <= 7:
		fmt.Println("Na média")
	case nota == 8:
		fmt.Println("Bom")
	case nota > 8 && nota <= 10:
		fmt.Println("Excelente")
	default:
		fmt.Println("Nota invaliza. Nota informada: ", nota)
	}
}

//---------------------------------------------------------------------------------------------
// Vetores
func funcVetor() {

	//Declarar sem inicializacao
	var numeros [7]int
	var somaArray int = 0

	fmt.Println("-----------| Array / Vetores |---------------")
	numeros[0] = 1
	numeros[1] = 1
	numeros[2] = 1
	numeros[3] = 1
	numeros[4] = 1
	numeros[5] = 1
	numeros[6] = 1

	fmt.Println(numeros)

	//Declarar com inicializacao
	var numeros2 = [7]int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println(numeros2)

	//usar for para acessar elementos
	for indice := 0; indice < len(numeros2); indice++ {
		fmt.Println(numeros2[indice])
	}

	for nid := 0; nid < len(numeros2); nid++ {
		somaArray += numeros2[nid]
	}

	fmt.Println("A soma do array é", somaArray)

}

//---------------------------------------------------------------------------------------------
// Slice
func funcSlice() {
	var numeros = [7]int{10, 20, 30, 40, 50, 60, 70}

	fmt.Println("-----------| Slice |---------------")
	fmt.Println("Array: ", numeros)
	fmt.Println("Slice[2:5]: ", numeros[2:5])
	fmt.Println("Slice[:5]: ", numeros[:5])
	fmt.Println("Slice[3:]: ", numeros[3:])
}

//---------------------------------------------------------------------------------------------
// Slice
func funcArrayString() {
	var nome = []string{
		"Ana",
		"Jose",
		"Maria",
	}
	fmt.Println("-----------| Array String |---------------")
	fmt.Println(nome)
	fmt.Println("Funcao cap", cap(nome))

	//diferença entre len e cap. (usar cap no o array for um slice)

}
