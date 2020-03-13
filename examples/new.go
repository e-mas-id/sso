package main

import (
	"fmt"
	"github.com/e-mas-id/sso/pkg/client"
)

func main() {
	// Create new client with your own client name.
	c := client.New("randomName")

	// Set client name (if you have more than 1 clients).
	c.SetClientName("randomName2")

	// Set client id and secret if you have.
	c.SetClientCredential("cId", "cSecret")

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
	customer, err := c.Login("e-mas@mail.com", "123456")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Authorize SSO login (for the first time only).
	err = c.Authorize()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Login with access token. Now return customer data.
	customer, err = c.Login()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Logout SSO login.
	err = c.Logout()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(c)
	fmt.Println(customer)
}
