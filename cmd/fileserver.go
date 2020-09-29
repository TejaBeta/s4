/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	fileFetcher "github.com/tejabeta/s4/internal/fetcher"

	"github.com/spf13/cobra"
)

// fileserverCmd represents the fileserver command
var fileserverCmd = &cobra.Command{
	Use:   "fileserver",
	Short: "Serves the purpose of fileserver from a specific handler",
	Long: `Use the fileserver command with S4 to serve as
a fileserver similar to NFS but stores files locally as well.`,
	Run: func(cmd *cobra.Command, args []string) {
		fileServer()
	},
}

func init() {
	rootCmd.AddCommand(fileserverCmd)
}

func fileServer() {

	fetcher := fileFetcher.Fetcher{IsAWS: isAWS, Bucket: bucket, AccessKey: accessKey, SecretKey: secretKey, Address: address, Region: region, AutoUpdate: autoUpdate, AppType: "fileserver"}

	fetcher.Run()
}
