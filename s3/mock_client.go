package s3

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// Mock do S3 para testes
type MockS3Client struct {
	CreateBucketFunc func(ctx context.Context, input *s3.CreateBucketInput, opts ...func(*s3.Options)) (*s3.CreateBucketOutput, error)

	DeleteBucketFunc  func(ctx context.Context, input *s3.DeleteBucketInput, opts ...func(*s3.Options)) (*s3.DeleteBucketOutput, error)
	DeleteObjectsFunc func(ctx context.Context, input *s3.DeleteObjectsInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error)

	ListObjectsV2Func func(ctx context.Context, input *s3.ListObjectsV2Input, opts ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)

	HeadBucketFunc func(ctx context.Context, input *s3.HeadBucketInput, opts ...func(*s3.Options)) (*s3.HeadBucketOutput, error)
	HeadObjectFunc func(ctx context.Context, input *s3.HeadObjectInput, opts ...func(*s3.Options)) (*s3.HeadObjectOutput, error)
	PutObjectFunc  func(ctx context.Context, input *s3.PutObjectInput, opts ...func(*s3.Options)) (*s3.PutObjectOutput, error)

	// AbortMultipartUploadFunc    func(ctx context.Context, input *s3.AbortMultipartUploadInput, opts ...func(*s3.Options)) (*s3.AbortMultipartUploadOutput, error)
	// CompleteMultipartUploadFunc func(ctx context.Context, input *s3.CompleteMultipartUploadInput, opts ...func(*s3.Options)) (*s3.CompleteMultipartUploadOutput, error)
	// CreateMultipartUploadFunc   func(ctx context.Context, input *s3.CreateMultipartUploadInput, opts ...func(*s3.Options)) (*s3.CreateMultipartUploadOutput, error)
	// PutObjectFunc               func(ctx context.Context, input *s3.PutObjectInput, opts ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	// UploadPartFunc              func(ctx context.Context, input *s3.UploadPartInput, opts ...func(*s3.Options)) (*s3.UploadPartOutput, error)
}

// Mock do paginator para testes
type MockPaginator struct {
	Pages []*s3.ListObjectsV2Output
	Index int
	Err   error
}

type MockObjectNotExists struct {
	WaitFunc          func(ctx context.Context, params *s3.HeadObjectInput, maxWaitDur time.Duration, optFns ...func(*s3.ObjectNotExistsWaiterOptions)) error
	WaitForOutputFunc func(ctx context.Context, params *s3.HeadObjectInput, maxWaitDur time.Duration, optFns ...func(*s3.ObjectNotExistsWaiterOptions)) (*s3.HeadObjectOutput, error)
}

type MockBucketNotExists struct {
	WaitFunc          func(ctx context.Context, params *s3.HeadBucketInput, maxWaitDur time.Duration, optFns ...func(*s3.BucketNotExistsWaiterOptions)) error
	WaitForOutputFunc func(ctx context.Context, params *s3.HeadBucketInput, maxWaitDur time.Duration, optFns ...func(*s3.BucketNotExistsWaiterOptions)) (*s3.HeadBucketOutput, error)
}

var NotFound *types.NotFound

// Implementação do CreateBucket do Mock
func (m *MockS3Client) CreateBucket(ctx context.Context, input *s3.CreateBucketInput, opts ...func(*s3.Options)) (*s3.CreateBucketOutput, error) {
	if m.CreateBucketFunc != nil {
		return m.CreateBucketFunc(ctx, input, opts...)
	}
	return nil, nil
}

// Implementação do DeleteBucket do Mock
func (m *MockS3Client) DeleteBucket(ctx context.Context, input *s3.DeleteBucketInput, opts ...func(*s3.Options)) (*s3.DeleteBucketOutput, error) {
	if m.DeleteBucketFunc != nil {
		return m.DeleteBucketFunc(ctx, input, opts...)
	}
	return nil, nil
}

// Implementação do DeleteObjects do Mock
func (m *MockS3Client) DeleteObjects(ctx context.Context, input *s3.DeleteObjectsInput, opts ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error) {
	if m.DeleteObjectsFunc != nil {
		return m.DeleteObjectsFunc(ctx, input, opts...)
	}
	return nil, nil
}

// Implementação do ListObjectsV2 do Mock
func (m *MockS3Client) ListObjectsV2(ctx context.Context, input *s3.ListObjectsV2Input, opts ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	if m.ListObjectsV2Func != nil {
		return m.ListObjectsV2Func(ctx, input, opts...)
	}
	return nil, nil
}

// Implementação do método HasMorePages do Mock
func (m *MockPaginator) HasMorePages() bool {
	return m.Index < len(m.Pages)
}

// Implementação do método NextPage do Mock
func (m *MockPaginator) NextPage(ctx context.Context, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	if !m.HasMorePages() {
		return nil, m.Err
	}
	page := m.Pages[m.Index]
	m.Index++
	return page, nil
}

// Implementação do método Wait do Mock: MockObjectNotExists
func (m *MockObjectNotExists) Wait(ctx context.Context, params *s3.HeadObjectInput, maxWaitDur time.Duration, optFns ...func(*s3.ObjectNotExistsWaiterOptions)) error {
	if m.WaitFunc != nil {
		return m.WaitFunc(ctx, params, maxWaitDur, optFns...)
	}
	return nil
}

// Implementação do método WaitForOutput do Mock: MockObjectNotExists
func (m *MockObjectNotExists) WaitForOutput(ctx context.Context, params *s3.HeadObjectInput, maxWaitDur time.Duration, optFns ...func(*s3.ObjectNotExistsWaiterOptions)) (*s3.HeadObjectOutput, error) {
	return nil, nil
}

// Implementação do método Wait do Mock: MockBucketNotExists
func (m *MockBucketNotExists) Wait(ctx context.Context, params *s3.HeadBucketInput, maxWaitDur time.Duration, optFns ...func(*s3.BucketNotExistsWaiterOptions)) error {
	if m.WaitFunc != nil {
		return m.WaitFunc(ctx, params, maxWaitDur, optFns...)
	}
	return nil
}

// Implementação do método WaitForOutput do Mock: MockBucketNotExists
func (m *MockBucketNotExists) WaitForOutput(ctx context.Context, params *s3.HeadBucketInput, maxWaitDur time.Duration, optFns ...func(*s3.BucketNotExistsWaiterOptions)) (*s3.HeadBucketOutput, error) {
	return nil, nil
}

// Implementação do método HeadBucket do Mock
func (m *MockS3Client) HeadBucket(ctx context.Context, input *s3.HeadBucketInput, opts ...func(*s3.Options)) (*s3.HeadBucketOutput, error) {
	if m.HeadBucketFunc != nil {
		return m.HeadBucketFunc(ctx, input, opts...)
	}
	return nil, nil
}

// Implementação do método HeadObject do Mock
func (m *MockS3Client) HeadObject(ctx context.Context, input *s3.HeadObjectInput, opts ...func(*s3.Options)) (*s3.HeadObjectOutput, error) {
	if m.HeadObjectFunc != nil {
		return m.HeadObjectFunc(ctx, input, opts...)
	}
	return nil, nil
}

// Implementação do método PutObject do Mock
func (m *MockS3Client) PutObject(ctx context.Context, input *s3.PutObjectInput, opts ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	if m.PutObjectFunc != nil {
		return m.PutObjectFunc(ctx, input, opts...)
	}
	return nil, nil
}

// func (m *MockS3Client) AbortMultipartUpload(ctx context.Context, input *s3.AbortMultipartUploadInput, opts ...func(*s3.Options)) (*s3.AbortMultipartUploadOutput, error) {
// 	if m.AbortMultipartUploadFunc != nil {
// 		return m.AbortMultipartUploadFunc(ctx, input, opts...)
// 	}

// 	return nil, nil
// }

// func (m *MockS3Client) CompleteMultipartUpload(ctx context.Context, input *s3.CompleteMultipartUploadInput, opts ...func(*s3.Options)) (*s3.CompleteMultipartUploadOutput, error) {
// 	if m.CompleteMultipartUploadFunc != nil {
// 		return m.CompleteMultipartUploadFunc(ctx, input, opts...)
// 	}

// 	return nil, nil
// }

// func (m *MockS3Client) CreateMultipartUpload(ctx context.Context, input *s3.CreateMultipartUploadInput, opts ...func(*s3.Options)) (*s3.CreateMultipartUploadOutput, error) {
// 	if m.CreateMultipartUploadFunc != nil {
// 		return m.CreateMultipartUploadFunc(ctx, input, opts...)
// 	}

// 	return nil, nil
// }

// func (m *MockS3Client) PutObject(ctx context.Context, input *s3.PutObjectInput, opts ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
// 	if m.PutObjectFunc != nil {
// 		return m.PutObjectFunc(ctx, input, opts...)
// 	}

// 	return nil, nil
// }

// func (m *MockS3Client) UploadPart(ctx context.Context, input *s3.UploadPartInput, opts ...func(*s3.Options)) (*s3.UploadPartOutput, error) {
// 	if m.UploadPartFunc != nil {
// 		return m.UploadPartFunc(ctx, input, opts...)
// 	}

// 	return nil, nil
// }

// Mock da inicialização do S3
func NewS3ClientMock(mock S3Api) *Client {

	log.Println("Carregando interface mock")

	return &Client{
		s3Client: mock,
		paginator: func(input *s3.ListObjectsV2Input) S3Paginator {
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
		objNotExistWaiter: func() S3NewObjectNotExists {
			return &MockObjectNotExists{
				WaitFunc: func(ctx context.Context, params *s3.HeadObjectInput, maxWaitDur time.Duration, optFns ...func(*s3.ObjectNotExistsWaiterOptions)) error {
					errorBucket := "bucket-timeout"
					pntrErrorBucket := &errorBucket
					if *params.Bucket == *pntrErrorBucket {
						return ErrWaiterTimeout
					}
					return nil
				},
			}
		},
		bucketNotExistsWaiter: func() S3NewBucketNotExists {
			return &MockBucketNotExists{
				WaitFunc: func(ctx context.Context, params *s3.HeadBucketInput, maxWaitDur time.Duration, optFns ...func(*s3.BucketNotExistsWaiterOptions)) error {
					errorBucket := "bucket-timeout"
					pntrErrorBucket := &errorBucket
					if *params.Bucket == *pntrErrorBucket {
						return ErrWaiterTimeout
					}
					return nil
				},
			}
		},
	}

}

// Cria o mock
func CreateS3ClientMock() *Client {
	mock := &MockS3Client{
		CreateBucketFunc: func(ctx context.Context, input *s3.CreateBucketInput, opts ...func(*s3.Options)) (*s3.CreateBucketOutput, error) {
			if *input.Bucket == "bucket-exists" {
				return nil, Exists
			} else if *input.Bucket == "bucket-owned" {
				return nil, Owned
			}
			return &s3.CreateBucketOutput{}, nil
		},
		DeleteBucketFunc: func(ctx context.Context, input *s3.DeleteBucketInput, opts ...func(*s3.Options)) (*s3.DeleteBucketOutput, error) {
			if *input.Bucket == "no-bucket" {
				return nil, NoBucket
			}
			return &s3.DeleteBucketOutput{}, nil
		},
		DeleteObjectsFunc: func(ctx context.Context, input *s3.DeleteObjectsInput, opts ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error) {
			if *input.Bucket == "no-bucket" {
				return nil, NoBucket
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
				return nil, NotFound
			}
			return &s3.HeadBucketOutput{}, nil
		},
	}

	return NewS3ClientMock(mock)
}
