package main

import (
	"sync"
	"time"
)

func loadBalance(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Minute)
}

func alocarMemoria(tamanho int) [][]byte {

  var memoria [][]byte

	for {

    bloco := make([]byte, tamanho)

    for i := range bloco {
			bloco[i] = byte(i % 256)
		}

		memoria = append(memoria, bloco)
		time.Sleep(500 * time.Millisecond)
	}

	return memoria
}

func main() {
	
  var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		loadBalance(i, &wg)
	}
	tamanho := 10 * 1024 * 1024
	alocarMemoria(tamanho)

}
