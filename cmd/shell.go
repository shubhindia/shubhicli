/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "A command used to ssh into my devices",
	Long:  `A simple command to ssh into my devices connected in same network.`,
	Run: func(cmd *cobra.Command, args []string) {
		device, _ := cmd.Flags().GetString("device")

		if device != "" {
			sshIntoDevice(device)
		} else {
			fmt.Println("Error. You must enter a device name")
		}
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	shellCmd.PersistentFlags().String("device", "", "Device name to which you want to ssh.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shellCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func sshIntoDevice(device string) {
	fmt.Println("SSHing into ", device)
}
