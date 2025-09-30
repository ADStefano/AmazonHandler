package s3handler

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Estrutura do client do S3
type Client struct {
	S3Client              S3Api
	ObjectPaginator       CreateObjectPaginator
	BucketPaginator	      CreateBucketPaginator
	ObjNotExistWaiter     CreateNewObjectNotExists
	BucketNotExistsWaiter CreateNewBucketNotExists
	PresignerClient       Presigner
}

// NewS3Client inicializa um client S3 com o client do Presign
func NewS3Client(client *s3.Client) *Client {

	return &Client{
		S3Client: client,
		PresignerClient: s3.NewPresignClient(client),
		BucketPaginator: func(input *s3.ListBucketsInput) S3BucketPaginator {
			return s3.NewListBucketsPaginator(client, input)
		},
		ObjectPaginator: func(input *s3.ListObjectsV2Input) S3ObjectPaginator {
			return s3.NewListObjectsV2Paginator(client, input)
		},
		ObjNotExistWaiter: func() S3NewObjectNotExists {
			return s3.NewObjectNotExistsWaiter(client)
		},
		BucketNotExistsWaiter: func() S3NewBucketNotExists {
			return s3.NewBucketNotExistsWaiter(client)
		},
	}

}
