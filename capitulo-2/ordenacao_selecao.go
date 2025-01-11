package main

import "fmt"

func buscaMenor(arr [] int) int {
	//inicializar as variaveis

	menor := arr[0]
	menor_indice := 0

	for i := 1; i< len(arr); i++ {
		if(arr[i] < menor){
			menor = arr[i]
			menor_indice = i
		}
	}
	return menor_indice
}

func delete_at_index(slice []int, index int) []int {

    // Append function used to append elements to a slice
    // first parameter as the slice to which the elements
    // are to be added/appended second parameter is the
    // element(s) to be appended into the slice
    // return value as a slice
	/*
	O operador ... é essencial aqui porque transforma a slice slice[index+1:] em uma sequência de elementos, permitindo que eles sejam adicionados corretamente ao final de slice[:index].
	*/
    return append(slice[:index], slice[index+1:]...)
}



func ordenacaoSelecao(arr []int) []int {
	novoArr := []int{}
	//navegar no array
	// buscar o menor elemento
	// adicionar o menor elemento no novo array
	// remover o menor elemento do array original

	/*
	se utilizar o for i := 0; i < len(arr); i++ {, o tamanho do array vai mudar a cada iteração, pois estamos removendo elementos do array original.
	*/
	size := len(arr)
	for i := 0; i < size; i++ {
		menor := buscaMenor(arr)
		novoArr = append(novoArr, arr[menor])
		arr = delete_at_index(arr, menor)
	}

	return novoArr
}

func main() {
	arr := []int{5, 3, 6, 2, 10}
	// [2, 3, 5, 6, 10]
	fmt.Println(ordenacaoSelecao(arr))

}
