package goutils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"mime/multipart"
)

func UpdloadInS3(file multipart.File, path, fileName string) string {
	// The session the S3 Uploader will use
	sess := ConnectAws()

	if path != "" {
		fileName = fmt.Sprint(path, "/", fileName)
	}

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	myBucket := Godotenv("BUCKET_NAME")

	//file, header, err := c.Request.FormFile("photo")
	//upload to the s3 bucket
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		CreateFileDay(err.Error())
		return ""
	}
	CreateFileDay(fmt.Sprint("Upload do Arquivo: ", up.UploadID))
	return "https://" + myBucket + "." + "s3-" + MyRegion + ".amazonaws.com/" + fileName
}
