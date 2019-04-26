package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"time"
)

//var messageHandler MQTT.MessageHandler = func(client MQTT.Client, message MQTT.Message) {
//	fmt.Println("message id : ",message.Topic())
//	fmt.Println("message payload : ",message.Payload())
//}
func main() {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("go-simple1")

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
	//from the server after sending each message
	for i := 0; i < 60000; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := c.Publish("go-mqtt/test", 0, true, text)
		fmt.Println("send : ", text)
		token.Wait()
		time.Sleep(time.Second * 2)
	}

}
