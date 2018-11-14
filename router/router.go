package router

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/lexkong/log"
	"mqttserver/client"
	"mqttserver/handler"
)

type router struct {
	route      string
	handleFunc MQTT.MessageHandler
	qos        byte
}

var mux = []*router{
	{route: "user", handleFunc: handler.UserHandler, qos: 1},
}

func Init() error {
	for _, r := range mux {
		if token := client.Client.Subscribe(r.route, r.qos, r.handleFunc); token.Wait() && token.Error() != nil {
			log.Errorf(token.Error(), "%s sub err", r.route)
			return token.Error()
		}
	}

	return nil
}
