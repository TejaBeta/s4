/*
Copyright © 2020 Tejasvi Thota <tejasvi.thota@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package controllers

import (
	"log"
	"net/http"
	"s4/handlers"
	"time"
)

type Fetcher struct {
	IsAWS      bool
	Bucket     string
	AccessKey  string
	SecretKey  string
	Address    string
	Region     string
	AutoUpdate bool
	LStore     map[string]time.Time
	AppType    string
}

func (fetcher *Fetcher) Run() {
	fetcher.LStore = make(map[string]time.Time)

	switch {
	case fetcher.IsAWS:
		fetcher.s3Handle()
		if fetcher.AutoUpdate {
			go fetcher.autoUpdater()
		}
		break
	}

	fs := http.FileServer(http.Dir("./local"))
	log.Fatal(http.ListenAndServe(fetcher.Address, fs))
	log.Println("Server started listening on: ", fetcher.Address)
}

func (fetcher *Fetcher) s3Handle() {
	s3 := handlers.S3Info{Bucket: fetcher.Bucket, AccessKey: fetcher.AccessKey, SecretKey: fetcher.SecretKey, Region: fetcher.Region}

	s3.BucketReader()

	for _, item := range s3.S3Objects {
		if v, ok := fetcher.LStore[item.Name]; !ok || v != item.LastModified {
			fetcher.LStore[item.Name] = item.LastModified
			s3.ObjectDownloader(item.Name, "local")
		}
	}
}

func (fetcher *Fetcher) autoUpdater() {
	autoUpdater := time.NewTicker(15 * time.Minute)
	for {
		select {
		case <-autoUpdater.C:
			fetcher.s3Handle()
		}
	}
}
