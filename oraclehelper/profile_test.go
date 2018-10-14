package oraclehelper

import (
	"testing"
)

func TestProfileService(t *testing.T) {
	profile := ResourceProfile{
		Profile: "ORA_STIG_PROFILE",
	}
	param, _ := c.ProfileService.ReadProfile(profile)

	if "15" != param["IDLE_TIME"] {
		t.Errorf("%v; want %v\n", param["IDLE_TIME"], "15")
	}
}

func TestProfileServiceCreateUpdateDropProfile(t *testing.T) {
	profile := ResourceProfile{
		Profile:      "TEST01",
		ResourceName: "IDLE_TIME",
		Limit:        "30",
	}
	c.ProfileService.CreateProfile(profile)

	before, _ := c.ProfileService.ReadProfile(profile)

	if "TEST01" != before["PROFILE"] {
		t.Errorf("%v; want %v\n", before["PROFILE"], "TEST01")
	}

	c.ProfileService.UpdateProfile(profile)

	after, _ := c.ProfileService.ReadProfile(profile)
	if after["IDLE_TIME"] == before["IDLE_TIME"] {
		t.Errorf("%v; want %v\n", before["IDLE_TIME"], after["IDLE_TIME"])
	}

	c.ProfileService.DeleteProfile(profile)
}
