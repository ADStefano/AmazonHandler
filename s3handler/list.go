package s3handler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// ListBuckets lista os buckets do usuário autenticado
func (client *Client) ListBuckets(prefix string, ctx context.Context) ([]types.Bucket, error) {

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

	return buckets, nil
}

// ListObjects lista os objetos dentro de um bucket, caso o maxKeys seja menor que 1, o valor padrão será 1000
func (client *Client) ListObjects(bucketName, prefix string, maxKeys int32, ctx context.Context) ([]types.Object, error) {

	if bucketName == "" {
		return nil, ErrEmptyParam
	}

	var objects []types.Object

	if maxKeys < 1 {
		maxKeys = 1000
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
