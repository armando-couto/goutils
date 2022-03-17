package goutils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"mime/multipart"
	"time"
)

func UpdloadInS3(file multipart.File, path, fileName string) string {
	// The session the S3 Uploader will use
	sess := ConnectAws()

	fileName = fmt.Sprint(time.Now().Format(LAYOUT_YYYYMMDDHHMMSS), fileName)

	// Caso tenha om PATH ai ele concatena
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
		CreateFileDayError(FormatMessage(MessageError{Error: err.Error()}))
		return ""
	}
	CreateFileDayInfo(fmt.Sprint("Upload do Arquivo: ", up.UploadID))
	return "https://" + myBucket + "." + "s3.amazonaws.com/" + fileName
}

func UpdloadInS3Base64(b64 string, path, fileName string) string {
	// The session the S3 Uploader will use
	sess := ConnectAws()

	fileName = fmt.Sprint(time.Now().Format(LAYOUT_YYYYMMDDHHMMSS), fileName)

	// Caso tenha om PATH ai ele concatena
	if path != "" {
		fileName = fmt.Sprint(path, "/", fileName)
	}

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	myBucket := Godotenv("BUCKET_NAME")

	dec, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		CreateFileDayError(FormatMessage(MessageError{Error: err.Error()}))
		return ""
	}
	//file, header, err := c.Request.FormFile("photo")
	//upload to the s3 bucket
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(fileName),
		Body:   bytes.NewBuffer(dec),
	})
	if err != nil {
		CreateFileDayError(FormatMessage(MessageError{Error: err.Error()}))
		return ""
	}
	CreateFileDayInfo(fmt.Sprint("Upload do Arquivo: ", up.UploadID))
	return "https://" + myBucket + "." + "s3.amazonaws.com/" + fileName
}

func UpdloadInS3Base64Byte(b64 []byte, path, fileName string) string {
	// The session the S3 Uploader will use
	sess := ConnectAws()

	fileName = fmt.Sprint(time.Now().Format(LAYOUT_YYYYMMDDHHMMSS), fileName)

	// Caso tenha om PATH ai ele concatena
	if path != "" {
		fileName = fmt.Sprint(path, "/", fileName)
	}

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	myBucket := Godotenv("BUCKET_NAME")

	//upload to the s3 bucket
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(fileName),
		Body:   bytes.NewBuffer(b64),
	})
	if err != nil {
		CreateFileDayError(FormatMessage(MessageError{Error: err.Error()}))
		return ""
	}
	CreateFileDayInfo(fmt.Sprint("Upload do Arquivo: ", up.UploadID))
	return "https://" + myBucket + "." + "s3.amazonaws.com/" + fileName
}
