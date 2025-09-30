package s3handler

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
)

var (
	ErrApi            smithy.APIError
	ErrExists         *types.BucketAlreadyExists
	ErrOwned          *types.BucketAlreadyOwnedByYou
	ErrNoSuchBucket   *types.NoSuchBucket
	ErrNotFound       *types.NotFound
	ErrNoSuchKey      *types.NoSuchKey
	ErrEntityTooLarge = errors.New("EntityTooLarge")
	ErrAccessDenied   = errors.New("AccessDenied")
	ErrWaiterTimeout  = errors.New("exceeded max wait time for BucketNotExists waiter")
	ErrEmptyParam   = errors.New("bucket name or object key cannot be empty")
)
