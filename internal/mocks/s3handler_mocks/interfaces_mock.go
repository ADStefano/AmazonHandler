package mock

import (
	"context"
	"time"
	"amazon-handler/s3handler"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Mock do S3 para testes
type MockS3Client struct {
	CreateBucketFunc func(ctx context.Context, input *s3.CreateBucketInput, opts ...func(*s3.Options)) (*s3.CreateBucketOutput, error)

	DeleteBucketFunc  func(ctx context.Context, input *s3.DeleteBucketInput, opts ...func(*s3.Options)) (*s3.DeleteBucketOutput, error)
	DeleteObjectsFunc func(ctx context.Context, input *s3.DeleteObjectsInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error)

	ListBucketsFunc   func(ctx context.Context, input *s3.ListBucketsInput, opts ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
	ListObjectsV2Func func(ctx context.Context, input *s3.ListObjectsV2Input, opts ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)

	HeadBucketFunc func(ctx context.Context, input *s3.HeadBucketInput, opts ...func(*s3.Options)) (*s3.HeadBucketOutput, error)
	HeadObjectFunc func(ctx context.Context, input *s3.HeadObjectInput, opts ...func(*s3.Options)) (*s3.HeadObjectOutput, error)
	PutObjectFunc  func(ctx context.Context, input *s3.PutObjectInput, opts ...func(*s3.Options)) (*s3.PutObjectOutput, error)

	GetObjectFunc func(ctx context.Context, input *s3.GetObjectInput, opts ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

var _ s3handler.S3Api = (*MockS3Client)(nil)

// Mock do paginator para testes
type MockPaginator struct {
	Pages []*s3.ListObjectsV2Output
	Index int
	Err   error
}

// Mock para o waiter de objeto não existente
type MockObjectNotExists struct {
	WaitFunc          func(ctx context.Context, params *s3.HeadObjectInput, maxWaitDur time.Duration, optFns ...func(*s3.ObjectNotExistsWaiterOptions)) error
	WaitForOutputFunc func(ctx context.Context, params *s3.HeadObjectInput, maxWaitDur time.Duration, optFns ...func(*s3.ObjectNotExistsWaiterOptions)) (*s3.HeadObjectOutput, error)
}

// Mock para o waiter de bucket não existente
type MockBucketNotExists struct {
	WaitFunc          func(ctx context.Context, params *s3.HeadBucketInput, maxWaitDur time.Duration, optFns ...func(*s3.BucketNotExistsWaiterOptions)) error
	WaitForOutputFunc func(ctx context.Context, params *s3.HeadBucketInput, maxWaitDur time.Duration, optFns ...func(*s3.BucketNotExistsWaiterOptions)) (*s3.HeadBucketOutput, error)
}

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

// Implementação do ListBuckets do Mock
func (m *MockS3Client) ListBuckets(ctx context.Context, input *s3.ListBucketsInput, opts ...func(*s3.Options)) (*s3.ListBucketsOutput, error) {
	if m.ListBucketsFunc != nil {
		return m.ListBucketsFunc(ctx, input, opts...)
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

func (m *MockS3Client) GetObject(ctx context.Context, input *s3.GetObjectInput, opts ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	if m.GetObjectFunc != nil {
		return m.GetObjectFunc(ctx, input, opts...)
	}
	return nil, nil
}
