package goutils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"net/http"
	"time"
)

func Svc() *sqs.SQS {
	sess := session.Must(session.NewSession(&aws.Config{
		MaxRetries:                    aws.Int(1),
		CredentialsChainVerboseErrors: aws.Bool(true),
		HTTPClient:                    &http.Client{Timeout: 10 * time.Second},
	}))

	cfgs := aws.Config{}
	regian := "us-east-1"
	cfgs.Region = &regian

	svc := sqs.New(sess, &cfgs)
	return svc
}
