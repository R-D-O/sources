//-----------------------------------------
//Estudo em golang
//ref: NBK Mundo Tech : https://www.youtube.com/playlist?list=PLUbb2i4BuuzCX8CLeArvx663_0a_hSguW

package main
import (
	"fmt"
	"math"
	//"strings"
)
func main() {
	opcao := -1

	for opcao != 0 {
		fmt.Println("|-------------[Menu]------------|")
		fmt.Println("|1 : Função Anonima             |")
		fmt.Println("|2 : Função Arg Função          |")
		fmt.Println("|3 : Closure                    |")
		fmt.Println("|4 : Método                     |")
		fmt.Println("|5 : Interface                  |")
		fmt.Println("|0 : Sair                       |")
		fmt.Println("|-------------------------------|")
		fmt.Println("Entre com a Opção: ")
		fmt.Scanln(&opcao)

		switch {
		case opcao == 1:
			funcAnonima()
		case opcao == 2:
			funcParamFunc()
		case opcao == 3:
			funcIncItemPed()
		case opcao == 4:
			funcMetodos()
		case opcao == 5:
			funcInterfaces()
		case opcao == 0:
			break
		default:
			fmt.Println("Opção Inválida")
		}
	}
}
//---------------------------------------------------------------------------------
//Salvar a função numa variaveis
func funcAnonima(){

	//Construção comando longo
	var funcSomar func(float64, float64) float64 = func(a,b float64) float64{
		return(a+b)
	}
	fmt.Println("Soma 01...: ",funcSomar(5.5,10.5))

	//comando resumido
	var funcSomar2  = func(a,b float64) float64{
		return(a+b)
	}
	fmt.Println("Soma 02...: ",funcSomar2(5.5,10.5))

}

//---------------------------------------------------------------------------------
//Função com parametros contendo outra função
//Calculado simples

func funcParamFunc(){
	somar := func(a,b float64) float64{
		return a+b
	}

	subtrair := func(a,b float64) float64{
		return a-b
	}

	multiplicar := func(a,b float64) float64{
		return a*b
	}

	dividir := func(a,b float64) float64{
		return a/b
	}

	oper := " "

	fmt.Println("Entre com a operação [ +,-,*,/ ] ")
	fmt.Scanln(&oper)

	switch  {
	case oper == "+":
		fmt.Println("A soma é", computar(somar))
	case oper == "-":
		fmt.Println("A subtração é", computar(subtrair))
	case oper == "*":
		fmt.Println("A multiplicação é", computar(multiplicar))
	case oper == "/":
		fmt.Println("A divisao é", computar(dividir))
	default:
		fmt.Println("Opção invalida!")
	}
}

func computar(fn func(float64, float64) float64) float64{
	valor01 := 0.0
	valor02 := 0.0

	fmt.Print("Entre com valor 01...:")
	fmt.Scan(&valor01)

	fmt.Print("Entre com valor 02...:")
	fmt.Scan(&valor02)

	return fn(valor01,valor02)
}

//---------------------------------------------------------------------------------
//Função closure
//Cadastro
func funcContarItemPed() func() int{
	var totItem int = 0
	return func() int{
		totItem++
		return(totItem)
	}
}

func funcIncItemPed(){
	var addItemPed = funcContarItemPed()
	fmt.Println("Número de Itens no Pedido: ", addItemPed() )
	fmt.Println("Número de Itens no Pedido: ", addItemPed() )
	fmt.Println("Número de Itens no Pedido: ", addItemPed() )
	fmt.Println("Número de Itens no Pedido: ", addItemPed() )
}


type Pessoa struct{
	Nome string
	SobreNome string
}

//metodo
func (p Pessoa) NomeCompleto() string{
	return p.Nome + " " + p.SobreNome
}

// sem metodo
func SemMetodo(p Pessoa) string{
	return p.Nome + " " + p.SobreNome
}


//---------------------------------------------------------------------------------
//Metodos
func funcMetodos(){
	alguem := Pessoa{"Jose ", "da Silva"}

	fmt.Println(alguem.NomeCompleto())
	fmt.Println(SemMetodo(alguem))
}



//---------------------------------------------------------------------------------
//Interfaces
type Geometrica interface {
	area() float64
}

type Quadrado struct {
	lado float64
}

func (q Quadrado) area() float64{
	return q.lado * q.lado
}

type Circulo struct {
	raio float64
}

func (c Circulo) area() float64{
	return math.Pi * c.raio * c.raio
}
func funcInterfaces(){
	var g Geometrica
	g=Quadrado{3}
	fmt.Printf("A area do quadrado é %10.2f é  \n", g.area() )

	g = Circulo{5}
	fmt.Printf("A area do circulo é %10.2f é  \n", g.area() )
}






