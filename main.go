package main

import (
	"flag"
	"log"
	"net/http"
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
	fs := http.FileServer(http.Dir("./static_test"))

	log.Println("Server started listening on: ", address)
	log.Fatal(http.ListenAndServe(address, fs))
}
