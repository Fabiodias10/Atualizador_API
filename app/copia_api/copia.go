package copia

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/otiai10/copy"
)

var listaNomesAceitos = map[string]struct{}{
	"ApiAccess":       {},
	"ApiMobile":       {},
	"ApiDataParking":  {},
	"ApiPartner":      {},
	"ApiPayment":      {},
	"ApiRegistration": {},
	// Adicione mais nomes conforme necessário
}

// copia todas apis e renomea a api de pagamento diferentona
func CopiarApi() {
	dataAtual := time.Now().Format("02-01-2006")
	pastaOrigem := filepath.Join("C:/temp", "API-"+dataAtual)
	// pastaOrigem := "c:/temp/API-09-03-2024"
	// pastaDestino := "c:/temp/destino"
	pastaDestino := "C:/inetpub/wwwroot/Parking"

	diretorio := pastaOrigem
	nomeAntigo := "ApiPagamentoXE"
	novoNome := "ApiPayment"

	percorrerPastaERenomear(diretorio, nomeAntigo, novoNome)
	// if erra != nil {
	// 	fmt.Println("Erro percorrer pasta e renomear:", erra)
	// }

	err := filepath.Walk(pastaOrigem, func(caminhoPasta string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Ignora os arquivos, processamos apenas pastas
		if !info.IsDir() {
			return nil
		}

		// Obtém o nome da pasta
		nomePasta := info.Name()

		// Verifica se o nome da pasta está na lista de nomes aceitos
		if _, existe := listaNomesAceitos[nomePasta]; existe {
			// Copia a pasta inteira para o destino
			destinoPasta := filepath.Join(pastaDestino, nomePasta)

			// Remove a pasta de destino se ela já existir
			if err := os.RemoveAll(destinoPasta); err != nil {
				fmt.Printf("Erro ao remover pasta de destino %s: %s\n", pastaDestino, err)
				return nil
			}

			err := copy.Copy(caminhoPasta, destinoPasta)
			if err != nil {
				fmt.Printf("Erro ao copiar pasta %s: %s\n", caminhoPasta, err)
			} else {
				// fmt.Printf("Pasta %s copiada para %s com sucesso.\n", caminhoPasta, destinoPasta)
				fmt.Printf("%s copiada com sucesso.\n", filepath.Base(destinoPasta))
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Erro ao percorrer as pastas:", err)
	}

}

func percorrerPastaERenomear(diretorio, nomeAntigo, novoNome string) error {

	err := filepath.Walk(diretorio, func(caminho string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Verifica se é um diretório e se o nome corresponde ao nomeAntigo
		if info.IsDir() && info.Name() == nomeAntigo {
			// Gera o novo caminho com o novo nome
			novoCaminho := filepath.Join(filepath.Dir(caminho), novoNome)
			// Renomeia o diretório
			err := os.Rename(caminho, novoCaminho)
			if err != nil {
				return err
			}
			fmt.Printf("Diretório renomeado de %s para %s\n", caminho, novoCaminho)
		}
		return nil
	})

	return err
}
