package main

import (
	"flag"
	"log"
	"net/http"
	"s4/handlers"
)

var (
	bucket    string
	accessKey string
	secretKey string
	address   string
	region    string
)

func init() {
	flag.StringVar(&bucket, "bucket", "", "S3 bucket name")
	flag.StringVar(&accessKey, "accessKey", "", "AWS access key")
	flag.StringVar(&secretKey, "secretKey", "", "AWS secret key")
	flag.StringVar(&region, "region", "", "AWS Region Bucket resides")
	flag.StringVar(&address, "address", "127.0.0.1:3000", "address:port to serve the s3 content")
}

func main() {
	flag.Parse()

	s3 := handlers.S3Info{Bucket: bucket, AccessKey: accessKey, SecretKey: secretKey, Region: region}

	s3.BucketReader()

	for _, item := range s3.S3Objects {
		s3.ObjectDownloader(item.Name, "local")
	}

	log.Println("Server started listening on: ", address)
	fs := http.FileServer(http.Dir("./local"))
	log.Fatal(http.ListenAndServe(address, fs))

}
