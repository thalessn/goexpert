package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.WriteString("Hello World")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d\n", tamanho)
	f.Close()

	//leitura
	// arquivo, err := os.ReadFile("arquivo.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(arquivo))

	//Leitura por Stream
	//Abre o arquivo para ter uma refencia
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	// Para ler em chuncks é necessario criar um reader
	reader := bufio.NewReader(arquivo2)
	// O reader precisar de um slice de bytes definido o tamanho da "linha/pedaço" do que será lido.
	buffer := make([]byte, 3)

	//Cria-se um for infinito para percorrer todo o arquivo e quando chegar ao final sairá do loop
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
}
