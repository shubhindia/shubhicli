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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Prints a random joke",
	Long:  `Fetches a random joke from https://icanhazdadjoke.com `,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

}

type Joke struct {
	ID     string `json:"id`
	Joke   string `json:joke`
	Status int    `json:Status`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com"
	responseBytes := getJokeData(url)
	joke := Joke{}
	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		fmt.Println("error:", err)

	}
	fmt.Println(string(joke.Joke))
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		fmt.Println("Can't fetch a joke :(")
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Shubhicli (github.com/shubhindia/shubhicli)")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Println("error:", err)
	}
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error:", err)
	}

	return responseBytes
}
