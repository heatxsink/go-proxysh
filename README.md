go-proxysh
==========

A golang module that hits the api over at [proxy.sh](https://proxy.sh) for vpn server load.

Install
-------
You'll need to have golang installed.

	$ go get github.com/heatxsink/go-proxysh

Usage
-----

```go

import (
  "fmt"
  "github.com/heatxsink/go-proxysh"
)

const (
  username = "your-username-here"
  password = "your-password-here"
)

func main() {
  pp := proxysh.New(username, password)
  data, err := pp.GetServerLoad()
  if err != nil {
    fmt.Println("Error: ", err)
  }
  for _, s := range data.ServerList {
    fmt.Println("Location:    ", s.Location)
    fmt.Println("Address:     ", s.Address)
    fmt.Println("Server Load: ", s.ServerLoad)
    fmt.Println()
  }

```