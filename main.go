package main

import (
	"fmt"
	"main/config"
	"main/product"
	"main/rabbit"
	"main/supplier"
)

func main() {

	configurationManager := config.NewConfigurationManager()
	rabbitConfig := configurationManager.GetRabbitConfig()
	queuesConfig := configurationManager.GetQueuesConfig()
	productConsumer := product.Consumer{}
	supplierConsumer := supplier.Consumer{}

	rabbitClient := rabbit.NewRabbitClient(rabbitConfig, queuesConfig, productConsumer, supplierConsumer)
	defer rabbitClient.CloseConnection()

	rabbitClient.DeclareExchangeQueueBindings()

	consumerChan := make(chan bool)

	rabbitClient.InitializeConsumers()

	fmt.Println("Started consumers")

	<-consumerChan // close(consumerChan)
}
