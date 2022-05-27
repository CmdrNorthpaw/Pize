package main

import (
	"flag"
	"log"
	"net"
	"strings"
)

const (
	ok = "ok"
	Error = "error"
)

var logger = log.Default()

func main() {
	dummy := NewDrone()

	port := flag.String("p", "8889", "Set the port to listen on. Defaults to 8889")
	
	// Resolve the UDP address and set up a listener for it
	address, err := net.ResolveUDPAddr("udp4", *port)
	if err != nil {
		panic(err)
	}
	
	connection, err := net.ListenUDP("udp4", address)
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	buffer := make([]byte, 1024)

	// The success and failiure functions detail what to do in the event that the command succeeds or fails.
	// The word "ok" or "error" will be sent accross the UDP connection to mimic the drone
	// and more detailed information will be outputted to the server logs.
	success := func(message string) {
		connection.WriteToUDP([]byte(ok), address)
		logger.Println("SUCCESS: " + message)
	}

	failiure := func(message string) {
		connection.WriteToUDP([]byte(Error), address)
		logger.Println("ERROR: " + message)
	}

	// Continually checks for new data on the UDP connection.
	for {
		n, _, _ := connection.ReadFromUDP(buffer)
		recieved := strings.Split(string(buffer[0:n-1]), " ")

		if command := Commands[recieved[0]]; command == nil {
			failiure("Unknown command " + recieved[0])
		} else {
			// The Tello has to have "SDK mode" enabled before it will accept any other commands.
			// We check if SDK mode is enabled, and if it isn't, the command has to be the SDK activation command
			// or the command will fail.
			if !dummy.SDKMode {
				if recieved[0] == "command" {
					dummy.SDKMode = true
					success("Drone SDK mode activated")
				} else {
					failiure("Drone is not in SDK mode yet. Please send the command \"command\" to enable it")
				}
			// Now that we know SDK mode is enabled, we check the result of the command.
			// If it succeeds, we send an "ok" and log the success message,
			// otherwise we send "error" and log the failiure.
			} else if success, message := command(dummy, recieved[1:]...); success {
				connection.WriteToUDP([]byte(ok), address)
				logger.Println("INFO: " + message)
			} else {
				failiure(message)
			}
		}
	}
}