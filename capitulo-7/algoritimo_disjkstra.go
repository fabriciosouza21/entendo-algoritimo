package main

import (
	"fmt"
	"math"
)

// 1 passo é mapear o grafo
// o grafo vai ser um map de map
// o primeiro map vai ser a chave do vertice
// o segundo map vai ser a chave do vertice adjacente e o valor vai ser o peso da aresta

var grafo = make(map[string]map[string]int)


// 2 passo é mapear os custos
var custos = make(map[string]int)

// 3 passo é mapear os pais
var pais = make(map[string]string)

//iniciando o array de processados
var processados = make([]string, 0)

func contains(slice []string, valor string) bool {
    for _, v := range slice {
        if v == valor {
            return true
        }
    }
    return false
}

func ache_no_menor_custo() string{
	custo_menor := math.MaxInt

	nodo_menor := ""

	//interar sobre os custos

	for nodo, custo := range custos {
		if(custo < custo_menor && !contains(processados, nodo)){
			custo_menor = custo
			nodo_menor = nodo
		}
	}

	return nodo_menor

}

func main() {

    // Inicializando o grafo
    grafo["a"] = make(map[string]int) // Inicializa o map interno para "a"
    grafo["a"]["fim"] = 1

    grafo["b"] = make(map[string]int) // Inicializa o map interno para "b"
    grafo["b"]["a"] = 3
    grafo["b"]["fim"] = 5

    grafo["fim"] = make(map[string]int) // Inicializa o map interno para "fim"

	//iniciando os custos
	custos["a"] = 6
	custos["b"] = 2
	custos["fim"] = math.MaxInt

	//iniciando os pais
	pais["a"] = "inicio"
	pais["b"] = "inicio"
	pais["fim"] = ""

	//encontrar o nodo com o menor custo
	nodo := ache_no_menor_custo()

	//enquanto houver nodos para processar
	for nodo != "" {
		custo := custos[nodo]
		//pegar os vizinhos do nodo
		vizinhos := grafo[nodo]

		for vizinho, custo_vizinho :=range vizinhos {
			novo_custo := custo + custo_vizinho

			//verificar se o novo caminho é mais curto
			if(custos[vizinho] > novo_custo){
				//atualizar o custo
				custos[vizinho] = novo_custo

				//atualizar os pais
				pais[vizinho] = nodo

			}

		}
		//marcar o nodo como processado
		processados = append(processados, nodo)

		//busca novamente pelo nodo com o menor custo
		nodo = ache_no_menor_custo()

	}

	fmt.Println("Custos: ", custos)
	fmt.Println("Pais: ", pais)
	fmt.Println("Processados: ", processados)

}
