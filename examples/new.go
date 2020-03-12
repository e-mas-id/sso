package main

import (
	"fmt"
	"github.com/e-mas-id/sso/pkg/client"
)

func main() {
	// Create new client with your own client name.
	c := client.New("skrap")

	// Set environment.
	c.SetEnv(client.Dev)

	// (optional) Set e-mas domain.
	c.SetDomain("http://localhost:11029")

	// (optional) Set debug.
	c.SetDebug(true)

	// Request client credential (client_id and client_secret)
	c, err := c.RequestCredential()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Login with username & password.
	customer, err := c.Login("axel.oktavian@orori.com", "123456")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Login with access token.
	customer, err = c.Login()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(c)
	fmt.Println(customer)
}
