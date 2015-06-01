package proxysh

import (
	"testing"
	"fmt"
)

const (
	username = "your-username-here"
	password = "your-password-here"
)

func TestGetServerLoad(t *testing.T) {
	fmt.Println("proxysh.GetServerLoad()")
	zz := New(username, password)
	d, _ := zz.GetServerLoad()
	for _, s := range d.ServerList {
		fmt.Println("Location:    ", s.Location)
		fmt.Println("\tAddress:     ", s.Address)
		fmt.Println("\tServer Load: ", s.ServerLoad)
		fmt.Println()
	}
}