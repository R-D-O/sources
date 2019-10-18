//--------------------------------------------
//Estudo golang
//18/10/2019
//Ref: NBK Mundo Tech (canal youtube)

package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
)

func main(){

	opcao := -1

	for opcao != 0 {
		fmt.Println("|-------------[Menu]------------|")
		fmt.Println("|1 : Diferença entre Cap e Len  |")
		fmt.Println("|2 : Make                       |")
		fmt.Println("|3 : Adiciona elemento no vetor |")
		fmt.Println("|4 : Matriz                     |")
		fmt.Println("|5 : Ponteiro                   |")
		fmt.Println("|6 : Struct                     |")
		fmt.Println("|7 : Range                      |")
		fmt.Println("|8 : Maps                       |")
		fmt.Println("|0 : Sair                       |")
		fmt.Println("|-------------------------------|")
		fmt.Println("Entre com a Opção: ")
		fmt.Scanln(&opcao)

		switch {
		case opcao == 1:
			funcCapArray()
		case opcao == 2:
			funcMake()
		case opcao == 3:
			funcAppend()
		case opcao == 4:
			funcMatriz()
		case opcao == 5:
			funcPonteiro()
		case opcao == 6:
			funcEstrutura()
		case opcao == 7:
			funcRange()
		case opcao == 8:
			funcMaps()
		case opcao == 0:
			break
		default:
			fmt.Println("Opção Inválida")
		}
	}
}



//---------------------------------------------------------------------------------------------
// Pausar
func funcPausa() {
	fmt.Print("Pressione 'ENTER' para continuar")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}


//---------------------------------------------------------------------------------------------
//Diferença entre o Cap e Len
func funcCapArray(){
	var animes [3]string;
	animes = [3] string{"Dragon Ball","Sailon Moon","Yuyu Hakusho"}

	var doisPrimeiros = animes[:2]
	var doisUltimos = animes[1:]

	fmt.Println("-----------Len vs Cap----------")

	fmt.Println("Lista Completa")
	fmt.Println(animes)

	fmt.Println("Dois primeiros [:2]")
	fmt.Println(doisPrimeiros)

	fmt.Println("Dois ultimos [1:]")
	fmt.Println(doisUltimos)


	fmt.Printf("len = %d   cap = %d\n", len(doisPrimeiros), cap(doisPrimeiros))
	fmt.Printf("len = %d   cap = %d\n", len(doisUltimos), cap(doisUltimos))

	funcPausa()
}

//---------------------------------------------------------------------------------------------
//Função Make
func funcMake()  {
	numeros := make([]int, 5)

	fmt.Println("-----------make----------")
	fmt.Println(numeros)

	funcPausa()
}


//---------------------------------------------------------------------------------------------
//Adicionar elementos no vetor
func funcAppend()  {

	var nomes []string;
	var nId int = 1

	for nId <= 5 {
		fmt.Println("Entre com um Nome: ", nId, "de 5")
		var input string
		fmt.Scanln(&input)

		if len(input) > 0{
			nomes = append(nomes,input)
			fmt.Println("Nomes digitados: ", nomes)
			nId++
		}else{
			fmt.Println("Não informou o nome. Tente novamente")
		}

	}



	funcPausa()
}
//---------------------------------------------------------------------------------------------
//Array Matriz
func funcMatriz(){

	agenda := [][]string{}
	var nId int = 1
	var nome string = ""
	var telefone string = ""

	for nId <= 3{
		nome = ""
		telefone = ""

		fmt.Println("Entre com um Nome: ")
		fmt.Scanln(&nome)

		fmt.Println("Entre com telefone: ")
		fmt.Scanln(&telefone)

		if len(nome) > 0 && len(telefone) > 0{
			agenda = append(agenda, []string{nome,telefone} )
			nId++

		}
	}

	fmt.Println("-------------------------[ Agenda ]-------------------------")
	nId = 0
	for nId < len(agenda){
		fmt.Printf("|Nome: %15s |Telefone: %12s|\n", agenda[nId][0], agenda[nId][1])
		//fmt.Printf("Nome / Telefone:  %20s\n",strings.Join(agenda[nId]," "))
		nId++
	}
	fmt.Println("------------------------------------------------------------")
	funcPausa()
}
//---------------------------------------------------------------------------------------------
//Ponteiro: endereço de memoria das variaveis
//

func funcPonteiro(){
	//Declarar as variavie e atibuir valor
	Variavel := 55

	//Declaração de ponteiro
	var ponteir *int

	//Mostrar valor da variavel
	fmt.Println("Variavel =",Variavel)

	//Atribuir o ponteiro da varivel
	PonteiroVariaveis := &Variavel
	fmt.Println("Endereco de memoria da Variavel = ", PonteiroVariaveis)

	//Alterar a varivel pelo ponteiro
	*PonteiroVariaveis = 25

	//Imprimir
	fmt.Println("Variavel = ", Variavel)
	fmt.Println("Ponteiro = ", *PonteiroVariaveis)
	fmt.Println("Ponteiro sem atribuição = ", ponteir) //sem variavel associal eh nil
	funcPausa()
}
//---------------------------------------------------------------------------------------------
//Estrutura
//
func funcEstrutura(){

	//Declara uma struc
	type usuario struct {
		nome string
		email string
	}

	//Imprimir
	fmt.Println(usuario{"Rodrigo","rodrigo@mail.com"})

	//Criar
	stUser := usuario{"Jose","jose.silva@mail.com"}

	//Imprimir por elemento
	fmt.Println("Nome...: ", stUser.nome )
	fmt.Println("E-mail.: ", stUser.email)

	//Alterar um elemento
	stUser.nome = "Joao"

	//Imprimir por elemento
	fmt.Println("Nome [alterado]...: ", stUser.nome )
	fmt.Println("E-mail............: ", stUser.email)

	//identificando os campos das struct, não precisa seguir a ordem da declaração
	stUser2 := usuario{email:"email@email.com",nome:"carla"}

	//Impressao
	fmt.Println(stUser2)

	//Ponteiro
	var stUser3 usuario = usuario{"Bianca","bianca@email.com"}
	var ponteiroStruc *usuario = &stUser3

	//Impressao
	fmt.Println("Impressão da struct:", stUser3)

	//impressão do ponteiro
	fmt.Println(  "Impressão do ponteiro da struct: ",*ponteiroStruc )

	//impressão do campo do ponteiro
	fmt.Println(  "Impressão do ponteiro da struct: ", (*ponteiroStruc).nome )

	//impressao do enderedo de memoria
	fmt.Println(  "Endereço memoria: ", ponteiroStruc)




	funcPausa()
}


//---------------------------------------------------------------------------------------------
//Range
//
func funcRange(){

	numeros := []int{100,200,300,400}
	//relembrando com for

	fmt.Println("|------------[Com for tradicional]------------------")
	for i:=0; i < len(numeros); i++{
		fmt.Printf("|Indice %d |valor %d\n", i,numeros[i])
	}

	//range com indice e valor
	fmt.Println("|------------[Range: indice e valor]-----------------")
	for indice, v := range numeros{
		fmt.Printf("|Indice %d |valor %d\n", indice,v)
	}


	//com range, omitindo o indice
	fmt.Println("|------------[Range: valor]------------------------")
	for _, v := range numeros{
		fmt.Printf("|valor %d\n", v)
	}

	//com range, omitindo o valor
	fmt.Println("|------------[Range: indice]------------------------")
	for indice, _ := range numeros{
		fmt.Printf("|Indice %d |valor %d\n", indice,numeros[indice])
	}

}

//---------------------------------------------------------------------------------------------
//Maps
//
func funcMaps(){
	//Declarar
	var notas map[string]int

	//Montar
	notas = make(map[string]int)

	//Atribuir valores
	notas["Ana"] = 9
	notas["Maria"] = 10

	//Imprimir
	fmt.Println(notas)

	//verificar se a chave exist

	valor, existe := notas["Joao"]

	fmt.Println(valor) //valor indexado
	fmt.Println(existe) //true: existe, false:não existe

	if existe{
		fmt.Println(notas["joao"])
	}

	//maps literais
	var notas2 = map[string]int{
		"Ana":9,
		"Maria":10,
	}

	fmt.Println("Maps Literal: ", notas2)

	//Maps com struct
	type Perfil struct{
		Altura float64
		Ativo bool
		Profissao string
	}

	var perfis map[string]Perfil = make(map[string]Perfil)

	perfis["Joao"] = Perfil{
		1.74, true , "Medico",
	}

	perfis["Maria"] = Perfil{
		1.63, false , "Engenheira",
	}

	fmt.Println(perfis)
}
