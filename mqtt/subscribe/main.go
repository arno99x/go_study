package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"time"
)

var messageHandler MQTT.MessageHandler = func(client MQTT.Client, message MQTT.Message) {
	fmt.Println("message id : ", message.Topic())
	fmt.Println("message payload : ", string(message.Payload()))
	message.Ack()
}

func main() {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("go-simple2")
	//opts.SetDefaultPublishHandler(messageHandler)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	if token := c.Subscribe("go-mqtt/test", 0, messageHandler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())

		//os.Exit(1)
	}

	time.Sleep(time.Second * 200)

	//unsubscribe from /go-mqtt/sample
	//if token := c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
	//	fmt.Println(token.Error())
	//	os.Exit(1)
	//}
	//
	//c.Disconnect(250000)
}
