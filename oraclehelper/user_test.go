package oraclehelper

import (
	"testing"
)

func TestUserService(t *testing.T) {

	c.UserService.CreateUser(ResourceUser{Username: "TEST01"})

	c.UserService.ModifyUser(ResourceUser{Username: "TEST01", DefaultTablespace: "SYSTEM"})

	user, _ := c.UserService.ReadUser(ResourceUser{Username: "TEST01"})

	if "SYSTEM" != user.DefaultTablespace {
		t.Errorf("%v; want %v\n", user.DefaultTablespace, "SYSTEM")
	}

	c.UserService.DropUser(ResourceUser{Username: "TEST01"})

}
