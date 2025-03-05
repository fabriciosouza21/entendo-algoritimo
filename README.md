# Resumo

esse é um resumo do livro Entendendo Algoritmos - Um Guia Ilustrado Para Programadores e Outros Curiosos - Autor (Aditya Y. Bhargava).

A finalidade desse repositorio e fazer um resumo de cada capitulo do livro, as implementações vão ser feitas em Go


# Introdução

# capítulo 1

- A pesquisa binária é muito mais rápida do que a pesquisa simples.
- O(log n) é mais rápido do que O(n), e O(log n) fica ainda mais rápido
conforme os elementos da lista aumentam.
- A rapidez de um algoritmo não é medida em segundos.
- O tempo de execução de um algoritmo é medido por meio de seu crescimento.
- O tempo de execução dos algoritmos é expresso na notação Big O.

# capítulo 2

Ordenação por seleção

- arrays e lista encadeadas
- algoritmo de Ordenação

# capítulo 3

> Leigh Caldwell, do Stack Overflow, diz "Os loops podem melhorar o desempenho do seu programa . A recursão melhorar o desenpenho do seu progrador. Escolha o que for mais importante para a sua situação."

- Toda função recursiva tem dois casos: o caso base e o caso recursivo.
- Todas as chamadas de função vão para a pilha de chamadas.
- A pilha de chamadas pode ficar muito grande e ocupar muita memória.
- As funções recursivas podem ser substituídas por iterativas.
- A recursão é muito utilizada em algoritmos de divisão e conquista.
- Toda função recursiva pode ser transformada em iterativa.


# capítulo 4

Dividir para conquistar e quicksort

- A estratégia DC funciona por meio da divisão do problema em problemas menores. Se vocẽ estiver utilizando DC em uma lista, o caso-base provavelmente será uma lista vazia ou com apenas um elemento
- se você esive implementado o quicksort, escolha um elemento aleatório como pivô para garantir que o tempo de execução seja O(n log n) no caso médio.
- A constante, na notação Big O, pode ser relevante em alguns casos. Esta é
a razão pela qual o quicksort é mais rápido do que o merge sort.

# capítulo 5

Tabelas hash


- Você pode fazer uma tabela hash ao combinar uma função hash com um
array.
- Colisões são problemas. É necessário haver uma função hash que
minimize colisões.
- As tabelas hash são extremamente rápidas para pesquisar, inserir e
remover itens.
- Tabelas hash são boas para modelar relações entre dois itens.
- Se o seu fator de carga for maior que 0.7, será necessário redimensionar a
sua tabela hash.
- As tabelas hash são utilizadas como cache de dados (como em um
servidor da web, por exemplo).
- Tabelas hash são ótimas para localizar duplicatas.

# capítulo 6

A pesquisa em largura permite enconra o menor caminho entre dois objetos.


- A pesquisa em Largura lhe diz se há um caminho A para B
- Se esse caminho existir, a pesquisa em largura lhe dará o caminho mínimo
- se você tem um problema do tipo "encontre o meno X", tente modela o seu problema utilizando grafos e use a pesquisa em largura pra resolvê-lo
- um dígrafo contém setas e as relações seguem a direção das setas
- Grafos não direcionados não contêm setas, e a relação acontece nos dois sentidos
- Filas são FIFO (primeiro a entrar, primeiro a sair)
- Pilhas são LIFO (último a entrar, primeiro a sair).
- Você precisa vericar as pessoas na ordem em que elas foram adicionadas à lista de pesquisa. Portanto a lista de pesquisa deve ser uma fila; caso contrário, você não obterá o caminho mínimo.
- Cada vez que você precisar vericar alguém, procure não vericá-lo novamente. Caso contrário, poderá acabar em um loop infinito.

# capítulo 7

Utilizamos dijkstra em grafos aonde temos um grafo direcionado e ponderado, onde as arestas tem um peso.Dijkstra é um algoritmo de que encontra o caminho mais rapido, enquanto o algoritmo de pesquisa em largura encontra o caminho mais curto(com a menor quantidade de arestas).

- A pesquisa em largura é usada para calcular o caminho mínimo para um
grafo não ponderado.
- O algoritmo de Dijkstra é usado para calcular o caminho mínimo para
um grafo ponderado.
- O algoritmo de Dijkstra funciona quando todos os pesos são positivos.
- Se o seu grafo tiver pesos negativos, use o algoritmo de Bellman-Ford.


# capítulo 8

Um algoritmo guloso é simples: a cada etapa, deve-se escolher o movimento ideal.

- Algoritmos gulosos otimizam localmente na esperança de acabar em uma
otimização global.
- Problemas NP-completo não têm uma solução rápida.
- Se você estiver tentando resolver um problema NP-completo, o melhor a
fazer é usar um algoritmo de aproximação.
- Algoritmos gulosos são fáceis de escrever e têm tempo de execução baixo,
portanto eles são bons algoritmos de aproximação.


# capítulo 9

Programação dinâmica

A programção dinamica inicia com problemas menores e os resolve até chagar ao prblema geral.

o problema parte do principio que existe um tabela de itens, as linhas são os itens e as colunas são o peso da mochila.

- A programação dinâmica é útil quando você está tentando otimizar algo
em relação a um limite.
- Você pode utilizar a programação dinâmica quando o problema puder ser
dividido em subproblemas discretos.
- Todas as soluções em programação dinâmica envolvem uma tabela.
- Os valores nas células são, geralmente, o que você está tentando otimizar.
- Cada célula é um subproblema, então pense sobre como é possível dividir
este subproblema em outros subproblemas.
- Não existe uma fórmula única para calcular uma solução em
programação dinâmica.

# capítulo 10

KNN - O algoritmo usa os vizinhos mais proximo para classificar um novo ponto.


-  O algoritmo dos k-vizinhos mais próximos é utilizado na classicação e
também na regressão. Ele envolve observar os K-vizinhos mais próximos.
- Classicação = classicar em grupos.
- Regressão = adivinhar uma resposta (como um número).
- Extrair características signica converter um item (como uma fruta ou
um usuário) em uma lista de números que podem ser comparados.
- Escolher boas características é uma parte importante para que um
algoritmo dos k-vizinhos mais próximos opere corretamente.





