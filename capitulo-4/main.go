package main

import "fmt"

func quicksort(list []int) []int {

	// o caso base
	// arrays com somente um elemento
	if(len(list) <2){
		return list
	}else {
		// escolher um elemento para ser o pivô
		// neste caso, o primeiro elemento
		pivo := list[0]

		// subarray de elementos menores que o pivo
		// sem ser importa com a ordenação

		menores := []int{}
		maiores := []int{}
		// ignorar o primeiro elemento
		// pois ele é o pivo
		for _, item := range list[1:] {
			if(item <= pivo){
				menores = append(menores, item)
			}else {
				maiores = append(maiores, item)
			}
		}

		// chamada recursiva
		// juntar os elementos ordenados
		return  append(append(quicksort(menores), pivo), quicksort(maiores)...)
	}


}

//receber uma lista de interios e retornar a soma de todos os elementos
func somaRecursiva(lista []int) int {
	//caso base
	if(len(lista) == 0){
		return 0
	}
	//chamada recursiva
	return lista[0] + somaRecursiva(lista[1:])
}

func main() {
	lista := []int{2,4, 6}
	fmt.Println(somaRecursiva(lista))

	fmt.Println("somaRecursiva: ", somaRecursiva(lista))

	lista2 := []int{10, 5, 2, 3}
	fmt.Println("quicksort: ", quicksort(lista2))
}
