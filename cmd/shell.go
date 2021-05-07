/*
Copyright © 2021 shubhindia <shubhindia123@gmail.com>

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
	"os/exec"
	"syscall"

	"github.com/spf13/cobra"
)

// shellCmd represents the shell command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "A command used to ssh into my devices",
	Long:  `A simple command to ssh into my devices connected in same network.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		device, _ := cmd.Flags().GetString("device")
		user, _ := cmd.Flags().GetString("user")
		if user == "" {
			fmt.Println("Switching to default user shubhcyanogen")
			user = "shubhcyanogen"
		}

		if device != "" {
			return sshIntoDevice(device, user)

		} else {
			fmt.Println("Error. You must enter a device name")
			err := "no device"
			panic(err)

		}
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)

	shellCmd.PersistentFlags().String("device", "", "Device name to which you want to ssh.")
	shellCmd.PersistentFlags().String("user", "", "SSH username")

}

func Exec(command []string) error {
	binary, err := exec.LookPath(command[0])
	if err != nil {
		return err
	}
	err = syscall.Exec(binary, command, os.Environ())
	// Panic rather than returning since this should never happen.
	panic(err)
}

func sshIntoDevice(device string, user string) error {
	fmt.Printf("SSHing into %v with user %v", device, user)
	/* Adding a temporary logic to get ip of the device from env variables. This is not the most optimized way to do this,
	but will use this untill I figure out an optimised way */
	deviceIP := ""
	if device == "warmachine" {
		deviceIP = os.Getenv("warmachine")
	}
	if device == "rpi" {
		deviceIP = os.Getenv("rpi")
	}
	command := []string{"ssh", user + "@" + deviceIP}
	return Exec(command)
}
