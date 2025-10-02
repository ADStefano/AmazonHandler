package s3handler

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Download baixa um objeto do bucket S3
func (client *Client) DownloadS3(bucketName, objectKey string, ctx context.Context) (*s3.GetObjectOutput, error) {

	if bucketName == "" || objectKey == "" {
		return nil, &S3Error{
			Operation: "Download",
			Bucket:    bucketName,
			Object:    objectKey,
			Message:   "EmptyParam",
			Err:       ErrEmptyParam,
		}
	}

	log.Printf("Baixando objeto: %s do bucket: %s", objectKey, bucketName)

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	output, err := client.S3Client.GetObject(ctx, input)

	if err != nil {

		log.Printf("Erro: %s", err)

		parsedErr := ParseError(err)

		return nil, &S3Error{
			Operation: "Download",
			Bucket:    bucketName,
			Object:    objectKey,
			Message:   "GetObjectError",
			Err:       parsedErr,
		}

	}

	return output, nil

}
