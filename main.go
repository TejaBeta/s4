package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	bucket    string
	accessKey string
	secretKey string
	address   string
	sslCert   string
)

func init() {
	flag.StringVar(&bucket, "bucket", "", "S3 bucket name")
	flag.StringVar(&accessKey, "accessKey", "", "AWS access key")
	flag.StringVar(&secretKey, "secretKey", "", "AWS secret key")
	flag.StringVar(&address, "address", "127.0.0.1:3000", "address:port to serve the s3 content")
}

func main() {
	flag.Parse()
	fs := http.FileServer(http.Dir("./local"))

	s3BucketContent()

	log.Println("Server started listening on: ", address)
	log.Fatal(http.ListenAndServe(address, fs))

}

func s3BucketContent() {

	os.Setenv("AWS_ACCESS_KEY", accessKey)
	os.Setenv("AWS_SECRET_KEY", secretKey)

	// fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-2"), Credentials: credentials.NewEnvCredentials()},
	)

	// Create S3 service client
	svc := s3.New(sess)

	// Get the list of items
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
	if err != nil {
		exitErrorf("Unable to list items in bucket %q, %v", bucket, err)
	}

	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")

		dir := strings.Split(*item.Key, "/")

		finalDir := strings.Join(dir[:len(dir)-1], "/")

		// fmt.Println(strings.Join(dir[:len(dir)-1], "/"))

		if _, err := os.Stat("local/" + finalDir); os.IsNotExist(err) {
			// your file does not exist
			os.MkdirAll("local/"+finalDir, 0700)
		}

		file, err := os.Create("local/" + *item.Key)

		if err != nil {
			fmt.Println(err)
		}

		defer file.Close()

		downloader := s3manager.NewDownloader(sess)
		numBytes, err := downloader.Download(file,
			&s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(*item.Key),
			})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

	}
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
