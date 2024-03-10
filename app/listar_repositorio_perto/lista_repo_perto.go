package listarrepositorioperto

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Listarrepositorioperto(url string) []string {
	fmt.Println("")
	// URL da página que contém os links para os arquivos

	// Obter a lista de nomes de arquivo da página
	nomesArquivos, err := getNomesArquivos(url)
	if err != nil {
		fmt.Println("Erro ao obter os nomes dos arquivos:", err)
		// return
	}
	return nomesArquivos
}

func getNomesArquivos(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Erro na solicitação. Código de status: %d", resp.StatusCode)
	}

	// Usar goquery para analisar o HTML da página
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var nomesArquivos []string

	// Encontrar e extrair os nomes dos arquivos do HTML
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists && strings.HasPrefix(link, "/") {
			// Obter o segmento final do URL como nome do arquivo
			segments := strings.Split(link, "/")
			nomeArquivo := segments[len(segments)-1]
			// Verificar se o nome do arquivo não está vazio antes de adicionar ao slice
			if nomeArquivo != "" {
				nomesArquivos = append(nomesArquivos, nomeArquivo)
			}
		}
	})

	return nomesArquivos, nil
}
