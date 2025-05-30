package s3

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var Exists *types.BucketAlreadyExists
var Owned *types.BucketAlreadyOwnedByYou
var NoBucket *types.NoSuchBucket
var ErrWaiterTimeout = errors.New("exceeded max wait time for BucketNotExists waiter")

// CreateBucket cria um bucket S3 caso não exista
func (client *Client) CreateBucket(bucketName string) (bool, error) {

	log.Printf("Criando bucket: %s\n", bucketName)

	params := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint("sa-east-1"),
		},
	}
	_, err := client.s3Client.CreateBucket(context.TODO(), params)
	if err != nil {

		if errors.As(err, &Owned) {
			log.Printf("Você já criou um bucket com esse nome: %s", bucketName)
			return false, Owned
		} else if errors.As(err, &Exists) {
			log.Printf("Bucket: %s já existe", bucketName)
			return false, Exists
		} else {
			log.Printf("Não foi possível criar o bucket %s", err.Error())
			return false, err
		}
	}

	log.Printf("Bucket %s criado com sucesso", bucketName)
	return true, nil
}

// DeleteBucket apaga um bucket do S3 (O bucket precisa estar vazio)
func (client *Client) DeleteBucket(bucketName string) (bool, error) {
	log.Printf("Deletando bucket %s\n", bucketName)

	params := &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	}

	_, err := client.s3Client.DeleteBucket(context.TODO(), params)
	if err != nil {

		if errors.As(err, &NoBucket) {
			log.Printf("Bucket %s não encontrado", bucketName)
			return false, NoBucket
		} else{
			log.Printf("Erro ao deletar bucket %s: %v", bucketName, err)
			return false, err
		}
	} else {
		headBucketParams := &s3.HeadBucketInput{
			Bucket: aws.String(bucketName),
		}
		
		err = client.bucketNotExistsWaiter().Wait(context.TODO(), headBucketParams, time.Minute)
		if err != nil {
			log.Printf("Erro ao esperar bucket %s ser deletado", bucketName)
			return false, ErrWaiterTimeout
		}
	}

	log.Printf("Bucket %s deletado com sucesso", bucketName)
	return true, nil
}
