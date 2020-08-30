package main

import (
	"flag"
	"fmt"
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
	flag.StringVar(&address, "address", "", "address:port to serve the s3 content")
	flag.StringVar(&sslCert, "sslCert", "", "Location to SSL certificate")
}

func main() {
	flag.Parse()
	fmt.Println("Hello, I am S4!")
}
