package s3

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// ListBuckets lista os buckets
func (client *Client) ListBuckets() (string, error) {
	return "", nil
}

// ListObjects lista os objetos dentro de um bucket TODO ADD PREFIX
func (client *Client) ListObjects(bucketName string, maxKeys int32) ([]types.Object, error) {
	log.Printf("Buscando objetos no bucket %s", bucketName)

	var objects []types.Object

	if maxKeys < 1 {
		log.Printf("Parâmetro maxKeys menor que o permitido, utilizando default de 1000")
		maxKeys = 1000
	}

	params := &s3.ListObjectsV2Input{
		Bucket:  aws.String(bucketName),
		MaxKeys: &maxKeys,
	}

	paginator := client.paginator(params)

	for paginator.HasMorePages() {

		output, err := paginator.NextPage(context.TODO())

		if err != nil {

			var noBucket *types.NoSuchBucket

			if errors.As(err, &noBucket) {

				log.Printf("Bucket %s não existe.\n", bucketName)
				return nil, err

			} else {

				log.Printf("Erro ao buscar os objetos no bucket: %s\n", err)
				log.Panicf("Erro: %e", err)
				return nil, err
			}

		} else {

			objects = append(objects, output.Contents...)
		}
	}

	return objects, nil
}
