package oraclehelper

import (
	"testing"
)

func TestUserService(t *testing.T) {
	quota := make(map[string]string)
	quota["USERS"] = "unlimited"
	quota["SYSTEM"] = "10m"
	quota["SYSAUX"] = "10G"
	c.UserService.CreateUser(ResourceUser{Username: "TEST01"})
	c.ProfileService.CreateProfile(ResourceProfile{Profile: "PP01"})
	user, _ := c.UserService.ReadUser(ResourceUser{Username: "TEST01"})
	if user.Profile != "DEFAULT" {
		t.Errorf("want: %s, gott: %s", "DEFAULT", user.Profile)
	}
	c.UserService.ModifyUser(ResourceUser{Username: "TEST01", DefaultTablespace: "SYSTEM", Quota: quota, Profile: "PP01"})

	user, _ = c.UserService.ReadUser(ResourceUser{Username: "TEST01"})
	if "SYSTEM" != user.DefaultTablespace {
		t.Errorf("%v; want %v\n", user.DefaultTablespace, "SYSTEM")
	}
	if user.Quota["SYSTEM"] != "10M" {
		t.Errorf("gott: %s; want:%s\n", user.Quota["SYSTEM"], "10M")
	}
	if user.Quota["USERS"] != "unlimited" {
		t.Errorf("%s; want %s\n", user.Quota["USERS"], "unlimited")
	}
	if user.Quota["SYSAUX"] != "10G" {
		t.Errorf("%s; want %s\n", user.Quota["SYSAUX"], "10G")
	}
	if user.Profile != "PP01" {
		t.Errorf("want: %s, gott: %s", "PP01", user.Profile)
	}

	c.UserService.DropUser(ResourceUser{Username: "TEST01"})
	c.ProfileService.DeleteProfile(ResourceProfile{Profile: "PP01"})
}
