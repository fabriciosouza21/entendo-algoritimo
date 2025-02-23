package main

import "fmt"

//criar uma estrutura para representar a fila

type Fila struct {
	elementos []string
}

//enfileirar
//receber uma lista de strings
func (f *Fila) Enfileirar(elementos ...string) {
	// o append é uma função que adiciona um elemento ao final de um array
	f.elementos = append(f.elementos, elementos...)
}

//desenfileirar
func (f *Fila) Desenfileirar() (string){
	if(len(f.elementos) == 0){
		return ""
	}
	// o primeiro elemento da fila
	elemento := f.elementos[0]

	// atualizar a fila
	f.elementos = f.elementos[1:]
	return elemento
}

// pessoa é um vendedor de manga se o nome termina com m
func pessoaEhVendedor(nome string) bool {
	return nome[len(nome)-1] == 'm'
}

var grafo = make(map[string][]string)

func contains(slice []string, valor string) bool {
    for _, v := range slice {
        if v == valor {
            return true
        }
    }
    return false
}

func pesquisa(nome string) bool {
	// primeiro precisamos pesquisar em todos os amigos

	fila := &Fila{elementos: []string{}}

	// adicionar todos os amigos deste nó
	fila.Enfileirar(grafo[nome]...)

	// lista com as pessoas que foram verificadas
	verificados := []string{}

	// enquanto a fila não estiver vazia
	for len(fila.elementos) > 0 {
		// pegar o primeiro elemento da fila
		pessoa := fila.Desenfileirar()

		// somente verificar se a pessoa não foi verificada
		if !contains(verificados, pessoa) {

			//verificar se a pessoa é um vendedor
			if pessoaEhVendedor(pessoa) {
				fmt.Println(pessoa + " é um vendedor de manga!")
				return true
			}else {
				// adicionar os amigos da pessoa na fila
				fila.Enfileirar(grafo[pessoa]...)

				// marcar a pessoa como verificada
				verificados = append(verificados, pessoa)
			}
		}
	}


	return false
}



func main(){

	// criar um mapa que vai presenta o grafo de cada amigo
	grafo["voce"] = []string{"alice", "bob", "claire"}
	grafo["bob"] = []string{"anuj", "peggy"}
	grafo["alice"] = []string{"peggy"}
	grafo["claire"] = []string{"thom", "jonny"}
	grafo["anuj"] = []string{}
	grafo["peggy"] = []string{}
	grafo["thom"] = []string{}
	grafo["jonny"] = []string{}
	pesquisa("voce")
}
