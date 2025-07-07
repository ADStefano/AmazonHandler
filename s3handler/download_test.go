package s3handler_test

import (
	"amazon-handler/s3handler"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type TestDownload struct {
	TestBucketName string
	TestObjectKey  string
	ExpectedOutput *s3.GetObjectOutput
	ExpectedError  error
}

var mockGetObjectOutput = &s3.GetObjectOutput{}

var testDownloadCases = []TestDownload{
	{
		TestBucketName: "test",
		TestObjectKey:  "uploads/exemplo.html",
		ExpectedOutput: mockGetObjectOutput,
		ExpectedError:  nil,
	},
	{
		TestBucketName: "no-such-key",
		TestObjectKey:  "uploads/exemplo.html",
		ExpectedOutput: nil,
		ExpectedError:  s3handler.ErrNoSuchKey,
	},
	{
		TestBucketName: "no-such-bucket",
		TestObjectKey:  "uploads/exemplo.html",
		ExpectedOutput: nil,
		ExpectedError:  s3handler.ErrNoSuchBucket,
	},
}

func TestDownloadFunc(t *testing.T) {
	for _, testCase := range testDownloadCases {
		t.Run(testCase.TestBucketName, func(t *testing.T) {

			t.Logf("Testando Download do bucket: %s, objeto: %s", testCase.TestBucketName, testCase.TestObjectKey)

			output, err := mockClient.DownloadS3(testCase.TestBucketName, testCase.TestObjectKey)

			if testCase.ExpectedOutput != nil {
				if output == nil {
					t.Errorf("Esperado output não nulo, recebido: %v", output)
				}
			} else {
				if output != nil {
					t.Errorf("Esperado output não nulo, recebido: %v", output)
				}
			}

			if !errors.Is(err, testCase.ExpectedError) {
				t.Errorf("Erro esperado: %v , Recebido: %v\n", testCase.ExpectedError, err)
			}
		})
	}
}
