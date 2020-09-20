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
package cmd

import (
	"flag"
	"log"
	"net/http"
	"s4/handlers"
	"time"

	"github.com/spf13/cobra"
)

// staticCmd represents the static command
var staticCmd = &cobra.Command{
	Use:   "static",
	Short: "Static serves a website from a specific handler",
	Long: `Use the static command with S4 to serve a static
website with index.html pointing inside the object storage.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("static called")
		staticWebsite()
	},
}

func init() {
	rootCmd.AddCommand(staticCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// staticCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// staticCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func staticWebsite() {
	flag.Parse()

	lstore = make(map[string]time.Time)

	switch {
	case isAWS:
		s3Handle()
		if autoUpdate {
			go autoUpdater()
		}
		break
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
