package s3handler_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/ADStefano/AmazonHandler/s3handler"
)

type TestUpload struct {
	TestName       string
	TestBucketName string
	Prefix         string
	Filename       string
	File           *os.File
	ExpectedOutput bool
	ExpectedError  error
}

var mockFile *os.File

var testUploadCases = []TestUpload{
	{
		TestName:       "Success - upload without prefix",
		TestBucketName: "test",
		Prefix:         "",
		Filename:       "exemplo.html",
		File:           mockFile,
		ExpectedOutput: true,
		ExpectedError:  nil,
	},
	{
		TestName:       "Success",
		TestBucketName: "test",
		Prefix:         "uploads",
		Filename:       "exemplo.html",
		File:           mockFile,
		ExpectedOutput: true,
		ExpectedError:  nil,
	},
	{
		TestName:       "Fail - Entity Too Large",
		TestBucketName: "entity-too-large",
		Prefix:         "uploads",
		Filename:       "exemplo.html",
		File:           mockFile,
		ExpectedOutput: false,
		ExpectedError:  s3handler.ErrEntityTooLarge,
	},
}

func TestUploads(t *testing.T) {

	for _, testCase := range testUploadCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			output, err := mockClient.UploadS3(testCase.TestBucketName, testCase.Prefix, testCase.Filename, testCase.File, context.Background())

			if output != testCase.ExpectedOutput {
				t.Errorf("Output esperado %v, recebido %v", testCase.ExpectedOutput, output)
			}

			if !errors.Is(err, testCase.ExpectedError) {
				t.Errorf("Erro esperado %v, recebido %v", testCase.ExpectedError, err)
			}
		})
	}
}
