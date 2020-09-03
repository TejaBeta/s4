package handlers

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// S3ObjectInfo is a struct holds the metadata related to s3 object
type S3ObjectInfo struct {
	Name         string
	LastModified time.Time
	Size         int64
	StorageClass string
}

// S3Info struct containing basics what is required to read
type S3Info struct {
	Bucket    string
	AccessKey string
	SecretKey string
	Region    string
	S3Objects []S3ObjectInfo
}

// BucketReader will hold the logic related reading and checking the content of a bucket
func (s3Info *S3Info) BucketReader() {
	os.Setenv("AWS_ACCESS_KEY", s3Info.AccessKey)
	os.Setenv("AWS_SECRET_KEY", s3Info.SecretKey)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s3Info.Region), Credentials: credentials.NewEnvCredentials()},
	)

	svc := s3.New(sess)

	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(s3Info.Bucket)})

	if err != nil {
		log.Fatal("Unable to list items in bucket %q, %v", s3Info.Bucket, err)
	}

	for _, item := range resp.Contents {
		s3ObjectInfo := S3ObjectInfo{Name: *item.Key, LastModified: *item.LastModified, Size: *item.Size, StorageClass: *item.StorageClass}
		s3Info.S3Objects = append(s3Info.S3Objects, s3ObjectInfo)
	}
}

// ObjectDownloader will download specific objects from s3 bucket
func (s3Info *S3Info) ObjectDownloader(ObjectName string, dirLocation string) {
	dirLocation = dirLocation + "/"
	dir := strings.Split(ObjectName, "/")

	finalDir := strings.Join(dir[:len(dir)-1], "/")

	if _, err := os.Stat(dirLocation + finalDir); os.IsNotExist(err) {
		os.MkdirAll(dirLocation+finalDir, 0700)
	}

	file, err := os.Create(dirLocation + ObjectName)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s3Info.Region), Credentials: credentials.NewEnvCredentials()},
	)

	downloader := s3manager.NewDownloader(sess)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(s3Info.Bucket),
			Key:    aws.String(ObjectName),
		})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}
