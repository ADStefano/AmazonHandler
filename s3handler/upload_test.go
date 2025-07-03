package s3handler_test

import (
	"amazon-handler/s3handler"
	"errors"
	"testing"
)

type TestUpload struct {
	TestBucketName string
	Prefix         string
	Path           string
	ExpectedOutput bool
	ExpectedError  error
}

var testUploadCases = []TestUpload{
	{
		TestBucketName: "test",
		Prefix:         "",
		Path:           "/home/angelo/Documentos/Programação/exemplo.html",
		ExpectedOutput: true,
		ExpectedError:  nil,
	},
	{
		TestBucketName: "test",
		Prefix:         "uploads",
		Path:           "/home/angelo/Documentos/Programação/exemplo.html",
		ExpectedOutput: true,
		ExpectedError:  nil,
	},
	{
		TestBucketName: "entity-too-large",
		Prefix:         "uploads",
		Path:           "/home/angelo/Documentos/Programação/exemplo.html",
		ExpectedOutput: false,
		ExpectedError:  s3handler.ErrEntityTooLarge,
	},
}

func TestUploads(t *testing.T) {
	for _, testCase := range testUploadCases {
		t.Run(testCase.Path, func(t *testing.T) {
			output, err := mockClient.Upload(testCase.TestBucketName, testCase.Prefix, testCase.Path)

			if output != testCase.ExpectedOutput {
				t.Errorf("Output esperado %v, recebido %v", testCase.ExpectedOutput, output)
			}

			if !errors.Is(err, testCase.ExpectedError) {
				t.Errorf("Erro esperado %v, recebido %v", testCase.ExpectedError, err)
			}
		})
	}
}
