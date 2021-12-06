package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (o *ObjectClient) Delete(ctx context.Context, filename string) error {
	_, err := o.client.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(o.bucket),
		Key:    aws.String(filename),
	})

	if err != nil {
		return fmt.Errorf("error deleting object: %w", err)
	}

	return nil
}
