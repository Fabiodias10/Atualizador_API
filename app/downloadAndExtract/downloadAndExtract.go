package downloadAndExtract

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/nwaples/rardecode"
)

// sadf
func Download(fileURL, fileName string) error {

	// Obter a data e hora atual no formato brasileiro
	dataAtual := time.Now().Format("02-01-2006")
	diretorioDestino := filepath.Join("C:/temp", "API-"+dataAtual)

	// Criar o diretório se não existir
	if err := os.MkdirAll(diretorioDestino, os.ModePerm); err != nil {
		return err
	}
	caminhoArquivo := filepath.Join(diretorioDestino, fileName)
	// Baixar o arquivo
	resp, err := http.Get(fileURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Criar arquivo local para salvar o conteúdo
	out, err := os.Create(caminhoArquivo)
	if err != nil {
		return err
	}
	defer out.Close()

	// Copiar o conteúdo do arquivo baixado para o arquivo local
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func Efetiva_download(selectedFiles []string, url string) {

	// Obter a data e hora atual no formato brasileiro
	dataAtual := time.Now().Format("02-01-2006")
	diretorioDestino := filepath.Join("C:/temp", "API-"+dataAtual)
	// Remove a pasta de destino se ela já existir
	if err := os.RemoveAll(diretorioDestino); err != nil {
		fmt.Printf("Erro ao remover pasta de destino %s: %s\n", diretorioDestino, err)
		return
	}
	// Cria um WaitGroup para esperar que todas as goroutines terminem
	var wg sync.WaitGroup
	// Canal para comunicar erros entre goroutines
	errCh := make(chan error, len(selectedFiles))
	// Baixar e descompactar cada arquivo selecionado
	for _, arquivo := range selectedFiles {
		wg.Add(1)
		go func(arquivo string) {
			defer wg.Done()
			fmt.Printf("Baixando: %s\n", arquivo)
			err := Download(url+arquivo, arquivo)
			if err != nil {
				fmt.Printf("Erro ao baixar e descompactar %s: %v\n", arquivo, err)
				errCh <- err
			}
		}(arquivo)
	}

	// Função anônima para fechar o canal de erro quando todas as goroutines terminarem
	go func() {
		wg.Wait()
		close(errCh)
	}()

	// Aguarda todas as goroutines terminarem e verifica erros
	for err := range errCh {
		if err != nil {
			fmt.Printf("Erro durante o download: %v\n", err)
		}
	}

	fmt.Println("Downloads concluídos.")
	fmt.Printf("Descompactando...")
	Descompactando()
	fmt.Println(" OK")

}

func DescompactarZip(arquivoZip, destino string) error {
	reader, err := zip.OpenReader(arquivoZip)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, arquivo := range reader.File {
		caminhoArquivo := filepath.Join(destino, arquivo.Name)

		if arquivo.FileInfo().IsDir() {
			os.MkdirAll(caminhoArquivo, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(caminhoArquivo), os.ModePerm); err != nil {
			return err
		}

		arquivoExtraido, err := os.Create(caminhoArquivo)
		if err != nil {
			return err
		}
		defer arquivoExtraido.Close()

		entrada, err := arquivo.Open()
		if err != nil {
			return err
		}
		defer entrada.Close()

		if _, err := io.Copy(arquivoExtraido, entrada); err != nil {
			return err
		}
	}

	return nil
}

func DescompactarRar(arquivoRar, destino string) error {
	arquivo, err := os.Open(arquivoRar)
	if err != nil {
		return err
	}
	defer arquivo.Close()

	// rarReader, err := rardecode.NewReaderSize(arquivo, nil, 0)
	rarReader, err := rardecode.NewReader(arquivo, "")
	if err != nil {
		return err
	}

	for {
		header, err := rarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		caminhoArquivo := filepath.Join(destino, header.Name)

		if header.IsDir {
			os.MkdirAll(caminhoArquivo, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(caminhoArquivo), os.ModePerm); err != nil {
			return err
		}

		arquivoExtraido, err := os.Create(caminhoArquivo)
		if err != nil {
			return err
		}
		defer arquivoExtraido.Close()

		if _, err := io.Copy(arquivoExtraido, rarReader); err != nil {
			return err
		}
	}

	return nil
}
func Descompactando() {
	// Obter a data e hora atual no formato brasileiro
	dataAtual := time.Now().Format("02-01-2006")
	pastaOrigem := filepath.Join("C:/temp", "API-"+dataAtual)

	// Canal para receber mensagens de conclusão
	doneCh := make(chan struct{})

	// Cria um WaitGroup para esperar que todas as goroutines terminem
	var wg sync.WaitGroup

	// Canal para comunicar erros entre goroutines
	errCh := make(chan error, 1)

	// Use um semáforo para controlar o número máximo de goroutines em paralelo
	semaphore := make(chan struct{}, 5) // altere o número conforme necessário

	err := filepath.Walk(pastaOrigem, func(caminho string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		switch extensao := filepath.Ext(caminho); extensao {
		case ".zip", ".rar":
			// Incrementa o WaitGroup
			wg.Add(1)

			// Inicia uma goroutine para processar o arquivo
			go func(caminho, pastaOrigem string) {
				// Libera o semáforo quando a goroutine terminar
				defer func() { <-semaphore }()
				defer wg.Done()

				// Adquire um espaço no semáforo
				semaphore <- struct{}{}

				var err error
				switch extensao {
				case ".zip":
					err = DescompactarZip(caminho, pastaOrigem)
				case ".rar":
					err = DescompactarRar(caminho, pastaOrigem)
				}

				if err != nil {
					errCh <- fmt.Errorf("Erro ao descompactar %s: %s", caminho, err)
				}

				// Remover o arquivo após a descompactação
				err = os.Remove(caminho)
				if err != nil {
					errCh <- fmt.Errorf("Erro ao remover %s: %s", caminho, err)
				}
			}(caminho, pastaOrigem)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Erro ao percorrer a pasta:", err)
	}

	// Função anônima para fechar o canal de conclusão quando todas as goroutines terminarem
	go func() {
		wg.Wait()
		close(doneCh)
	}()

	// Aguarda a conclusão das goroutines e verifica erros
	select {
	case <-doneCh:
		// fmt.Println("Descompactação concluída.")
	case err := <-errCh:
		if err != nil {
			fmt.Println(err)
		}
	}
}
