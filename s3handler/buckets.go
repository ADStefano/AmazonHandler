package s3handler

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// CreateBucket cria um bucket S3 caso não exista
func (client *Client) CreateBucket(bucketName string, ctx context.Context) (bool, error) {

	if bucketName == "" {
		return false, &S3Error{
			Operation: "CreateBucket",
			Bucket:    bucketName,
			Message:   "EmptyBucketName",
			Err:       ErrEmptyParam,
		}
	}

	params := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint("sa-east-1"),
		},
	}
	_, err := client.S3Client.CreateBucket(ctx, params)
	if err != nil {
		parsedErr := ParseError(err)
		return false, &S3Error{
			Operation: "CreateBucket",
			Bucket:    bucketName,
			Message:   "CreateBucketError",
			Err:       parsedErr,
		}
	}

	return true, nil
}

// DeleteBucket apaga um bucket do S3 (O bucket precisa estar vazio)
func (client *Client) DeleteBucket(bucketName string, ctx context.Context) (bool, error) {

	if bucketName == "" {
		return false, &S3Error{
			Operation: "DeleteBucket",
			Bucket:    bucketName,
			Message:   "EmptyBucketName",
			Err:       ErrEmptyParam,
		}
	}

	params := &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	}

	_, err := client.S3Client.DeleteBucket(ctx, params)
	if err != nil {
		parsedErr := ParseError(err)
		return false, &S3Error{
			Operation: "DeleteBucket",
			Bucket:    bucketName,
			Message:   "DeleteBucketError",
			Err:       parsedErr,
		}
	} else {
		headBucketParams := &s3.HeadBucketInput{
			Bucket: aws.String(bucketName),
		}

		err = client.BucketNotExistsWaiter().Wait(context.TODO(), headBucketParams, time.Minute)
		if err != nil {
			return false, &S3Error{
				Operation: "DeleteBucket",
				Bucket:    bucketName,
				Message:   "WaiterTimeout",
				Err:       ErrWaiterTimeout,
			}
		}
	}

	return true, nil
}
