package cmd

import (
	"fmt"
	"log"

	"github.com/chuckleheads/hurtlocker/components/agent/dispatcher"
	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Builder Agent",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := ConfigFromViper()
		rmqConfig := config.RabbitMQ
		conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", rmqConfig.Username, rmqConfig.Password, rmqConfig.Host, rmqConfig.Port))
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()
		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

		err = ch.ExchangeDeclare(
			rmqConfig.Exchange, // name
			"topic",            // type
			true,               // durable
			false,              // auto-deleted
			false,              // internal
			false,              // no-wait
			nil,                // arguments
		)
		failOnError(err, "Failed to declare an exchange")

		q, err := ch.QueueDeclare(
			"",    // name
			false, // durable
			false, // delete when usused
			true,  // exclusive
			false, // no-wait
			nil,   // arguments
		)
		failOnError(err, "Failed to declare a queue")

		for _, s := range rmqConfig.Topic {
			log.Printf("Binding queue %s to exchange %s with routing key %s",
				q.Name, rmqConfig.Exchange, s)
			err = ch.QueueBind(
				q.Name,             // queue name
				s,                  // routing key
				rmqConfig.Exchange, // exchange
				false,
				nil)
			failOnError(err, "Failed to bind a queue")
		}

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto ack
			false,  // exclusive
			false,  // no local
			false,  // no wait
			nil,    // args
		)
		failOnError(err, "Failed to register a consumer")

		forever := make(chan bool)

		go func() {
			for d := range msgs {
				switch d.RoutingKey {
				case "build":
					dispatcher.Build(d.Body)
				case "deploy":
					dispatcher.Deploy(d.Body)
				case "export":
					dispatcher.Export(d.Body)
				default:
					panic("unreachable")
				}
			}
		}()

		log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
		<-forever
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
