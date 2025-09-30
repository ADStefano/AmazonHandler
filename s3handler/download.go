package s3handler

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
)

// Download baixa um objeto do bucket S3
func (client *Client) DownloadS3(bucketName, objectKey string) (*s3.GetObjectOutput, error) {
	log.Printf("Baixando objeto: %s do bucket: %s", objectKey, bucketName)

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	output, err := client.S3Client.GetObject(context.TODO(), input)

	if err != nil {

		log.Printf("Erro: %s", err)

		var errApi smithy.APIError

		if errors.As(err, &ErrNoSuchKey) {
			log.Printf("O objeto: %s não existe no bucket: %s", objectKey, bucketName)
			return nil, ErrNoSuchKey
		} else if errors.As(err, &errApi) && errApi.ErrorCode() == "NoSuchBucket" {
			log.Printf("O bucket: %s não existe", bucketName)
			return nil, ErrNoSuchBucket
		} else {
			log.Printf("Erro ao baixar o objeto: %s do bucket: %s, erro: %s", objectKey, bucketName, err)
			return nil, err
		}

	}

	return output, nil

}
