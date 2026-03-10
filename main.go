package main

import "fmt"

// função para interação com usuario
func main() {

	numeros := []int{}

	for {

		fmt.Println("1 - Adicionar número")
		fmt.Println("2 - Listar números")
		fmt.Println("3 - Remover por índice")
		fmt.Println("4 - Estatísticas")
		fmt.Println("5 - Divisão segura")
		fmt.Println("6 - Limpar lista")
		fmt.Println("0 - Sair")

		var opcao int
		fmt.Print("Escolha: ")
		fmt.Scan(&opcao)

		switch opcao {

		case 1:
			numeros = adicionarNumero(numeros)

		case 2:
			listarNumeros(numeros)

		case 3:
			numeros = removerIndice(numeros)

		case 4:
			min, max, media := estatisticas(numeros)
			fmt.Println("Min:", min)
			fmt.Println("Max:", max)
			fmt.Println("Média:", media)

		case 5:
			divisaoSegura()

		case 6:
			numeros = limparLista()

		case 0:
			fmt.Println("Encerrando")
			return

		default:
			fmt.Println("Opção inválida")
		}
	}

}

func adicionarNumero(lista []int) []int {

	var n int

	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	lista = append(lista, n)

	return lista
}
func listarNumeros(lista []int) {

	if len(lista) == 0 {
		fmt.Println("Lista vazia")
		return
	}

	for i, v := range lista {
		fmt.Println(i, ":", v)
	}
}


func removerIndice(lista []int) []int {

	var indice int

	fmt.Print("Digite o índice: ")
	fmt.Scan(&indice)

	if indice < 0 || indice >= len(lista) {
		fmt.Println("Índice inválido")
		return lista
	}

	lista = append(lista[:indice], lista[indice+1:]...)

	return lista
}


func estatisticas(lista []int) (int, int, float64) {

	if len(lista) == 0 {
		return 0, 0, 0
	}

	min := lista[0]
	max := lista[0]
	soma := 0

	for _, v := range lista {

		if v < min {
			min = v
		}

		if v > max {
			max = v
		}

		soma += v
	}

	media := float64(soma) / float64(len(lista))

	return min, max, media
}

func dividir(a int, b int) (int, error) {

	if b == 0 {
		return 0, fmt.Errorf("divisão por zero")
	}

	return a / b, nil
}
func limparLista() []int {
	return []int{}
}

func divisaoSegura() {

	var a int
	var b int

	fmt.Print("Primeiro número: ")
	fmt.Scan(&a)

	fmt.Print("Segundo número: ")
	fmt.Scan(&b)

	resultado, err := dividir(a, b)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Resultado:", resultado)
}