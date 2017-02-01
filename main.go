package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func main() {
	config := &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{
			ssh.Password("yourpassword"),
		},
	}
	fmt.Printf("%v\n", config)
}
