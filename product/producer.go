package product

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Producer struct {
}

func (c *Producer) ProduceProductCreated(delivery amqp.Delivery) error {

	payload := string(delivery.Body)
	fmt.Printf("ConsumeProductCreated, %s \n", payload)

	return nil
}

func (c *Producer) ProduceProductUpdated(delivery amqp.Delivery) error {

	payload := string(delivery.Body)
	fmt.Printf("ConsumeProductUpdated, %s \n", payload)

	return nil
}

func (c *Producer) ProduceProductStatusChanged(delivery amqp.Delivery) error {

	payload := string(delivery.Body)
	fmt.Printf("ConsumeProductStatusChanged, %s \n", payload)

	return nil
}
