package s3handler

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Interface para o cliente S3
type S3Api interface {
	CreateBucket(ctx context.Context, input *s3.CreateBucketInput, opts ...func(*s3.Options)) (*s3.CreateBucketOutput, error)
	DeleteBucket(ctx context.Context, input *s3.DeleteBucketInput, opts ...func(*s3.Options)) (*s3.DeleteBucketOutput, error)
	DeleteObjects(ctx context.Context, params *s3.DeleteObjectsInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error)
	ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, opts ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
	HeadBucket(ctx context.Context, input *s3.HeadBucketInput, opts ...func(*s3.Options)) (*s3.HeadBucketOutput, error)
	HeadObject(ctx context.Context, input *s3.HeadObjectInput, opts ...func(*s3.Options)) (*s3.HeadObjectOutput, error)
	PutObject(ctx context.Context, input *s3.PutObjectInput, opts ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
	GetObject(ctx context.Context, input *s3.GetObjectInput, opts ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

// Interface para o paginador do S3
type S3Paginator interface {
	HasMorePages() bool
	NextPage(ctx context.Context, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
}

// Função para criar um paginador do S3
type CreatePaginator func(input *s3.ListObjectsV2Input) S3Paginator

// Interface para o waiter de objeto não existente do S3
type S3NewObjectNotExists interface {
	Wait(ctx context.Context, params *s3.HeadObjectInput, maxWaitDur time.Duration, optFns ...func(*s3.ObjectNotExistsWaiterOptions)) error
	WaitForOutput(ctx context.Context, params *s3.HeadObjectInput, maxWaitDur time.Duration, optFns ...func(*s3.ObjectNotExistsWaiterOptions)) (*s3.HeadObjectOutput, error)
}

// Função para criar um waiter de objeto não existente do S3
type CreateNewObjectNotExists func() S3NewObjectNotExists

// Interface para o waiter de bucket não existente do S3
type S3NewBucketNotExists interface {
	Wait(ctx context.Context, params *s3.HeadBucketInput, maxWaitDur time.Duration, optFns ...func(*s3.BucketNotExistsWaiterOptions)) error
	WaitForOutput(ctx context.Context, params *s3.HeadBucketInput, maxWaitDur time.Duration, optFns ...func(*s3.BucketNotExistsWaiterOptions)) (*s3.HeadBucketOutput, error)
}

// Função para criar um waiter de bucket não existente do S3
type CreateNewBucketNotExists func() S3NewBucketNotExists

// Interface para o presigner do S3
type Presigner interface {
	PresignGetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignPutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignPostObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.PresignPostOptions)) (*s3.PresignedPostRequest, error)
	PresignDeleteBucket(ctx context.Context, params *s3.DeleteBucketInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
	PresignDeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
}
