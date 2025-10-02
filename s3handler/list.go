package s3handler

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// ListBuckets lista os buckets do usuário autenticado
func (client *Client) ListBuckets(prefix string, ctx context.Context) ([]types.Bucket, error) {

	log.Printf("Buscando buckets...")

	var buckets []types.Bucket

	params := &s3.ListBucketsInput{
		Prefix: aws.String(prefix),
	}

	bucketPaginator := client.BucketPaginator(params)

	for bucketPaginator.HasMorePages() {

		output, err := bucketPaginator.NextPage(ctx)

		if err != nil {

			parsedErr := ParseError(err)

			return nil, &S3Error{
				Operation: "ListBuckets",
				Message:   "ListBucketsError",
				Err:       parsedErr,
			}
		}

		buckets = append(buckets, output.Buckets...)
	}

	log.Printf("Total de buckets encontrados: %d", len(buckets))

	return buckets, nil
}

// ListObjects lista os objetos dentro de um bucket
func (client *Client) ListObjects(bucketName, prefix string, maxKeys int32, ctx context.Context) ([]types.Object, error) {
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

		output, err := paginator.NextPage(ctx)

		if err != nil {

			parsedErr := ParseError(err)

			return nil, &S3Error{
				Operation: "ListObjects",
				Bucket:    bucketName,
				Message:   "ListObjectsError",
				Err:       parsedErr,
			}

		} else {

			objects = append(objects, output.Contents...)
		}
	}

	return objects, nil
}
