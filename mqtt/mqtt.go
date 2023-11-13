package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
	"log"
	"time"
)

var Client mqtt.Client

func Connect() {
	opts := mqtt.NewClientOptions()
	opts.SetClientID(viper.GetString("mqtt.client_id"))
	opts.SetCleanSession(true)
	opts.SetConnectRetryInterval(time.Second * 5)
	opts.SetConnectTimeout(time.Second)
	opts.SetAutoReconnect(true)
	opts.SetConnectionLostHandler(mqttConnectionLost)
	opts.SetOnConnectHandler(mqttConnected)

	opts.AddBroker(viper.GetString("mqtt.broker"))
	if user := viper.GetString("mqtt.username"); user != "" {
		opts.SetUsername(user)
	}
	if pass := viper.GetString("mqtt.password"); pass != "" {
		opts.SetPassword(pass)
	}
	Client = mqtt.NewClient(opts)

	for token := Client.Connect(); token.Wait() && token.Error() != nil; token = Client.Connect() {
		log.Println("Screen: error connecting:", token.Error())
		time.Sleep(time.Second * 5)
	}
}

func mqttConnectionLost(_ mqtt.Client, err error) {
	log.Println("Screen: Connection lost:", err)
}

func mqttConnected(_ mqtt.Client) {
	log.Println("Screen: Connected!")
	Client.Subscribe("kistan/in_tv2/#", 0, mqttMessageReceived)
}

func mqttMessageReceived(_ mqtt.Client, msg mqtt.Message) {

}
