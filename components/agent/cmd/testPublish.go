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

	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
)

// testPublishCmd represents the testPublish command
var testPublishCmd = &cobra.Command{
	Use:   "testPublish",
	Short: "A command to test publishing to our builder agent",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := ConfigFromViper()
		rmqConfig := config.RabbitMQ
		connString := fmt.Sprintf("amqp://%s:%s@%s:%d/", rmqConfig.Username, rmqConfig.Password, rmqConfig.Host, rmqConfig.Port)
		fmt.Printf(connString)
		conn, err := amqp.Dial(connString)
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

		err = ch.ExchangeDeclare(
			rmqConfig.Exchange, // name
			"direct",           // type
			true,               // durable
			false,              // auto-deleted
			false,              // internal
			false,              // no-wait
			nil,                // arguments
		)
		failOnError(err, "Failed to declare an exchange")
		// publishBuild(ch, rmqConfig)
		// publishExport(ch, rmqConfig)
	},
}

func init() {
	rootCmd.AddCommand(testPublishCmd)
}

// func publishBuild(ch *amqp.Channel, rmqConfig config.RabbitMQConfig) {
// 	buildReq := &build.Build{
// 		PlanPath:    "redis",
// 		RepoUrl:     "https://github.com/elliott-davis/core-plans",
// 		Channel:     "unstable",
// 		Environment: []string{"HAB_ORIGIN=edavis"},
// 	}

// 	body, err := proto.Marshal(buildReq)
// 	if err != nil {
// 		log.Fatalln("Failed to encode address book:", err)
// 	}
// 	err = ch.Publish(
// 		rmqConfig.Exchange, // exchange
// 		"build.linux",      // routing key
// 		false,              // mandatory
// 		false,              // immediate
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte(body),
// 		})
// 	failOnError(err, "Failed to publish a message")

// 	log.Printf(" [x] Sent %s", body)
// }

// func publishExport(ch *amqp.Channel, rmqConfig config.RabbitMQConfig) {
// 	exportReq := &export.Export{
// 		Ident:   "core/redis",
// 		Channel: "unstable",
// 	}

// 	body, err := proto.Marshal(exportReq)
// 	if err != nil {
// 		log.Fatalln("Failed to encode address book:", err)
// 	}
// 	err = ch.Publish(
// 		rmqConfig.Exchange, // exchange
// 		"export.docker",    // routing key
// 		false,              // mandatory
// 		false,              // immediate
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte(body),
// 		})
// 	failOnError(err, "Failed to publish a message")

// 	log.Printf(" [x] Sent %s", body)
// }
