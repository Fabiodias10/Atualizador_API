package listarAws

import (
	"fmt"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// lista o bucket: setup-dias/API-OTHERS/ e retorna os objetos "
func Listar() ([]string, []string, error) {
	// Configuração da sessão AWS
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		// Você também pode configurar suas credenciais aqui, ou elas serão obtidas automaticamente do ambiente AWS
	})
	if err != nil {
		fmt.Println("Erro ao criar a sessão:", err)
		return nil, nil, err
	}

	// Nome do bucket
	bucket := "setup-dias"
	prefix := "API-OTHERS/"

	// Criar o serviço S3
	svc := s3.New(sess)

	// Configurar os parâmetros para a listagem
	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	}

	// Realizar a operação de listagem
	resp, err := svc.ListObjectsV2(params)
	if err != nil {
		fmt.Println("Erro ao listar objetos no S3:", err)
		return nil, nil, err
	}

	// Exibir a lista de arquivos
	// fmt.Println("Arquivos no diretório s3://" + bucket + "/" + prefix)
	var nomesArquivos []string
	var nomesCompleto []string
	for indice, item := range resp.Contents {
		// fmt.Println(*item.Key)
		// fmt.Println(*item.Key)
		if indice > 0 {

			// formata := *item.Key
			formata := filepath.Base(*item.Key)
			// fmt.Println(formata)
			nomesArquivos = append(nomesArquivos, formata)
			nomesCompleto = append(nomesCompleto, *item.Key)
			// fmt.Println(item)
		}
	}

	return nomesCompleto, nomesArquivos, nil
}
