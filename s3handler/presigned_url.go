package s3handler

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// GetPreSignedURL gera uma URL pré assinada, o tempo de expiração padrão é de 5 minutos 
func (client *Client) GetPreSignedURL(bucketName, objectKey string, expiration time.Duration) (*v4.PresignedHTTPRequest, error) {

	if bucketName == "" || objectKey == "" {
		log.Printf("Bucket ou objeto não podem ser vazios")
		return nil, ErrEmptyParam
	}

	log.Printf("Gerando Presigned URL para o objeto: %s no bucket: %s com expiração de: %s", objectKey, bucketName, expiration)

	params := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	if expiration <= 0 {
		log.Printf("Valor de expiração menor ou igual a zero, utilizando padrão de 5 minutos")
		expiration = 5 * time.Minute
	}

	opts := func(o *s3.PresignOptions) {
		o.Expires = expiration
	}

	request, err := client.PresignerClient.PresignGetObject(context.TODO(),params, opts)
	if err != nil {
		log.Printf("Erro ao gerar URL pré assinada: %s", err.Error())
		return nil, err
	}

	return request, err
}

func (client *Client) PutPreSignedURL(bucketName, objectKey string, expiration time.Duration) (*v4.PresignedHTTPRequest, error) {

	if bucketName == "" || objectKey == "" {
		log.Printf("Bucket ou objeto não podem ser vazios")
		return nil, ErrEmptyParam
	}

	log.Printf("Gerando Presigned URL para o objeto: %s no bucket: %s com expiração de: %s", objectKey, bucketName, expiration)

	params := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	if expiration <= 0 {
		log.Printf("Valor de expiração menor ou igual a zero, utilizando padrão de 5 minutos")
		expiration = 5 * time.Minute
	}

	opts := func(o *s3.PresignOptions) {
		o.Expires = expiration
	}

	request, err := client.PresignerClient.PresignPutObject(context.TODO(),params, opts)

	return request, err
}

func (client *Client) DeletePreSignedURL(bucketName, objectKey string, expiration time.Duration) (*v4.PresignedHTTPRequest, error) {

	if bucketName == "" || objectKey == "" {
		log.Printf("Bucket ou objeto não podem ser vazios")
		return nil, ErrEmptyParam
	}

	log.Printf("Gerando Presigned URL para o objeto: %s no bucket: %s com expiração de: %s", objectKey, bucketName, expiration)

	params := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	if expiration <= 0 {
		log.Printf("Valor de expiração menor ou igual a zero, utilizando padrão de 5 minutos")
		expiration = 5 * time.Minute
	}

	opts := func(o *s3.PresignOptions) {
		o.Expires = expiration
	}

	request, err := client.PresignerClient.PresignDeleteObject(context.TODO(),params, opts)

	return request, err
}
