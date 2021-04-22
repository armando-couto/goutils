package goutils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var AccessKeyID string
var SecretAccessKey string
var MyRegion string

func ConnectAws() *session.Session {
	AccessKeyID = Godotenv("AWS_ACCESS_KEY_ID")
	SecretAccessKey = Godotenv("AWS_SECRET_ACCESS_KEY")
	MyRegion = Godotenv("AWS_REGION")
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}
