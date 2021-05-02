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
	"log"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List IP and MAC",
	Long:  `Lists IP and MAC addresses of devices connected in your home network.`,
	Run: func(cmd *cobra.Command, args []string) {
		ListDevices()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}

type Devices struct {
	IP  string `json:IP`
	MAC string `json:MAC`
}

func ListDevices() {
	output, err := exec.Command("arp", "-a").Output()

	if err != nil {
		log.Fatal(err)
	}
	numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	regexPattern := numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock
	re := regexp.MustCompile(regexPattern)

	IPs := re.FindAllString(string(output), -1)
	FinalDevices := []Devices{}
	for i := 0; i < len(IPs); i++ {
		IP := IPs[i]
		FinalDevices = append(FinalDevices, Devices{
			IP:  IP,
			MAC: "NIL",
		})
	}
	for i := 0; i < len(FinalDevices); i++ {
		fmt.Println("Device - ", i)
		fmt.Println("IP:", FinalDevices[i].IP)
		fmt.Println("MAC:", FinalDevices[i].MAC)
	}

}
