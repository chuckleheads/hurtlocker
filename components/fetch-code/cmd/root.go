// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/spf13/cobra"
	git "gopkg.in/src-d/go-git.v4"
)

var path string
var key string
var appID int
var installID int
var URL string
var isPrivate bool

var rootCmd = &cobra.Command{
	Use:   "fetch-code",
	Short: "Fetches source code from GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		customURL, err := url.Parse(URL)
		if err != nil {
			panic(err)
		}

		if isPrivate {
			itr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, appID, installID, key)
			if err != nil {
				panic(err)
			}
			token, err := itr.Token()
			customURL.User = url.UserPassword("x-access-token", token)
		}

		_, err = git.PlainClone(path, false, &git.CloneOptions{
			URL:      customURL.String(),
			Progress: os.Stdout,
		})
		if err != nil {
			panic(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&path, "path", "", "Path to clone repo into")
	rootCmd.Flags().StringVar(&URL, "url", "", "Backend Type. Example: github.com")
	rootCmd.Flags().IntVarP(&installID, "installId", "i", 0, "GitHub Install ID")
	rootCmd.Flags().IntVarP(&appID, "appId", "a", 0, "GitHub App ID")
	rootCmd.Flags().StringVarP(&key, "key", "k", "", "GitHub Secret Key")
	rootCmd.Flags().BoolVarP(&isPrivate, "private", "p", false, "Mark a repo as private")
	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("path")
}
