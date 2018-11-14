package client

// MQTT 客户端单例
// 初始化之后全局可调用

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var (
	Client MQTTClient
)

type MQTTClient struct {
	MQTT.Client
}

func Init(options *MQTT.ClientOptions) {
	Client = MQTTClient{
		MQTT.NewClient(options),
	}
}

func Connect() error {
	if token := Client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func Disconnect(quiesce uint) {
	Client.Disconnect(quiesce)
}
