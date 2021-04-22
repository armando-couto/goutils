package goutils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"math/rand"
	"time"
)

/*
	ConectionSQS o antigo nome era: Svc
*/
func ConectionSQS() *sqs.SQS {
	// The session the S3 Uploader will use
	sess := ConnectAws()
	regian := "us-east-1"
	cfgs := aws.Config{Region: &regian}
	return sqs.New(sess, &cfgs)
}

/*
	TokenGeneratorMessageId get token do Message.
*/
func TokenGeneratorMessageId() string {
	rand.Seed(time.Now().UnixNano())
	return string(fmt.Sprint(RandSeq(8), "-", RandSeq(4), "-", RandSeq(4), "-", RandSeq(4), "-", RandSeq(12)))
}
