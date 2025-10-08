package s3handler_test

import (
	"github.com/ADStefano/AmazonHandler/s3handler"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"
)

type testListObjects struct {
	testName       string
	bucketName     string
	prefix         string
	expectedOutput []types.Object
	expectedError  error
	maxKeys        int32
}

var output = []types.Object{
	{Key: aws.String("exemplo.html"), Size: aws.Int64(2048), LastModified: aws.Time(time.Now().Add(-24 * time.Hour))},
}

var testListFile = []testListObjects{
	{testName: "Teste - Sucesso", bucketName: "test", prefix: "", maxKeys: 5, expectedOutput: output, expectedError: nil},
	{testName: "Teste - Erro: No Such Bucket", bucketName: "no-such-bucket", prefix: "", maxKeys: 5, expectedOutput: output, expectedError: s3handler.ErrNoSuchBucket},
}

type listBuckets struct {
	testName       string
	bucketName     string
	prefix         string
	expectedOutput []types.Bucket
	expectedError  error
}

var outputBuckets = []types.Bucket{
	{Name: aws.String("bucket-teste"), CreationDate: aws.Time(time.Now().Add(-240 * time.Hour))},
}

var testListBuckets = []listBuckets{
	{testName: "Teste - Sucesso", bucketName: "test", prefix: "", expectedOutput: outputBuckets, expectedError: nil},
	{testName: "Teste - Erro: Access Denied", bucketName: "access-denied", prefix: "access-denied", expectedOutput: outputBuckets, expectedError: s3handler.ErrAccessDenied},
}

// Teste para ListObjects
func TestListObjects(t *testing.T) {

	for _, testCase := range testListFile {
		t.Run(testCase.bucketName, func(t *testing.T) {

			t.Logf("Testando bucket: %s, maxKeys: %d", testCase.bucketName, testCase.maxKeys)

			output, err := mockClient.ListObjects(testCase.bucketName, testCase.prefix, testCase.maxKeys, context.Background())

			for _, outputItem := range testCase.expectedOutput {
				for _, item := range output {

					if *item.Key != *outputItem.Key {
						t.Errorf("Esperado key: %s, Recebido: %v\n", *item.Key, *outputItem.Key)
					}
					if *item.Size != *outputItem.Size {
						t.Errorf("Esperado size: %d, Recebido: %d\n", *item.Size, *outputItem.Size)
					}
				}
			}

			if !errors.Is(err, testCase.expectedError) {
				t.Errorf("Esperado error: %v , Recebido: %v\n", testCase.expectedError, err)
			}
		})
	}
}

func TestListBuckets(t *testing.T) {

	for _, testCase := range testListBuckets {

		t.Run("ListBuckets", func(t *testing.T) {

			t.Logf("%s - Testando bucket: %s", testCase.testName, testCase.bucketName)

			output, err := mockClient.ListBuckets(testCase.prefix, context.Background())

			for _, outputItem := range testCase.expectedOutput {
				for _, item := range output {

					if *item.Name != *outputItem.Name {
						t.Errorf("Esperado name: %s, Recebido: %v\n", *item.Name, *outputItem.Name)
					}
				}
			}

			if !errors.Is(err, testCase.expectedError) {
				t.Errorf("Esperado error: %v , Recebido: %v\n", testCase.expectedError, err)
			}
		})

	}
}
