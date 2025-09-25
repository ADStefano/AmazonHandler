package s3handler_test

import v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"

type TestPresignedURL struct {
	TestBucketName string
	TestObjectKey  string
	ExpectedOutput *v4.PresignedHTTPRequest
	ExpectedError  error
}

var testPresignedURL = []TestPresignedURL{
	{
		TestBucketName: "test-bucket",
		TestObjectKey:  "test-object",
		ExpectedOutput: &v4.PresignedHTTPRequest{
			URL: "https://httpbin.org/status/500",
			Method: "GET",
			SignedHeader: map[string][]string{
				"Host": {"test-bucket.s3.amazonaws.com"},
				"X-Amz-Content-Sha256": {"UNSIGNED-PAYLOAD"},
				"X-Amz-Date": {"20240605T123456Z"},
				"X-Amz-Expires": {"300"},
				"X-Amz-SignedHeaders": {"host"},
				"X-Amz-Signature": {"EXAMPLE_SIGNATURE"},
			},
		},
		ExpectedError: nil,
	},
}