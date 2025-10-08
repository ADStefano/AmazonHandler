package s3handler_test

import (
	"github.com/ADStefano/AmazonHandler/s3handler"
	"context"
	"errors"
	"testing"
)

type testDeleteFile struct {
	testBucketName string
	filesNames     []string
	expectedOutput bool
	expectedError  error
}

var testDeleteObjects = []testDeleteFile{
	{testBucketName: "", filesNames: []string{""}, expectedOutput: false, expectedError: s3handler.ErrEmptyParam},
	{testBucketName: "test", filesNames: []string{"teste", "teste1", "teste2"}, expectedOutput: true, expectedError: nil},
	{testBucketName: "no-bucket", filesNames: []string{"teste", "teste1", "teste2"}, expectedOutput: false, expectedError: s3handler.ErrNoSuchBucket},
	{testBucketName: "bucket-timeout", filesNames: []string{"teste", "teste1", "teste2"}, expectedOutput: false, expectedError: s3handler.ErrWaiterTimeout},
}

// Teste para DeleteFiles
func TestDeleteFiles(t *testing.T) {
	for _, testCase := range testDeleteObjects {

		output, err := mockClient.DeleteObjects(testCase.filesNames, testCase.testBucketName, context.Background())

		t.Logf("Testando bucket: %s, arquivos: %v", testCase.testBucketName, testCase.filesNames)

		if output != testCase.expectedOutput {
			t.Errorf("Output diferente do esperado\nNome: %s, Esperado: %v , Recebido: %v\n", testCase.testBucketName, testCase.expectedOutput, output)
		}

		if !errors.Is(err, testCase.expectedError) {
			t.Errorf("Erro diferente do esperado\nNome: %s, Esperado: %v , Recebido: %v\n", testCase.testBucketName, testCase.expectedError, err)
		}
	}
}

func TestEmptyBucket(t *testing.T) {
	for _, testCase := range testDeleteObjects {

		t.Logf("Testando bucket: %s", testCase.testBucketName)

		output, err := mockClient.EmptyBucket(testCase.testBucketName, context.Background())

		if output != testCase.expectedOutput {
			t.Errorf("Output diferente do esperado. Nome: %s, Esperado: %v , Recebido: %v\n", testCase.testBucketName, testCase.expectedOutput, output)
		}

		if !errors.Is(err, testCase.expectedError) {
			t.Errorf("Erro diferente do esperado. Nome: %s, Esperado: %v , Recebido: %v\n", testCase.testBucketName, testCase.expectedError, err)
		}
	}
}
