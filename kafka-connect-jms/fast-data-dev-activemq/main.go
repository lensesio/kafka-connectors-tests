/*
A simple program to listen on an ActiveMQ (or other STOMP server)
queue or topic and print out the messages.
*/
package main

import (
	"flag"
	"log"
	"os"

	"github.com/go-stomp/stomp"
)

var serverAddr = flag.String("server", "localhost:61613", "STOMP server endpoint")
var originName = flag.String("origin", "/topic/connect_topic", "origin queue or topic")
var logFilename = flag.String("log", "", "file to write output (and logs), stdout if left empty")

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

	conn, err := stomp.Dial("tcp", *serverAddr)
	if err != nil {
		log.Fatalf("Error connecting to server: %v\n", err)
	}
	defer conn.Disconnect()

	sub, err := conn.Subscribe(*originName, stomp.AckClient)
	if err != nil {
		log.Fatalf("Error subscribing to origin: %v\n", err)
	}

	for msg := range sub.C {
		log.Println(string(msg.Body))
	}

}
