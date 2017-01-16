/*
A simple program to format a couple messages in avro and send them to
a MQTT server. Also tries to read back in order to provide some kind
of verification that it works.
*/
package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/linkedin/goavro"
)

var (
	mqttServerAddr = flag.String("server", "localhost:1883", "MQTT server endpoint")
	originName     = flag.String("origin", "/topic/connect_topic", "origin queue or topic")
	logFilename    = flag.String("log", "", "file to write output (and logs), stdout if left empty")
)

func main() {
	flag.Parse()

	if len(*logFilename) > 0 {
		logFile, err := os.OpenFile(*logFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Error opening log file: %v\n", err)
		}
		defer logFile.Close()
		log.SetOutput(logFile)
	}

	schema := `
{
  "type": "record",
  "name": "measurements",
  "namespace": "com.landoop.test.mqtt",
  "fields": [
    {
      "name": "output",
      "type": "int"
    },
    {
      "name": "deviceid",
      "type":"string"
    }
  ]
}
`
	codec, err := goavro.NewCodec(schema)
	if err != nil {
		log.Println(err)
	}

	record, err := goavro.NewRecord(goavro.RecordSchema(schema))
	if err != nil {
		log.Println(err.Error())
	}

	// Send messages
	go func() {
		opts := MQTT.NewClientOptions()
		opts.AddBroker("tcp://" + *mqttServerAddr)
		opts.SetClientID("test-client-publisher")
		opts.SetCleanSession(true)

		client := MQTT.NewClient(opts)
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		var bb *bytes.Buffer
		for i := 0; i < 1; i = 0 {
			log.Println("---- doing publish ----")
			record.Set("output", int32(5))
			record.Set("deviceid", "plant1-room3-pod4")
			bb = new(bytes.Buffer)
			if err = codec.Encode(bb, record); err != nil {
				log.Fatal(err)
			}
			token := client.Publish(*originName, byte(1), false, bb.Bytes())
			token.Wait()
			time.Sleep(2 * time.Second)

			record.Set("output", int32(16))
			record.Set("deviceid", "ship5-deck2-rack3")
			bb = new(bytes.Buffer)
			if err = codec.Encode(bb, record); err != nil {
				log.Fatal(err)
			}
			token = client.Publish(*originName, byte(1), false, bb.Bytes())
			token.Wait()
			time.Sleep(2 * time.Second)
		}

		client.Disconnect(250)
	}()

	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://" + *mqttServerAddr)
	opts.SetClientID("test-client-subscriber")
	opts.SetCleanSession(true)

	receiveCount := 0
	choke := make(chan [2]string)
	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		choke <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	defer client.Disconnect(250)

	if token := client.Subscribe(*originName, byte(1), nil); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		os.Exit(1)
	}

	for receiveCount < 100 {
		incoming := <-choke
		log.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], incoming[1])
		receiveCount++
	}

}
