package downloadAws

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Download(objetos []string) {

	dataAtual := time.Now().Format("02-01-2006")

	// Caminho desejado para a pasta
	caminho := filepath.Join("C:/temp", "API-"+dataAtual)

	// Verificar se a pasta já existe
	_, err := os.Stat(caminho)
	if os.IsNotExist(err) {
		// Criar a pasta se ela não existir
		err := os.MkdirAll(caminho, os.ModePerm)
		if err != nil {
			fmt.Println("Erro ao criar pasta", err)
		}
	}

	// Configuração da sessão AWS
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		// Você também pode configurar suas credenciais aqui, ou elas serão obtidas automaticamente do ambiente AWS
	})
	if err != nil {
		fmt.Println("Erro ao criar a sessão:", err)
		return
	}

	// Nome do bucket e caminho do objeto no S3
	bucket := "setup-dias"
	// Prefixo do caminho no S3
	prefixoCaminho := "API-OTHERS/"
	// key := "Totem-update/SitDemo.zip"

	// Criar o serviço S3
	svc := s3.New(sess)
	// Iterar sobre os objetos e realizar o download
	for _, objeto := range objetos {
		fmt.Println("Baixando:", objeto)
		// Combinar o prefixo do caminho com o nome do objeto
		caminhoCompleto := prefixoCaminho + objeto
		// Configurar os parâmetros de download
		params := &s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(caminhoCompleto),
		}

		// Realizar a operação de download
		resp, err := svc.GetObject(params)
		if err != nil {
			fmt.Println("Erro ao baixar o arquivo do S3:", err)
			continue // Continuar com o próximo objeto em caso de erro
		}
		// defer resp.Body.Close()

		dataAtual := time.Now().Format("02-01-2006")
		pastaOrigem := filepath.Join("C:/temp", "API-"+dataAtual, filepath.Base(objeto))

		// Criar o arquivo local para escrita
		file, err := os.Create(pastaOrigem)
		if err != nil {
			fmt.Println("Erro ao criar o arquivo local:", err)
			continue
		}
		// defer file.Close()

		//Copiar o conteúdo do corpo da resposta para o arquivo local
		// _, err = io.Copy(file, resp.Body)
		// if err != nil {
		// 	fmt.Println("Erro ao copiar o conteúdo para o arquivo local:", err)
		// 	file.Close()
		// 	continue
		// }
		const bufferSize = 8192 // Tamanho do buffer para cópia

		// ...

		// Copiar o conteúdo do corpo da resposta para o arquivo local usando CopyBuffer
		buffer := make([]byte, bufferSize)
		_, err = io.CopyBuffer(file, resp.Body, buffer)
		if err != nil {
			fmt.Println("Erro ao copiar o conteúdo para o arquivo local:", err)
			file.Close()
			resp.Body.Close()
			continue
		}

		file.Close()
		resp.Body.Close()
	}

}
