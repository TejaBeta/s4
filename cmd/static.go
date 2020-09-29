/*
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

const appType string = "static"

// staticCmd represents the static command
var staticCmd = &cobra.Command{
	Use:   "static",
	Short: "Static serves a website from a specific handler",
	Long: `Use the static command with S4 to serve a static
website with index.html pointing inside the object storage.`,
	Run: func(cmd *cobra.Command, args []string) {
		staticWebsite()
	},
}

func init() {
	rootCmd.AddCommand(staticCmd)
}

func staticWebsite() {

	fetcher := fileFetcher.Fetcher{IsAWS: isAWS, Bucket: bucket, AccessKey: accessKey, SecretKey: secretKey, Address: address, Region: region, AutoUpdate: autoUpdate, AppType: appType}

	fetcher.Run()
}
