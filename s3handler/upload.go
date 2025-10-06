package s3handler

import (
	"context"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Upload faz upload do objeto para o bucket S3 (Objetos de até 5GB) e cria um bucket baseado no prefix se não existir.
func (client *Client) UploadS3(bucketName, prefix, path string, ctx context.Context) (bool, error) {

	if bucketName == "" || path == "" {
		return false, &S3Error{
			Operation: "upload",
			Bucket:    bucketName,
			Object:    path,
			Message:   "ErrEmptyParam",
			Err:       ErrEmptyParam,
		}
	}

	filename := filepath.Base(path)

	file, err := os.Open(path)
	if err != nil {
		return false, &S3Error{
			Operation: "upload",
			Bucket:    bucketName,
			Object:    path,
			Message:   "FileOpenError",
			Err:       err,
		}
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
