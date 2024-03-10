package confirma

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func Confirma() bool {
	var resposta bool

	// Pergunta ao usuário se deseja prosseguir
	prompt := &survey.Confirm{
		Message: "Deseja prosseguir com download e atualização automatica das Api's?",
	}
	survey.AskOne(prompt, &resposta)

	// Verifica a resposta
	if resposta {
		fmt.Println("")
		// Coloque aqui o código que você deseja executar se a resposta for "sim"
	} else {
		fmt.Println("Operação cancelada.")
		// Coloque aqui o código que você deseja executar se a resposta for "não"
	}
	return resposta
}
