package main

import "fmt"

func pesquisaBinaria(lista []int, item int) int {

	baixo := 0
	alto := len(lista) - 1

	for baixo <= alto {
		// Calcula o meio
		meio := (baixo + alto) / 2

		// Chute o elemento do meio
		chute := lista[meio]

		// Encontrou o item
		if chute == item {
			return meio
		}

		// O chute foi muito alto
		if chute < item {
			baixo = meio + 1
		}

		// O chute foi muito baixo
		if chute > item {
			alto = meio - 1
		}
	}

	// O item não existe
	return -1

}

func main() {

	lista := []int {1, 3, 5, 7, 9}

	fmt.Println(pesquisaBinaria(lista, 3)) // 1
	fmt.Println(pesquisaBinaria(lista, -5)) // -1

}
