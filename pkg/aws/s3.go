package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// AssumeRole is the options for AssumeRole.
type AssumeRole struct {
	RoleARN    string
	ExternalID string
}

// NewS3ClientOptions is the options for NewS3Client.
type NewS3ClientOptions struct {
	EndpointURL string
	Region      string

	UseAnonymousCredentials bool

	AssumeRole *AssumeRole
}

// NewS3Client returns a new S3 client.
func NewS3Client(ctx context.Context, o NewS3ClientOptions) (*s3.Client, error) {
	var err error
	var conf aws.Config
	if o.UseAnonymousCredentials {
		conf, err = config.LoadDefaultConfig(ctx,
			config.WithCredentialsProvider(aws.AnonymousCredentials{}),
		)
	} else {
		conf, err = config.LoadDefaultConfig(ctx)
	}

	if err != nil {
		return nil, err
	}

	if ar := o.AssumeRole; ar != nil {
		conf.Credentials = stscreds.NewAssumeRoleProvider(
			sts.NewFromConfig(conf),
			ar.RoleARN,
			func(p *stscreds.AssumeRoleOptions) {
				if ar.ExternalID != "" {
					p.ExternalID = aws.String(ar.ExternalID)
				}
			},
		)
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
