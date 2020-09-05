package main

import (
	"flag"
	"log"
	"net/http"
	"s4/handlers"
	"time"
)

var (
	bucket     string
	accessKey  string
	secretKey  string
	address    string
	region     string
	autoUpdate bool
	lstore     map[string]time.Time
)

func init() {
	flag.StringVar(&bucket, "bucket", "", "S3 bucket name")
	flag.StringVar(&accessKey, "accessKey", "", "AWS access key")
	flag.StringVar(&secretKey, "secretKey", "", "AWS secret key")
	flag.StringVar(&region, "region", "", "AWS Region Bucket resides")
	flag.StringVar(&address, "address", "127.0.0.1:3000", "address:port to serve the s3 content")
	flag.BoolVar(&autoUpdate, "autoUpdate", true, "Bool to auto update")
}

func main() {
	flag.Parse()

	lstore = make(map[string]time.Time)

	aws := true

	if aws {
		s3Handle()
		if autoUpdate {
			go autoUpdater()
		}
	}

	fs := http.FileServer(http.Dir("./local"))
	log.Fatal(http.ListenAndServe(address, fs))
	log.Println("Server started listening on: ", address)
}

func s3Handle() {
	s3 := handlers.S3Info{Bucket: bucket, AccessKey: accessKey, SecretKey: secretKey, Region: region}

	s3.BucketReader()

	for _, item := range s3.S3Objects {
		if v, ok := lstore[item.Name]; !ok || v != item.LastModified {
			lstore[item.Name] = item.LastModified
			s3.ObjectDownloader(item.Name, "local")
		}
	}
}

func autoUpdater() {

	autoUpdater := time.NewTicker(15 * time.Minute)

	for {
		select {
		case <-autoUpdater.C:
			s3Handle()
		}
	}
}
