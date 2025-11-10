package goutils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"mime/multipart"
	"os"
	"strings"
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
		CreateFileDay(Message{Error: err.Error()}, &MessageGotifyGlobal)
		return ""
	}
	CreateFileDay(Message{Info: fmt.Sprint("Upload do Arquivo: ", up.UploadID)}, nil)
	return "https://" + myBucket + "." + "s3.amazonaws.com/" + fileName
}

func UpdloadInS3NotTime(file multipart.File, path, fileName string) string {
	// The session the S3 Uploader will use
	sess := ConnectAws()

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
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		CreateFileDay(Message{Error: err.Error()}, &MessageGotifyGlobal)
		return ""
	}
	CreateFileDay(Message{Info: fmt.Sprint("Upload do Arquivo: ", up.UploadID)}, nil)
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
		CreateFileDay(Message{Error: err.Error()}, &MessageGotifyGlobal)
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
		CreateFileDay(Message{Error: err.Error()}, &MessageGotifyGlobal)
		return ""
	}
	CreateFileDay(Message{Info: fmt.Sprint("Upload do Arquivo: ", up.UploadID)}, nil)
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
		CreateFileDay(Message{Error: err.Error()}, &MessageGotifyGlobal)
		return ""
	}
	CreateFileDay(Message{Info: fmt.Sprint("Upload do Arquivo: ", up.UploadID)}, nil)
	return "https://" + myBucket + "." + "s3.amazonaws.com/" + fileName
}

func UpdloadInS3ArqTxt(texto string, path, fileName string) string {
	// The session the S3 Uploader will use
	sess := ConnectAws()

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
		Body:   strings.NewReader(texto),
	})
	if err != nil {
		CreateFileDay(Message{Error: err.Error()}, &MessageGotifyGlobal)
		return ""
	}
	CreateFileDay(Message{Info: fmt.Sprint("Upload do Arquivo: ", up.UploadID)}, nil)
	return "https://" + myBucket + "." + "s3.amazonaws.com/" + fileName
}

func DownloadFromS3Public(fileName, localFilePath string) error {
	sess := ConnectAws()
	downloader := s3.New(sess)

	file, err := os.Create(localFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	bucketName := Godotenv("BUCKET_NAME")
	_, err = downloader.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
	})

	return err
}

func DownloadFromS3NotPublic(fileName, localFilePath string) error {
	accessKey := Godotenv("AWS_ACCESS_KEY_ID")
	secretKey := Godotenv("AWS_SECRET_ACCESS_KEY")

	// Configuração das credenciais
	creds := credentials.NewStaticCredentials(accessKey, secretKey, "")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(Godotenv("AWS_REGION")), // Substitua pela sua região AWS
		Credentials: creds,
	})

	downloader := s3.New(sess)

	file, err := os.Create(localFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	bucketName := Godotenv("BUCKET_NAME")
	_, err = downloader.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
	})

	return err
}
