package s3handler_test

import (
	"amazon-handler/s3handler"
	"amazon-handler/internal/mocks/s3handler_mocks"
	"errors"
	"testing"
)

type testBucket struct {
	testBucketName string
	expectedOutput bool
	expectedError  error
}

var mockClient = mock.CreateS3ClientMock()

var testCreateBuckets = []testBucket{
	{testBucketName: "test", expectedOutput: true, expectedError: nil},
	{testBucketName: "bucket-exists", expectedOutput: false, expectedError: s3handler.ErrExists},
	{testBucketName: "bucket-owned", expectedOutput: false, expectedError: s3handler.ErrOwned},
}

var TestDeleteBuckets = []testBucket{
	{testBucketName: "test", expectedOutput: true, expectedError: nil},
	{testBucketName: "no-bucket", expectedOutput: false, expectedError: s3handler.ErrNoSuchBucket},
	{testBucketName: "bucket-timeout", expectedOutput: false, expectedError: s3handler.ErrWaiterTimeout},
}

// Teste para CreateBucket
func TestCreateBucket(t *testing.T) {

	for _, testCase := range testCreateBuckets {
		t.Run(testCase.testBucketName, func(t *testing.T) {

			output, err := mockClient.CreateBucket(testCase.testBucketName)
			if output != testCase.expectedOutput {
				t.Errorf("Esperado: %v , Recebido: %v\n", testCase.expectedOutput, output)
			}

			if !errors.Is(err, testCase.expectedError) {
				t.Errorf("Esperado: %v , Recebido: %v\n", testCase.expectedError, err)
			}
		})
	}

}

// Teste para DeleteBucket
func TestDeleteBucket(t *testing.T) {
	for _, testCase := range TestDeleteBuckets {

		t.Run(testCase.testBucketName, func(t *testing.T) {

			t.Logf("Testando bucket: %s", testCase.testBucketName)

			output, err := mockClient.DeleteBucket(testCase.testBucketName)

			if output != testCase.expectedOutput {
				t.Errorf("Esperado: %v , Recebido: %v\n", testCase.expectedOutput, output)
			}

			if !errors.Is(err, testCase.expectedError) {
				t.Errorf("Esperado: %v , Recebido: %v\n", testCase.expectedError, err)
			}
		})
	}
}
