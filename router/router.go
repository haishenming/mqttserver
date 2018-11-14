package router

import (
	"mqttserver/client"
	"mqttserver/handler"
)

func Init() {
	client.Client.AddRoute("user", handler.UserHandler)
}
