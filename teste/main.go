package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Comando para parar o serviço do IIS
	// cmd := exec.Command("net", "stop", "w3svc")
	cmd := exec.Command("iisreset", "/stop")

	// Executa o comando
	err := cmd.Run()

	// Verifica se ocorreu algum erro durante a execução do comando
	if err != nil {
		fmt.Println("Erro ao parar o IIS:", err)
		return
	}

	fmt.Println("IIS parado com sucesso.")

}
