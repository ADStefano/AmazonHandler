package s3handler

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// GetPreSignedURL gera uma URL GET pré assinada, o tempo de expiração padrão é de 5 minutos
func (client *Client) GetPreSignedURL(bucketName, objectKey string, expiration time.Duration, ctx context.Context) (*v4.PresignedHTTPRequest, error) {

	if bucketName == "" || objectKey == "" {
		return nil, &S3Error{
			Operation: "PresignGetObject",
			Bucket:    bucketName,
			Object:    objectKey,
			Message:   "EmptyParam",
			Err:       ErrEmptyParam,
		}
	}

	params := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	if expiration <= 0 {
		expiration = 5 * time.Minute
	}

	opts := func(o *s3.PresignOptions) {
		o.Expires = expiration
	}

	request, err := client.PresignerClient.PresignGetObject(ctx, params, opts)
	if err != nil {
		parsedErr := ParseError(err)
		return nil, &S3Error{
			Operation: "PresignGetObject",
			Bucket:    bucketName,
			Object:    objectKey,
			Message:   "PresignGetObjectError",
			Err:       parsedErr,
		}
	}

	return request, nil
}

// PutPreSignedURL gera uma URL PUT pré assinada, o tempo de expiração padrão é de 5 minutos
func (client *Client) PutPreSignedURL(bucketName, objectKey string, expiration time.Duration, ctx context.Context) (*v4.PresignedHTTPRequest, error) {

	if bucketName == "" || objectKey == "" {
		return nil, &S3Error{
			Operation: "PresignPutObject",
			Bucket:    bucketName,
			Object:    objectKey,
			Message:   "EmptyParam",
			Err:       ErrEmptyParam,
		}
	}

	params := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	if expiration <= 0 {
		expiration = 5 * time.Minute
	}

	opts := func(o *s3.PresignOptions) {
		o.Expires = expiration
	}

	request, err := client.PresignerClient.PresignPutObject(ctx, params, opts)

	if err != nil {
		parsedErr := ParseError(err)
		return nil, &S3Error{
			Operation: "PresignPutObject",
			Bucket:    bucketName,
			Object:    objectKey,
			Message:   "PresignPutObjectError",
			Err:       parsedErr,
		}
	}

	return request, nil
}

// DeleteObjectPreSignedURL gera uma URL DELETE pré assinada para objetos dentro de um bucket, o tempo de expiração padrão é de 5 minutos
func (client *Client) DeleteObjectPreSignedURL(bucketName, objectKey string, expiration time.Duration, ctx context.Context) (*v4.PresignedHTTPRequest, error) {

	if bucketName == "" || objectKey == "" {
		return nil, &S3Error{
			Operation: "PresignDeleteObject",
			Bucket:    bucketName,
			Object:    objectKey,
			Message:   "EmptyParam",
			Err:       ErrEmptyParam,
		}
	}

	params := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	if expiration <= 0 {
		expiration = 5 * time.Minute
	}

	opts := func(o *s3.PresignOptions) {
		o.Expires = expiration
	}

	request, err := client.PresignerClient.PresignDeleteObject(ctx, params, opts)

	if err != nil {
		parsedErr := ParseError(err)
		return nil, &S3Error{
			Operation: "PresignDeleteObject",
			Bucket:    bucketName,
			Object:    objectKey,
			Message:   "PresignDeleteObjectError",
			Err:       parsedErr,
		}
	}

	return request, nil
}

// DeleteObjectPreSignedURL gera uma URL DELETE pré assinada para buckets, o tempo de expiração padrão é de 5 minutos
func (client *Client) DeleteBucketPreSignedURL(bucketName string, expiration time.Duration, ctx context.Context) (*v4.PresignedHTTPRequest, error) {

	if bucketName == "" {
		return nil, &S3Error{
			Operation: "PresignDeleteBucket",
			Bucket:    bucketName,
			Message:   "EmptyParam",
			Err:       ErrEmptyParam,
		}
	}

	params := &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	}

	if expiration <= 0 {
		expiration = 5 * time.Minute
	}

	opts := func(o *s3.PresignOptions) {
		o.Expires = expiration
	}

	request, err := client.PresignerClient.PresignDeleteBucket(ctx, params, opts)

	if err != nil {
		parsedErr := ParseError(err)
		return nil, &S3Error{
			Operation: "PresignDeleteBucket",
			Bucket:    bucketName,
			Message:   "PresignDeleteBucketError",
			Err:       parsedErr,
		}
	}

	return request, nil
}

// PostPreSignedURL gera uma URL POST pré assinada, o tempo de expiração padrão é de 5 minutos
func (client *Client) PostPreSignedURL(bucketName, objectKey string, expiration time.Duration, ctx context.Context) (*s3.PresignedPostRequest, error) {

	if bucketName == "" || objectKey == "" {
		return nil, &S3Error{
			Operation: "PresignPostObject",
			Bucket:    bucketName,
			Object:    objectKey,
			Message:   "EmptyParam",
			Err:       ErrEmptyParam,
		}
	}

	params := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	if expiration <= 0 {
		expiration = 5 * time.Minute
	}

	opts := func(o *s3.PresignPostOptions) {
		o.Expires = expiration
	}

	request, err := client.PresignerClient.PresignPostObject(ctx, params, opts)

	if err != nil {
		parsedErr := ParseError(err)
		return nil, &S3Error{
			Operation: "PresignPostObject",
			Bucket:    bucketName,
			Object:    objectKey,
			Message:   "PresignPostObjectError",
			Err:       parsedErr,
		}
	}

	return request, nil

}
