package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// NewS3ClientOptions is the options for NewS3Client.
type NewS3ClientOptions struct {
	EndpointURL string
	Region      string

	UseAnonymousCredentials bool

	AssumeRole *AssumeRole

	InsecureSkipVerify bool
}

// NewS3Client returns a new S3 client.
func NewS3Client(ctx context.Context, o NewS3ClientOptions) (*s3.Client, error) {
	conf, err := NewConfig(ctx, NewConfigOptions{
		Region:                  o.Region,
		UseAnonymousCredentials: o.UseAnonymousCredentials,
		AssumeRole:              o.AssumeRole,
		InsecureSkipVerify:      o.InsecureSkipVerify,
	})
	if err != nil {
		return nil, err
	}
	return s3.NewFromConfig(conf, func(opt *s3.Options) {
		if o.EndpointURL != "" {
			opt.BaseEndpoint = aws.String(o.EndpointURL)
		}

		opt.Region = o.Region
		// This is needed as the minio server does not support the virtual host style.
		opt.UsePathStyle = true
	}), nil
}
