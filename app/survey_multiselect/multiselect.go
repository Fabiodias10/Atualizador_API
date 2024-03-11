package surveymultiselect

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func Surveymultiselect(nomesArquivos []string) []string {
	fmt.Println("")
	// Criar prompt de múltipla seleção usando survey
	var selectedFiles []string
	// var selectedFilesAws []string
	prompt := &survey.MultiSelect{
		Message:  "Escolha as api's:",
		Options:  nomesArquivos,
		PageSize: len(nomesArquivos),
	}

	err := survey.AskOne(prompt, &selectedFiles)
	if err != nil {
		fmt.Println("Erro ao obter a seleção:", err)
		// return
	}

	return selectedFiles
}
