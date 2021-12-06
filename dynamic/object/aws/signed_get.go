package aws

import (
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/object"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (c *ObjectClient) SignedGet(filename string) (string, error) {
	req, _ := c.client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(filename),
	})

	url, err := req.Presign(object.GetSigningDuration)
	if err != nil {
		return "", fmt.Errorf("error presigning url: %w", err)
	}

	return url, nil
}
