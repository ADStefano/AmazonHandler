package mock

import (
	"amazon-handler/s3handler"
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// Mock da inicialização do S3
func NewS3ClientMock(mock s3handler.S3Api) *s3handler.Client {

	log.Println("Carregando interface mock")

	return &s3handler.Client{
		S3Client: mock,
		PresignerClient: &MockPresigner{
			PresignGetObjectFunc: func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
				return &v4.PresignedHTTPRequest{}, nil
			},
			PresignPutObjectFunc: func(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
				return &v4.PresignedHTTPRequest{}, nil
			},
			PresignDeleteBucketFunc: func(ctx context.Context, params *s3.DeleteBucketInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
				return &v4.PresignedHTTPRequest{}, nil
			},
			PresignDeleteObjectFunc: func(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
				return &v4.PresignedHTTPRequest{}, nil
			},
		},
		Paginator: func(input *s3.ListObjectsV2Input) s3handler.S3Paginator {
			return &MockPaginator{
				Pages: []*s3.ListObjectsV2Output{
					{
						Contents: []types.Object{
							{Key: aws.String("exemplo.html"), Size: aws.Int64(2048), LastModified: aws.Time(time.Now().Add(-24 * time.Hour))},
						},
					},
				},
				Index: 0,
			}
		},
		ObjNotExistWaiter: func() s3handler.S3NewObjectNotExists {
			return &MockObjectNotExists{
				WaitFunc: func(ctx context.Context, params *s3.HeadObjectInput, maxWaitDur time.Duration, optFns ...func(*s3.ObjectNotExistsWaiterOptions)) error {
					errorBucket := "bucket-timeout"
					pntrErrorBucket := &errorBucket
					if *params.Bucket == *pntrErrorBucket {
						return s3handler.ErrWaiterTimeout
					}
					return nil
				},
			}
		},
		BucketNotExistsWaiter: func() s3handler.S3NewBucketNotExists {
			return &MockBucketNotExists{
				WaitFunc: func(ctx context.Context, params *s3.HeadBucketInput, maxWaitDur time.Duration, optFns ...func(*s3.BucketNotExistsWaiterOptions)) error {
					errorBucket := "bucket-timeout"
					pntrErrorBucket := &errorBucket
					if *params.Bucket == *pntrErrorBucket {
						return s3handler.ErrWaiterTimeout
					}
					return nil
				},
			}
		},
	}

}

// Cria o mock do S3Client
func CreateS3ClientMock() *s3handler.Client {
	mock := &MockS3Client{
		CreateBucketFunc: func(ctx context.Context, input *s3.CreateBucketInput, opts ...func(*s3.Options)) (*s3.CreateBucketOutput, error) {
			switch *input.Bucket {
			case "bucket-exists":
				return nil, s3handler.ErrExists
			case "bucket-owned":
				return nil, s3handler.ErrOwned
			}
			return &s3.CreateBucketOutput{}, nil
		},
		DeleteBucketFunc: func(ctx context.Context, input *s3.DeleteBucketInput, opts ...func(*s3.Options)) (*s3.DeleteBucketOutput, error) {
			if *input.Bucket == "no-bucket" {
				return nil, s3handler.ErrNoSuchBucket
			}
			return &s3.DeleteBucketOutput{}, nil
		},
		DeleteObjectsFunc: func(ctx context.Context, input *s3.DeleteObjectsInput, opts ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error) {
			if *input.Bucket == "no-bucket" {
				return nil, s3handler.ErrNoSuchBucket
			}
			pntrBoolTrue := true
			key := "teste"
			versionId := "versao Teste"
			return &s3.DeleteObjectsOutput{
				Deleted: []types.DeletedObject{
					{
						DeleteMarker: &pntrBoolTrue,
						Key:          &key,
						VersionId:    &versionId,
					},
				},
			}, nil
		},
		HeadBucketFunc: func(ctx context.Context, input *s3.HeadBucketInput, opts ...func(*s3.Options)) (*s3.HeadBucketOutput, error) {
			if *input.Bucket != "bucket-still-exists" {
				return nil, s3handler.ErrNotFound
			}
			return &s3.HeadBucketOutput{}, nil
		},
		PutObjectFunc: func(ctx context.Context, input *s3.PutObjectInput, opts ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
			if *input.Bucket == "entity-too-large" {
				return nil, s3handler.ErrEntityTooLarge
			}
			return &s3.PutObjectOutput{}, nil
		},
		GetObjectFunc: func(ctx context.Context, input *s3.GetObjectInput, opts ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
			switch *input.Bucket {
			case "no-such-bucket":
				return nil, s3handler.ErrNoSuchBucket
			case "no-such-key":
				return nil, s3handler.ErrNoSuchKey
			}
			return &s3.GetObjectOutput{}, nil
		},
	}

	return NewS3ClientMock(mock)
}
