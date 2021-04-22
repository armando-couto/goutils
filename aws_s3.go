package goutils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"mime/multipart"
)

var filepath string

func UpdloadInS3(file multipart.File, header multipart.FileHeader) string {
	// The session the S3 Uploader will use
	sess := ConnectAws()

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	myBucket := Godotenv("BUCKET_NAME")

	//file, header, err := c.Request.FormFile("photo")
	filename := header.Filename
	//upload to the s3 bucket
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		CreateFileDay(err.Error())
		return ""
	}
	CreateFileDay(fmt.Sprint("Upload do Arquivo: ", up.UploadID))
	return "https://" + myBucket + "." + "s3-" + MyRegion + ".amazonaws.com/" + filename
}
