package separa

import "fmt"

func Separa(nomesArquivos []string) ([]string, []string) {
	// Definir a primeira lista
	// lista1 := []string{"Item1", "Item2", "Item3", "Item1", "OutroItem", "Item1"}
	lista1 := nomesArquivos
	// Criar a segunda lista
	items_para_localizar_separar := []string{"ApiMonitoring.zip", "GerenciadorLocal.zip", "ApiLPR_Hikvision.zip", "ApiLPR_Intelbras.zip", "ApiLPR_Quercus.zip"}
	lista2 := []string{}

	// Iterar sobre a primeira lista
	for _, itemParaLocalizar := range items_para_localizar_separar {
		// Verificar se o item está na lista de itens para localizar e separar
		for i, item := range lista1 {
			if item == itemParaLocalizar {
				// Adicionar o item à segunda lista
				lista2 = append(lista2, item)
				// Remover o item da primeira lista
				lista1 = append(lista1[:i], lista1[i+1:]...)
			}
		}
	}

	// Imprimir as listas atualizadas
	fmt.Println("Lista 1:", lista1)
	fmt.Println("Lista 2:", lista2)
	return lista1, lista2
}

// Função auxiliar para remover um item de uma slice
func removeItem(slice []string, item string) []string {
	for i, v := range slice {
		if v == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
