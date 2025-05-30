package s3

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Api interface {
	CreateBucket(ctx context.Context, input *s3.CreateBucketInput, opts ...func(*s3.Options)) (*s3.CreateBucketOutput, error)
	DeleteBucket(ctx context.Context, input *s3.DeleteBucketInput, opts ...func(*s3.Options)) (*s3.DeleteBucketOutput, error)
	DeleteObjects(ctx context.Context, params *s3.DeleteObjectsInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error)
	ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, opts ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
	HeadBucket(ctx context.Context, input *s3.HeadBucketInput, opts ...func(*s3.Options)) (*s3.HeadBucketOutput, error)
	HeadObject(ctx context.Context, input *s3.HeadObjectInput, opts ...func(*s3.Options)) (*s3.HeadObjectOutput, error)
	PutObject(ctx context.Context, input *s3.PutObjectInput, opts ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

type S3Paginator interface {
	HasMorePages() bool
	NextPage(ctx context.Context, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
}

type CreatePaginator func(input *s3.ListObjectsV2Input) S3Paginator

type S3NewObjectNotExists interface {
	Wait(ctx context.Context, params *s3.HeadObjectInput, maxWaitDur time.Duration, optFns ...func(*s3.ObjectNotExistsWaiterOptions)) error
	WaitForOutput(ctx context.Context, params *s3.HeadObjectInput, maxWaitDur time.Duration, optFns ...func(*s3.ObjectNotExistsWaiterOptions)) (*s3.HeadObjectOutput, error)
}

type CreateNewObjectNotExists func() S3NewObjectNotExists

type S3NewBucketNotExists interface {
	Wait(ctx context.Context, params *s3.HeadBucketInput, maxWaitDur time.Duration, optFns ...func(*s3.BucketNotExistsWaiterOptions)) error
	WaitForOutput(ctx context.Context, params *s3.HeadBucketInput, maxWaitDur time.Duration, optFns ...func(*s3.BucketNotExistsWaiterOptions)) (*s3.HeadBucketOutput, error)
}

type CreateNewBucketNotExists func() S3NewBucketNotExists

// Estrutura do client do S3
type Client struct {
	s3Client  S3Api
	paginator CreatePaginator
	objNotExistWaiter CreateNewObjectNotExists
	bucketNotExistsWaiter CreateNewBucketNotExists
}

// NewS3Client inicializa um client S3
func NewS3Client() *Client {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Erro ao carregar as configurações. (%e)", err)
	}

	client := s3.NewFromConfig(cfg)

	return &Client{
		s3Client: client,
		paginator: func(input *s3.ListObjectsV2Input) S3Paginator {
			return s3.NewListObjectsV2Paginator(client, input)
		},
		objNotExistWaiter: func() S3NewObjectNotExists {
			return s3.NewObjectNotExistsWaiter(client)
		},
		bucketNotExistsWaiter: func() S3NewBucketNotExists {
			return s3.NewBucketNotExistsWaiter(client)
		},
	}

}
