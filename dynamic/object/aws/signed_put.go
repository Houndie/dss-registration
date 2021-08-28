package aws

import (
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/object"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (c *ObjectClient) SignedPut(filesize int64, filename string) (string, error) {
	req, _ := c.client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(filename),
	})

	req.HTTPRequest.Header.Set("Content-Length", fmt.Sprintf("%v", filesize))

	url, err := req.Presign(object.PutSigningDuration)
	if err != nil {
		return "", fmt.Errorf("error presigning url: %w", err)
	}

	return url, nil
}
