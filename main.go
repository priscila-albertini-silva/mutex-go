package main

import (
	"fmt"
	"sync"
	"time"
)

// Dados compartilhados pelas goroutines
var (
	data = make(map[string]int)
	mtx  sync.Mutex
)

// Função que incrementa um valor no mapa compartilhado
func increment(key string) {
	// Bloqueia o acesso para garantir exclusividade
	mtx.Lock()
	defer mtx.Unlock()

	// Incrementa o valor da chave no mapa
	data[key]++
}

// Função que imprime o valor associado a uma chave no mapa
func readData(key string) {
	// Bloqueia o acesso para garantir exclusividade
	mtx.Lock()
	defer mtx.Unlock()

	// Lê o valor da chave no mapa e imprime
	fmt.Printf("Chave: %s, Valor: %d\n", key, data[key])
}

func main() {
	// Número de goroutines concorrentes
	numGoroutines := 5

	// WaitGroup é usado para esperar todas as goroutines terminarem
	var wg sync.WaitGroup

	// Inicializa o mapa com valores iniciais
	data["chave1"] = 10
	data["chave2"] = 20

	// Inicia as goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// Incrementa os valores das chaves em um loop
			for j := 0; j < 5; j++ {
				increment("chave1")
				increment("chave2")
				time.Sleep(time.Millisecond * 100)
			}

			readData("chave1")
			readData("chave2")
		}()
	}

	// Aguarda todas as goroutines terminarem
	wg.Wait()

	// Após todas as goroutines terminarem, imprime o estado final do mapa
	fmt.Println("Estado final do mapa:")
	for key, value := range data {
		fmt.Printf("Chave: %s, Valor: %d\n", key, value)
	}
}
