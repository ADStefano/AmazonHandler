package s3handler

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Upload faz upload do objeto para o bucket S3 (Objetos de até 5GB) e cria um bucket baseado no prefix se não existir.
func (client *Client) UploadS3(bucketName, prefix, path string, ctx context.Context) (bool, error) {

	if bucketName == "" || path == "" {
		log.Printf("Nome do bucket ou caminho do arquivo vazio")
		return false, &S3Error{
			Operation: "upload",
			Bucket:    bucketName,
			Object:    path,
			Message:   "ErrEmptyParam",
			Err:       ErrEmptyParam,
		}
	}

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

	_, err = client.S3Client.PutObject(ctx, input)
	if err != nil {

		parsedErr := ParseError(err)

		log.Printf("Erro ao fazer upload do objeto: %s, para o bucket: %s, prefixo: %s, erro: %s", filename, bucketName, prefix, err)

		return false, &S3Error{
			Operation: "upload",
			Bucket:    bucketName,
			Object:    key,
			Message:   "PutObjectError",
			Err:       parsedErr,
		}
	}

	return true, nil
}
