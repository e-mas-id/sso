# E-mas SSO 

Golang library to help implementing [E-mas](https://www.e-mas.com/) Single Sign-On. Similar with "Login with Facebook" thing.

## Features
* Client credential request
* Login with username & password
* Login with access token
* Login authorization
* Logout

_More will be coming soon..._

## Installation
Get the library with the usual `go get`.
1. `go get github.com/e-mas-id/sso`
2. That's it.

## Usage
```go
package main

import "github.com/e-mas-id/sso/pkg/client"

func main() {
  // Create new client with your own client name.
  c := client.New("your_client_name")

  // (optional) Set client name (if you have more than 1 clients).
  c.SetClientName("your_client_name_2")

  // (optional) Set client id and secret if you have.
  c.SetClientCredential("cId", "cSecret")

  // (optional) Set environment.
  c.SetEnv(client.Dev)

  // (optional) Set e-mas domain.
  c.SetDomain("http://localhost:123")

  // (optional) Set debug.
  c.SetDebug(true)

  // Request client credential (client_id and client_secret)
  c, err := c.RequestCredential()
  if err != nil {
    // do something
    return
  }

  // Login with username & password.
  customer, err := c.Login("e-mas@mail.com", "123456")
  if err != nil {
    // do something
    return
  }

  // Authorize SSO login (for the first time only).
  err = c.Authorize()
  if err != nil {
    // do something
    return
  }

  // Login with access token. Now return customer data.
  customer, err = c.Login()
  if err != nil {
    // do something
    return
  }

  // Logout SSO login.
  err = c.Logout()
  if err != nil {
    // do something
    return
  }
}
```
_For more detail, please go to [godoc](https://godoc.org/github.com/e-mas-id/sso/pkg/client)._

## Flow
General flow of E-mas SSO.
1. **Requesting Client Credentials**

   This should be done for the first time when registering new client name. You will get `client_id` and `client_secret` that will be used for all your client's SSO activity. So, keep it somewhere so you don't have to request it every initiation.
  
2. **Login with Username & Password**

   You don't need to do this if you already access `token`. This will return the customer's access `token` that will be used for most of customer's SSO activity. You won't get customer data (email, name, etc) if the customer is not authorized yet. (go to step 3 for authorization). 
   > Each customer will have different access `token` that will be expired after period of time. <br>To generate a new one, just login with username & password again.
   
3. **Authorization**

   This should be done once for every customer login using your client. Authorization requires customer access `token`. So, you need login with username & password first then use the `token` here. After customer is authorized, you will get customer data after login.
   
4. **Login with Token**

   This is an alternative way for login. You can login with username & password in this step too. Both login will return the same data. Just choose which one you prefer.
   
5. **Using Customer Data**

   After getting customer data, just do whatever you want with it. For example, registering new user to your app using E-mas customer data.
   
6. **Logout**

   Well, it's a logout. Invalidate current customer `token` so it can't be used anymore.
   

For more information and support please contact dev@e-mas.com.
