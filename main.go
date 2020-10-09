package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	gonanoid "github.com/matoous/go-nanoid"
)

func sleepWindows() {
	cmd := exec.Command("cmd.exe", "/C", "shutdown", "/h")

	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}

var powerHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	command := string(msg.Payload())
	if command == "OFF" {
		log.Println("POWER OFF")
		sleepWindows()
	}
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected to MQTT server")
}

var disconnectHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Println(fmt.Sprintf("Disconnected from MQTT server: %s", err))
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	err := godotenv.Load("gohass-mqtt-winclient.env")
	if err != nil {
		log.Fatal("Error: cannot load gohass-mqtt-winclient.env file")
	}
	mqttURI := os.Getenv("MQTT_URI")
	mqttUname := os.Getenv("MQTT_USERNAME")
	mqttPass := os.Getenv("MQTT_PASSWORD")
	mqttTopic := os.Getenv("MQTT_TOPIC")
	if mqttURI == "" {
		log.Fatal("Error: .env file does not contain \"MQTT_URI\"")
	}
	if mqttUname == "" {
		log.Fatal("Error: .env file does not contain \"MQTT_USERNAME\"")
	}
	if mqttPass == "" {
		log.Fatal("Error: .env file does not contain \"MQTT_PASSWORD\"")
	}
	if mqttTopic == "" {
		log.Fatal("Error: .env file does not contain \"MQTT_TOPIC\"")
	}
	nanoid, err := gonanoid.Nanoid()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: cannot generate client ID: %s", err))
	}
	clientid := fmt.Sprintf("gohass-mqtt-winclient-%s", nanoid)
	opts := mqtt.NewClientOptions().AddBroker(mqttURI).SetClientID(clientid)
	opts.SetUsername(mqttUname).SetPassword(mqttPass)
	opts.SetOnConnectHandler(connectHandler).SetConnectionLostHandler(disconnectHandler)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetPingTimeout(1 * time.Second)

	mqtt.ERROR = log.New(os.Stdout, "ERR:", 0)
	mqtt.CRITICAL = log.New(os.Stdout, "CRIT:", 0)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	log.Println("Starting MQTT client")

	// Subscribe
	if token := client.Subscribe(mqttTopic, 0, powerHandler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	<-c
}
