package mock

import (
	"github.com/ADStefano/AmazonHandler/s3handler"
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// Mock da inicialização do S3
func NewS3ClientMock(mock s3handler.S3Api) *s3handler.Client {

	return &s3handler.Client{
		S3Client: mock,
		PresignerClient: &MockPresigner{
			PresignGetObjectFunc: func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
				if *params.Bucket == "test-bucket-success"{
					return &v4.PresignedHTTPRequest{
						URL:    "https://httpbin.org/status/200",
						Method: "GET",
						SignedHeader: map[string][]string{
							"Host":                 {"test-bucket.s3.amazonaws.com"},
							"X-Amz-Content-Sha256": {"UNSIGNED-PAYLOAD"},
							"X-Amz-Date":           {"20240605T123456Z"},
							"X-Amz-Expires":        {"300"},
							"X-Amz-SignedHeaders":  {"host"},
							"X-Amz-Signature":      {"EXAMPLE_SIGNATURE"},
						},
					}, nil
				}
				return &v4.PresignedHTTPRequest{}, nil
			},
			PresignPutObjectFunc: func(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
				if *params.Bucket == "test-bucket-success"{
					return &v4.PresignedHTTPRequest{
						URL:    "https://httpbin.org/status/200",
						Method: "PUT",
						SignedHeader: map[string][]string{
							"Host":                 {"test-bucket.s3.amazonaws.com"},
							"X-Amz-Content-Sha256": {"UNSIGNED-PAYLOAD"},
							"X-Amz-Date":           {"20240605T123456Z"},
							"X-Amz-Expires":        {"300"},
							"X-Amz-SignedHeaders":  {"host"},
							"X-Amz-Signature":      {"EXAMPLE_SIGNATURE"},
						},
					}, nil
				}
				return &v4.PresignedHTTPRequest{}, nil
			},
			PresignDeleteBucketFunc: func(ctx context.Context, params *s3.DeleteBucketInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
				if *params.Bucket == "test-bucket-success"{
					return &v4.PresignedHTTPRequest{
						URL:    "https://httpbin.org/status/200",
						Method: "DELETE",
						SignedHeader: map[string][]string{
							"Host":                 {"test-bucket.s3.amazonaws.com"},
							"X-Amz-Content-Sha256": {"UNSIGNED-PAYLOAD"},
							"X-Amz-Date":           {"20240605T123456Z"},
							"X-Amz-Expires":        {"300"},
							"X-Amz-SignedHeaders":  {"host"},
							"X-Amz-Signature":      {"EXAMPLE_SIGNATURE"},
						},
					}, nil
				}
				return &v4.PresignedHTTPRequest{}, nil
			},
			PresignDeleteObjectFunc: func(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
				if *params.Bucket == "test-bucket-success"{
					return &v4.PresignedHTTPRequest{
						URL:    "https://httpbin.org/status/200",
						Method: "DELETE",
						SignedHeader: map[string][]string{
							"Host":                 {"test-bucket.s3.amazonaws.com"},
							"X-Amz-Content-Sha256": {"UNSIGNED-PAYLOAD"},
							"X-Amz-Date":           {"20240605T123456Z"},
							"X-Amz-Expires":        {"300"},
							"X-Amz-SignedHeaders":  {"host"},
							"X-Amz-Signature":      {"EXAMPLE_SIGNATURE"},
						},
					}, nil
				}	
				return &v4.PresignedHTTPRequest{}, nil
			},
			PresignPostObjectFunc: func(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.PresignPostOptions)) (*s3.PresignedPostRequest, error) {
				if *params.Bucket == "test-bucket-success"{
					return &s3.PresignedPostRequest{
						URL: "https://httpbin.org/status/200",
						Values: map[string]string{
							"key":             *params.Key,
							"bucket":          *params.Bucket,
							"x-amz-algorithm": "AWS4-HMAC-SHA256",
							"x-amz-credential": "EXAMPLECREDENTIAL/20240605/us-east-1/s3/aws4_request",
							"x-amz-date":      "20240605T123456Z",
							"x-amz-signature": "EXAMPLESIGNATURE",
						},
					}, nil
				}				
				return &s3.PresignedPostRequest{}, nil
			},
		},
		ObjectPaginator: func(input *s3.ListObjectsV2Input) s3handler.S3ObjectPaginator {
			if *input.Bucket == "no-such-bucket" {
				return &MockObjectPaginator{
					Err: s3handler.ErrNoSuchBucket,
					Pages: []*s3.ListObjectsV2Output{
						{
							Contents: []types.Object{
								{Key: aws.String("exemplo.html"), Size: aws.Int64(2048), LastModified: aws.Time(time.Now().Add(-24 * time.Hour))},
							},
						},
					},
					Index: 0,
			}
			}
			return &MockObjectPaginator{
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
		BucketPaginator: func(input *s3.ListBucketsInput) s3handler.S3BucketPaginator {
			if *input.Prefix == "access-denied"{
				return &MockBucketPaginator{
					Err: s3handler.ErrAccessDenied,
					Pages: []*s3.ListBucketsOutput{
						{
							Buckets: []types.Bucket{
								{Name: aws.String("bucket-teste"), CreationDate: aws.Time(time.Now().Add(-240 * time.Hour))},
							},
						},
					},
					Index: 0,
				}
			}
			return &MockBucketPaginator{
				Err: nil,
				Pages: []*s3.ListBucketsOutput{
					{
						Buckets: []types.Bucket{
							{Name: aws.String("bucket-teste"), CreationDate: aws.Time(time.Now().Add(-240 * time.Hour))},
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
