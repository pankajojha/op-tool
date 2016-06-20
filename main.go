package main

import (
	"fmt"

	"gopkg.in/hypersleep/easyssh.v0"
)

func main() {
	
	// Create MakeConfig instance with remote username, server address and path to private key.
	ssh := &easyssh.MakeConfig{
		User:   "pearson",
		Server: "10.168.12.89",
		// Optional key or Password without either we try to contact your agent SOCKET
		Key:  "/Documents/key/epen2_nonprod30.pem",
		Port: "22",
	}

	// Call Run method with command you want to run on remote server.
	//response, err := ssh.Run("ps ax")
	fmt.Println("\n  Server checking for \n \n "+ssh.Server)

	response, err := ssh.Run("ps -eaf | grep java")
	// Handle errors
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		fmt.Println(response)
	}
}