package iis

import (
	"fmt"
	"os/exec"
)

func IIS_Stop() {
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

	fmt.Println("IIS parado.")

}
func IIS_Start() {
	// Comando para parar o serviço do IIS
	// cmd := exec.Command("net", "stop", "w3svc")
	cmd := exec.Command("iisreset", "/start")

	// Executa o comando
	err := cmd.Run()

	// Verifica se ocorreu algum erro durante a execução do comando
	if err != nil {
		fmt.Println("Erro ao parar o IIS:", err)
		return
	}

	fmt.Println("IIS Iniciado.")

}
