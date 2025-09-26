package s3handler_test

import (
	"amazon-handler/s3handler"
	"errors"
	"testing"
	"time"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

type TestPresignedURLStruct struct {
	TestName       string
	TestBucketName string
	TestObjectKey  string
	TestExpiration time.Duration
	ExpectedOutput *v4.PresignedHTTPRequest
	ExpectedError  error
}

var testPresignedURL = []TestPresignedURLStruct{
	{
		TestName:       "GET Presigned URL - Sucesso",
		TestBucketName: "test-bucket-success",
		TestObjectKey:  "test-object-success",
		TestExpiration: 300 * time.Second,
		ExpectedOutput: &v4.PresignedHTTPRequest{
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
		},
		ExpectedError: nil,
	},
	{
		TestName:       "GET Presigned URL - Sem Bucket",
		TestBucketName: "",
		TestObjectKey:  "test-object",
		TestExpiration: 300 * time.Second,
		ExpectedOutput: nil,
		ExpectedError:  s3handler.ErrEmptyParam,
	},
}

func TestGetPresignURL(t *testing.T) {

	for _, testCase := range testPresignedURL {
		t.Run(testCase.TestName, func(t *testing.T) {

			t.Logf("Testando bucket: %s, objectKey: %s, expiration: %s", testCase.TestBucketName, testCase.TestObjectKey, testCase.TestExpiration)

			output, err := mockClient.GetPreSignedURL(testCase.TestBucketName, testCase.TestObjectKey, testCase.TestExpiration)

			if testCase.ExpectedOutput != nil && output != nil {
				if testCase.ExpectedOutput.URL != output.URL {
					t.Errorf("Esperado Output: %v , Recebido: %v\n", testCase.ExpectedOutput, output)
				}
			}

			if !errors.Is(err, testCase.ExpectedError) {
				t.Errorf("Esperado error: %v , Recebido: %v\n", testCase.ExpectedError, err)
			}
		})
	}
}

func TestPutPresignURL(t *testing.T) {

	for _, testCase := range testPresignedURL {
		t.Run(testCase.TestName, func(t *testing.T) {

			t.Logf("Testando bucket: %s, objectKey: %s, expiration: %s", testCase.TestBucketName, testCase.TestObjectKey, testCase.TestExpiration)

			output, err := mockClient.PutPreSignedURL(testCase.TestBucketName, testCase.TestObjectKey, testCase.TestExpiration)

			if testCase.ExpectedOutput != nil && output != nil {
				if testCase.ExpectedOutput.URL != output.URL {
					t.Errorf("Esperado Output: %v , Recebido: %v\n", testCase.ExpectedOutput, output)
				}
			}

			if !errors.Is(err, testCase.ExpectedError) {
				t.Errorf("Esperado error: %v , Recebido: %v\n", testCase.ExpectedError, err)
			}
		})
	}
}

func TestDeleteObjectPresignURL(t *testing.T) {

	for _, testCase := range testPresignedURL {
		t.Run(testCase.TestName, func(t *testing.T) {

			t.Logf("Testando bucket: %s, objectKey: %s, expiration: %s", testCase.TestBucketName, testCase.TestObjectKey, testCase.TestExpiration)

			output, err := mockClient.DeleteObjectPreSignedURL(testCase.TestBucketName, testCase.TestObjectKey, testCase.TestExpiration)

			if testCase.ExpectedOutput != nil && output != nil {
				if testCase.ExpectedOutput.URL != output.URL {
					t.Errorf("Esperado Output: %v , Recebido: %v\n", testCase.ExpectedOutput, output)
				}
			}

			if !errors.Is(err, testCase.ExpectedError) {
				t.Errorf("Esperado error: %v , Recebido: %v\n", testCase.ExpectedError, err)
			}
		})
	}
}

func TestDeleteBucketPresignURL(t *testing.T) {

	for _, testCase := range testPresignedURL {
		t.Run(testCase.TestName, func(t *testing.T) {

			t.Logf("Testando bucket: %s, objectKey: %s, expiration: %s", testCase.TestBucketName, testCase.TestObjectKey, testCase.TestExpiration)

			output, err := mockClient.DeleteBucketPreSignedURL(testCase.TestBucketName, testCase.TestExpiration)

			if testCase.ExpectedOutput != nil && output != nil {
				if testCase.ExpectedOutput.URL != output.URL {
					t.Errorf("Esperado Output: %v , Recebido: %v\n", testCase.ExpectedOutput, output)
				}
			}

			if !errors.Is(err, testCase.ExpectedError) {
				t.Errorf("Esperado error: %v , Recebido: %v\n", testCase.ExpectedError, err)
			}
		})
	}
}