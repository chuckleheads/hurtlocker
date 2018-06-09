package config

type DriverType int

const (
	Kubernetes DriverType = iota + 1
	Docker
)

type Config struct {
	RabbitMQ RabbitMQConfig
	Driver   DriverType
}
