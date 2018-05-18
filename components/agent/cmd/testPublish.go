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
	"log"

	"github.com/chuckleheads/hurtlocker/components/agent/proto/build"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
)

// testPublishCmd represents the testPublish command
var testPublishCmd = &cobra.Command{
	Use:   "testPublish",
	Short: "A command to test publishing to our builder agent",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := RabbitMQConfigFromViper()
		connString := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Username, config.Password, config.Host, config.Port)
		fmt.Printf(connString)
		conn, err := amqp.Dial(connString)
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

		err = ch.ExchangeDeclare(
			config.Exchange, // name
			"topic",         // type
			true,            // durable
			false,           // auto-deleted
			false,           // internal
			false,           // no-wait
			nil,             // arguments
		)
		failOnError(err, "Failed to declare an exchange")

		buildReq := &build.Build{
			PackagePath: ".",
		}

		body, err := proto.Marshal(buildReq)
		if err != nil {
			log.Fatalln("Failed to encode address book:", err)
		}
		err = ch.Publish(
			config.Exchange, // exchange
			"build",         // routing key
			false,           // mandatory
			false,           // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Failed to publish a message")

		log.Printf(" [x] Sent %s", body)
	},
}

func init() {
	rootCmd.AddCommand(testPublishCmd)
}
