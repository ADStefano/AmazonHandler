package s3handler

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

/* 
Upload faz upload do objeto para o bucket S3 (Objetos de até 5GB) e cria um bucket baseado no prefix se não existir.
O parâmetro prefix é opcional e pode ser uma string vazia.
O prefixo não deve conter barras ("/") no início ou no final.
*/
func (client *Client) UploadS3(bucketName, prefix, filename string, file *os.File, ctx context.Context) (bool, error) {

	if bucketName == "" {
		return false, &S3Error{
			Operation: "upload",
			Bucket:    bucketName,
			Object:    filename,
			Message:   "ErrEmptyParam",
			Err:       ErrEmptyParam,
		}
	}

	key := filename
	if prefix != "" {

		key = prefix + "/" + key
	}

	input := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   file,
	}

	_, err := client.S3Client.PutObject(ctx, input)
	if err != nil {

		parsedErr := ParseError(err)

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
