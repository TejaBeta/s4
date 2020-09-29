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
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var (
	isAWS      bool
	bucket     string
	accessKey  string
	secretKey  string
	address    string
	region     string
	autoUpdate bool
	localDir   string
	lstore     map[string]time.Time
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "s4",
	Short: "A CLI middleware to build services on AWS S3",
	Long: `A tiny CLI tool to that acts as a middleware to build
services making use of AWS S3 object store as a backend. 

Currently supports hosting a static website from private AWS S3
object store with pointing to index.html. And also supports the
hosting a private PyPi server with S3 object store as package storage.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.s4.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolVar(&isAWS, "isAWS", true, "Bool to pick a platform")
	rootCmd.PersistentFlags().StringVar(&bucket, "bucket", "", "S3 bucket name")
	rootCmd.PersistentFlags().StringVar(&accessKey, "accessKey", "", "AWS access key")
	rootCmd.PersistentFlags().StringVar(&secretKey, "secretKey", "", "AWS secret key")
	rootCmd.PersistentFlags().StringVar(&region, "region", "", "AWS Region Bucket resides")
	rootCmd.PersistentFlags().StringVar(&localDir, "localDir", "./local", "Local directory to sync and serve")
	rootCmd.PersistentFlags().StringVar(&address, "address", "127.0.0.1:3000", "address:port to serve the s3 content")
	rootCmd.PersistentFlags().BoolVar(&autoUpdate, "autoUpdate", true, "Bool to auto update")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".s4" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".s4")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
