package s3handler

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// GetPreSignedURL gera uma URL GET pré assinada, o tempo de expiração padrão é de 5 minutos
func (client *Client) GetPreSignedURL(bucketName, objectKey string, expiration time.Duration, ctx context.Context) (*v4.PresignedHTTPRequest, error) {

	if bucketName == "" || objectKey == "" {
		log.Printf("Bucket ou objeto não podem ser vazios")
		return nil, ErrEmptyParam
	}

	log.Printf("Gerando Presigned Get URL para o objeto: %s no bucket: %s com expiração de: %s", objectKey, bucketName, expiration)

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

	request, err := client.PresignerClient.PresignGetObject(ctx, params, opts)
	if err != nil {
		log.Printf("Erro ao gerar URL pré assinada: %s", err.Error())
		return nil, err
	}

	return request, err
}

// PutPreSignedURL gera uma URL PUT pré assinada, o tempo de expiração padrão é de 5 minutos
func (client *Client) PutPreSignedURL(bucketName, objectKey string, expiration time.Duration, ctx context.Context) (*v4.PresignedHTTPRequest, error) {

	if bucketName == "" || objectKey == "" {
		log.Printf("Bucket ou objeto não podem ser vazios")
		return nil, ErrEmptyParam
	}

	log.Printf("Gerando Presigned Put URL para o objeto: %s no bucket: %s com expiração de: %s", objectKey, bucketName, expiration)

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

	request, err := client.PresignerClient.PresignPutObject(ctx, params, opts)

	return request, err
}

// DeleteObjectPreSignedURL gera uma URL DELETE pré assinada para objetos dentro de um bucket, o tempo de expiração padrão é de 5 minutos
func (client *Client) DeleteObjectPreSignedURL(bucketName, objectKey string, expiration time.Duration, ctx context.Context) (*v4.PresignedHTTPRequest, error) {

	if bucketName == "" || objectKey == "" {
		log.Printf("Bucket ou objeto não podem ser vazios")
		return nil, ErrEmptyParam
	}

	log.Printf("Gerando Presigned Delete URL para o objeto: %s no bucket: %s com expiração de: %s", objectKey, bucketName, expiration)

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

	request, err := client.PresignerClient.PresignDeleteObject(ctx, params, opts)

	return request, err
}

// DeleteObjectPreSignedURL gera uma URL DELETE pré assinada para buckets, o tempo de expiração padrão é de 5 minutos
func (client *Client) DeleteBucketPreSignedURL(bucketName string, expiration time.Duration, ctx context.Context) (*v4.PresignedHTTPRequest, error) {

	if bucketName == "" {
		log.Printf("Bucket ou objeto não podem ser vazios")
		return nil, ErrEmptyParam
	}

	log.Printf("Gerando Presigned Delete URL para o bucket: %s com expiração de: %s", bucketName, expiration)

	params := &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	}

	if expiration <= 0 {
		log.Printf("Valor de expiração menor ou igual a zero, utilizando padrão de 5 minutos")
		expiration = 5 * time.Minute
	}

	opts := func(o *s3.PresignOptions) {
		o.Expires = expiration
	}

	request, err := client.PresignerClient.PresignDeleteBucket(ctx, params, opts)

	return request, err
}

// PostPreSignedURL gera uma URL POST pré assinada, o tempo de expiração padrão é de 5 minutos
func (client *Client) PostPreSignedURL(bucketName, objectKey string, expiration time.Duration, ctx context.Context) (*s3.PresignedPostRequest, error) {

	if bucketName == "" || objectKey == "" {
		log.Printf("Bucket ou objeto não podem ser vazios")
		return nil, ErrEmptyParam
	}

	log.Printf("Gerando Presigned POST URL para o objeto: %s no bucket: %s com expiração de: %s", objectKey, bucketName, expiration)

	params := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	if expiration <= 0 {
		log.Printf("Valor de expiração menor ou igual a zero, utilizando padrão de 5 minutos")
		expiration = 5 * time.Minute
	}

	opts := func(o *s3.PresignPostOptions) {
		o.Expires = expiration
	}

	request, err := client.PresignerClient.PresignPostObject(ctx, params, opts)

	return request, err

}
