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
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
)

// awsCmd represents the aws command
var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Just a bunch of AWS commands",
	Long:  `Bunch of AWS commands using AWS sdk for go`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic("Invalid opetaion")
		}

		operation := args[0]
		switch operation {
		case "s3":
			s3Commands(args)
		case "ec2":
			ec2Commands()
		default:
			fmt.Println("Invalid operation")
		}

	},
}

func init() {
	rootCmd.AddCommand(awsCmd)

}
func awsConfig() aws.Config {
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	return cfg
}

func s3Commands(args []string) {
	fmt.Println("S3 operation called", args[1])
	operation := args[1]
	// switch based on operations called
	switch operation {
	case "ls":

		cfg := awsConfig()
		client := s3.NewFromConfig(cfg)
		fmt.Println(client)
	}

}

func ec2Commands() {
	fmt.Println("EC2 operation called")
}
