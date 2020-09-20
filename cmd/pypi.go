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
	"fmt"

	"github.com/spf13/cobra"
)

// pypiCmd represents the pypi command
var pypiCmd = &cobra.Command{
	Use:   "pypi",
	Short: "PyPi command works as Python Index server with respective handler",
	Long: `PyPi command helps to serve S4 as private PyPi server using the
respective handler as the storage system.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pypi called")
	},
}

func init() {
	rootCmd.AddCommand(pypiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pypiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pypiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
