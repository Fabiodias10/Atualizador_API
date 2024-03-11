package main

import (
	confirma "atualiza_api/app/confirma_baixar"
	copia "atualiza_api/app/copia_api"
	download "atualiza_api/app/downloadAndExtract"
	"atualiza_api/app/downloadAws"
	"atualiza_api/app/iis"
	"atualiza_api/app/inicio"
	"atualiza_api/app/listarAws"
	listar "atualiza_api/app/listar_repositorio_perto"
	separa "atualiza_api/app/separa_slice"
	multiselect "atualiza_api/app/survey_multiselect"

	"fmt"
	"os"
	"os/exec"
	"runtime"

	// "teste/app/listarAws"
	"time"
)

func main() {
	url := "http://prkng-hml-a.dsr.net.br:32203/repository/Servidores/Apis/XE/"
	// url := "http://prkng-hml-a.dsr.net.br:32203/repository/Apps/CaixaWindows"

	resposta_admin, err := inicio.IsAdmin()
	if err != nil {
		fmt.Println("Erro ao verificar privilégios de administrador:", err)
		// os.Exit(1)
	}
	if resposta_admin {
		// fmt.Println("O programa está sendo executado em modo administrador.")

		// fmt.Println("É administrado do sistema? ", resposta_admin)
		//lista repositorio perto e e tras nome dos arquivos []string
		nomesArquivos := listar.Listarrepositorioperto(url)
		_, resposta, err := listarAws.Listar() //lista aws
		if err != nil {
			fmt.Println("Erro ao listar arquivos:", err)
		}
		juntar_slice := append(nomesArquivos, resposta...)

		//retorno dos arquivos selecionados []string
		selectedFiles := multiselect.Surveymultiselect(juntar_slice)

		lista1, lista2 := separa.Separa(selectedFiles)

		for len(selectedFiles) == 0 {

			// Limpar o console
			clearConsole()
			selectedFiles = multiselect.Surveymultiselect(juntar_slice)
			// fmt.Println("Nenhum arquivo selecionado. Por favor, tente novamente.")

		}

		resp_confirma := confirma.Confirma()
		if resp_confirma {
			inicio := time.Now()
			// download.Efetiva_download(selectedFiles, url)
			download.Efetiva_download(lista1, url)

			downloadAws.Download(lista2)
			fmt.Println("Downloads Concluidos")
			fmt.Printf("Descompactando...")
			download.Descompactando()
			fmt.Println(" OK")

			inicioIIS := time.Now()
			fmt.Printf("Tentando parar IIS...")
			iis.IIS_Stop()

			copia.CopiarApi()
			iis.IIS_Start()

			fmt.Println("")
			fmt.Printf("Tempo total que o IIS ficou fora do ar:  %s \n", time.Since(inicioIIS))
			fmt.Printf("Tempo total da atualização das Apis: %s \n", time.Since(inicio))
			// Aguardar qualquer tecla ser pressionada antes de sair
			fmt.Println("")
			fmt.Println("Pressione qualquer tecla para sair...")
			fmt.Scanln()
		} else {

			// Aguardar qualquer tecla ser pressionada antes de sair
			fmt.Println("")
			fmt.Println("Pressione qualquer tecla para sair...")
			fmt.Scanln()
		}
		//passa o slice de string para este metodo fazer o download e unzip dos arquivos

	} else {
		fmt.Println("Execute o programa como administrador para garantir privilégios necessários.")
		// Aguardar qualquer tecla ser pressionada antes de sair
		fmt.Println("")
		fmt.Println("Pressione qualquer tecla para sair...")
		fmt.Scanln()
	}
}

func clearConsole() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
