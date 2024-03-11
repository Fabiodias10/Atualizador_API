package main

import (
	"fmt"
	"teste/download"
	"teste/listar"
	"teste/survey"
)

func main() {
	fmt.Println("AWS S3")

	_, resposta, err := listar.Listar()
	if err != nil {
		fmt.Println("Erro ao listar arquivos:", err)
	}
	resp := survey.Surveymultiselect(resposta)
	// fmt.Println(key)
	download.Download(resp)

	// Aguardar a tecla ser pressionada antes de sair
	fmt.Println("Pressione Enter para sair...")
	fmt.Scanln()
}
