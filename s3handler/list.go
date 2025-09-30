package s3handler

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
)

// ListBuckets lista os buckets do usuário autenticado
func (client *Client) ListBuckets(prefix string) ([]types.Bucket, error) {

	log.Printf("Buscando buckets...")

	var buckets []types.Bucket

	params := &s3.ListBucketsInput{
		Prefix: aws.String(prefix),
	}

	bucketPaginator := client.BucketPaginator(params)

	for bucketPaginator.HasMorePages() {

		output, err := bucketPaginator.NextPage(context.TODO())

		if err != nil {

			var errApi smithy.APIError

			if errors.As(err, &errApi) && errApi.ErrorCode() == "AccessDenied" {

				log.Printf("Acesso negado ao listar buckets: %s\n", errApi.ErrorMessage())

				return nil, ErrAccessDenied

			} else {

				log.Printf("Erro ao listar buckets: %s\n", err)

				return nil, err
			}
		}

		buckets = append(buckets, output.Buckets...)
	}

	log.Printf("Total de buckets encontrados: %d", len(buckets))

	return buckets, nil
}

// ListObjects lista os objetos dentro de um bucket
func (client *Client) ListObjects(bucketName, prefix string, maxKeys int32) ([]types.Object, error) {
	log.Printf("Buscando objetos no bucket %s", bucketName)

	if bucketName == "" {
		log.Printf("Nome do bucket não pode ser vazio")
		return nil, ErrEmptyParam
	}

	var objects []types.Object

	if maxKeys < 1 {
		log.Printf("Parâmetro maxKeys menor que o permitido, utilizando default de 1000")
		maxKeys = 1000
	}

	if prefix != "" {
		log.Printf("Utilizando prefixo: %s", prefix)
	}

	params := &s3.ListObjectsV2Input{
		Bucket:  aws.String(bucketName),
		MaxKeys: &maxKeys,
		Prefix:  aws.String(prefix),
	}

	paginator := client.ObjectPaginator(params)

	for paginator.HasMorePages() {

		output, err := paginator.NextPage(context.TODO())

		if err != nil {

			if errors.As(err, &ErrNoSuchBucket) {

				log.Printf("Bucket %s não existe.\n", bucketName)
				return nil, err

			} else {

				log.Printf("Erro ao buscar os objetos no bucket %s: %s\n", bucketName, err)
				return nil, err
			}

		} else {

			objects = append(objects, output.Contents...)
		}
	}

	return objects, nil
}
