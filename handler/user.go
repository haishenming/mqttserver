package handler

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"

)

func UserHandler(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("UserHandler      ")
	fmt.Printf("[%s]  ", msg.Topic())
	fmt.Printf("%s\n", msg.Payload())
}
