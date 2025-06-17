package s3

import (
	"context"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
)

// Upload faz upload do objeto para o bucket S3 (Objetos de até 5GB).
func (client *Client) Upload(bucketName, prefix, path string) (bool, error) {

	filename := filepath.Base(path)
	log.Printf("Fazendo upload do arquivo: %s para o bucket: %s", filename, bucketName)

	file, err := os.Open(path)
	if err != nil {
		log.Printf("Erro ao abrir arquivo: %s", filename)
		return false, err
	}
	defer file.Close()

	key := filename
	if prefix != "" {
		key = prefix + "/" + filename
	}

	input := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   file,
	}

	println(bucketName)
	println(key)
	print(input)

	_, err = client.s3Client.PutObject(context.TODO(), input)
	if err != nil {

		var apiErr smithy.APIError

		if errors.As(err, &apiErr) && apiErr.ErrorCode() == "EntityTooLarge" {
			log.Printf("O arquivo %s é muito grande para ser enviado para o bucket %s", filename, bucketName)
			return false, errors.New("EntityTooLarge")
		}

		log.Printf("Erro ao fazer upload do objeto: %s, para o bucket: %s, prefixo: %s, erro: %s", filename, bucketName, prefix, err)
		return false, err
	}

	return true, nil
}
