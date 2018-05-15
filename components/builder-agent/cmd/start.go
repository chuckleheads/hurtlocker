package cmd

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Builder Agent",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := RabbitMQConfigFromViper()
		conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Username, config.Password, config.Host, config.Port))
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()
		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

		// only fetch one message at a time
		err = ch.Qos(
			1,     // prefetch count
			0,     // prefetch size
			false, // global
		)
		failOnError(err, "Failed to set QoS")

		q, err := ch.QueueDeclare(
			"task_queue", // name
			true,         // durable
			false,        // delete when usused
			false,        // exclusive
			false,        // no-wait
			nil,          // arguments
		)
		failOnError(err, "Failed to declare a queue")

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			false,  // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		failOnError(err, "Failed to register a consumer")

		forever := make(chan bool)

		go func() {
			for d := range msgs {
				log.Printf("Received a message: %s", d.Body)
				dotCount := bytes.Count(d.Body, []byte("."))
				t := time.Duration(dotCount)
				time.Sleep(t * time.Second)
				log.Printf("Done")
				d.Ack(false)
			}
		}()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
