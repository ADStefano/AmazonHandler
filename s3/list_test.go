package s3_test

import (
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"
)

type testListObjects struct {
	bucketName     string
	expectedOutput []types.Object
	expectedError  error
	maxKeys        int32
}

var output = []types.Object{
	{Key: aws.String("exemplo.html"), Size: aws.Int64(2048), LastModified: aws.Time(time.Now().Add(-24 * time.Hour))},
}

var testListFile = []testListObjects{
	{bucketName: "test", maxKeys: 5, expectedOutput: output, expectedError: nil},
}

// Teste para ListObjects
func TestListObjects(t *testing.T) {

	for _, testCase := range testListFile {
		t.Run(testCase.bucketName, func(t *testing.T) {

			t.Logf("Testando bucket: %s, maxKeys: %d", testCase.bucketName, testCase.maxKeys)

			output, err := mockClient.ListObjects(testCase.bucketName, testCase.maxKeys)

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
