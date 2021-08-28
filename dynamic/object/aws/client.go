package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type ObjectClient struct {
	bucket string
	client *s3.S3
}

func NewObjectClient(access_key, secret_key, region, bucket string) (*ObjectClient, error) {
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(access_key, secret_key, ""),
		Region:      aws.String(region),
	})
	if err != nil {
		return nil, fmt.Errorf("error creating new aws session: %w", err)
	}

	return &ObjectClient{
		bucket: bucket,
		client: s3.New(sess),
	}, nil

}
