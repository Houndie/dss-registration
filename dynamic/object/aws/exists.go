package aws

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (c *ObjectClient) Exists(ctx context.Context, filename string) (bool, error) {
	_, err := c.client.HeadObjectWithContext(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(filename),
	})

	if err != nil {
		var aerr awserr.RequestFailure
		if errors.As(err, &aerr) && (aerr.StatusCode() == 403 || aerr.StatusCode() == 404) {
			return false, nil
		}
		return false, fmt.Errorf("error fetching object metadata: %w", err)
	}

	return true, nil
}
