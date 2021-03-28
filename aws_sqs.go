package goutils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"math/rand"
	"net/http"
	"time"
)

/*
	ConectionSQS o antigo nome era: Svc
*/
func ConectionSQS() *sqs.SQS {
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

/*
	TokenGeneratorMessageId get token do Message.
*/
func TokenGeneratorMessageId() string {
	rand.Seed(time.Now().UnixNano())
	return string(fmt.Sprint(RandSeq(8), "-", RandSeq(4), "-", RandSeq(4), "-", RandSeq(4), "-", RandSeq(12)))
}
