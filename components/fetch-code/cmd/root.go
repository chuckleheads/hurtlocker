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

var key string
var appID int
var installID int
var repo string
var rootCmd = &cobra.Command{
	Use:   "fetch-code",
	Short: "Fetches source code from GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		itr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, appID, installID, key)
		if err != nil {
			panic("Oh Snap")
		}
		token, err := itr.Token()

		customURL := url.URL{
			Scheme: "https",
			User:   url.UserPassword("x-access-token", token),
			Host:   "github.com",
			Path:   fmt.Sprintf("%s.git", repo),
		}
		_, err = git.PlainClone("/tmp/foo", false, &git.CloneOptions{
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
	rootCmd.Flags().String("host", "github.com", "Backend Type. Example: github.com")
	rootCmd.Flags().IntVarP(&installID, "installId", "i", 0, "GitHub Install ID")
	rootCmd.Flags().IntVarP(&appID, "appId", "a", 0, "GitHub App ID")
	rootCmd.Flags().StringVarP(&key, "key", "k", "", "GitHub Secret Key")
	rootCmd.Flags().StringVarP(&repo, "repo", "r", "", "Repo to Clone")
	rootCmd.MarkFlagRequired("repo")
}
