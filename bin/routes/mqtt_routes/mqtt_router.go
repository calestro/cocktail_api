package mqttroutes

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func MqttConnect() {
	var broker = "192.168.0.103"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("api_client")
	client := mqtt.NewClient(opts)
	sub(client)
}

func sub(client mqtt.Client) {
	topic := "api/v1"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic %s", topic)
}
