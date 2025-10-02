package s3handler

import (
	"errors"
	"fmt"

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
	ErrEmptyParam     = errors.New("bucket name or object key cannot be empty")
)

type S3Error struct {
	Operation string
	Bucket    string
	Object    string
	Message   string
	Err       error
}

func (e *S3Error) Error() string {
	return fmt.Sprintf("s3 %s error [bucket=%s object=%s code=%s]: %v", e.Operation, e.Bucket, e.Object, e.Message, e.Err)
}

func (e *S3Error) Unwrap() error {
    return e.Err
}

func ParseError(err error) error {

	var errApi smithy.APIError

	if errors.As(err, &ErrNoSuchKey) {
		return ErrNoSuchKey
	} else if errors.As(err, &errApi) && errApi.ErrorCode() == "NoSuchBucket" {
		return ErrNoSuchBucket
	} else if errors.As(err, &errApi) && errApi.ErrorCode() == "AccessDenied"{
		return ErrAccessDenied
	} else if errors.As(err, &errApi) && errApi.ErrorCode() == "EntityTooLarge" {
		return ErrEntityTooLarge
	} else if errors.As(err, &ErrOwned){
		return ErrOwned
	} else if errors.As(err, &ErrExists){
		return ErrExists
	} else {
		return err
	}
}