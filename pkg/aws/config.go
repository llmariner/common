package aws

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// AssumeRole is the options for AssumeRole.
type AssumeRole struct {
	RoleARN    string
	ExternalID string
}

// Secret is the AWS secret.
type Secret struct {
	AccessKeyID     string
	SecretAccessKey string
}

// NewConfigOptions is the configuration options.
type NewConfigOptions struct {
	Region string

	UseAnonymousCredentials bool

	AssumeRole *AssumeRole
	Secret     *Secret

	InsecureSkipVerify bool
}

// NewConfig returns a new configuration.
func NewConfig(ctx context.Context, o NewConfigOptions) (aws.Config, error) {
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
		return conf, err
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
	} else if s := o.Secret; s != nil {
		conf.Credentials = credentials.NewStaticCredentialsProvider(
			s.AccessKeyID,
			s.SecretAccessKey,
			"", // Session token is empty for static credentials
		)
	}

	conf.Region = o.Region

	if o.InsecureSkipVerify {
		conf.HTTPClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
	}

	return conf, nil
}
