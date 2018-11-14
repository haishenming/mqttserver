package main

import (
	"mqttserver/client"
	"mqttserver/config"
	"mqttserver/router"
	
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path")
	
)

func main() {
	pflag.Parse()
	
	// 初始化配置
	if err := config.Init(*cfg); err != nil {
		panic(err)
		
	}
	
	// 连接配置
	connOpts := MQTT.NewClientOptions().AddBroker(viper.GetString("server")).SetClientID(viper.GetString("clientID"))
	connOpts.SetCleanSession(true)
	if viper.GetString("username") != "" {
		connOpts.SetUsername(viper.GetString("username"))
		if viper.GetString("password")  != "" {
			connOpts.SetPassword(viper.GetString("password"))
		}
	}
	
	log.Info("Add ClientOptions")
	
	// 初始化client
	err := client.Init(connOpts)
	if err != nil {
		log.Error("init client err", err)
		return
	}
	
	router.Init()
	
	defer client.Disconnect(uint(viper.GetInt("disconnect_quiesce")))
}