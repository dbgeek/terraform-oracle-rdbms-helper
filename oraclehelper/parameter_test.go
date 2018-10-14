package oraclehelper

import (
	"testing"
)

func TestParameterService(t *testing.T) {
	tf := ResourceParameter{
		Name: "undo_retention",
	}
	param, _ := c.ParameterService.Read(tf)

	if "undo_retention" != string(param.Name) {
		t.Errorf("%v; want %v\n", param.Name, "undo_retention")
	}

}

func TestSetParameter(t *testing.T) {
	tf := ResourceParameter{
		Name:  "undo_retention",
		Value: "400",
	}
	c.ParameterService.SetParameter(tf)
	param, _ := c.ParameterService.Read(tf)

	if param.Value != "400" {
		t.Errorf("%v; want %v\n", param.Value, "800")
	}
}

func TestResetParameter(t *testing.T) {
	tf := ResourceParameter{
		Name:  "undo_retention",
		Value: "400",
	}
	c.ParameterService.SetParameter(tf)
	c.ParameterService.ResetParameter(tf)
	param, _ := c.ParameterService.Read(tf)

	if param.Value == "400" {
		t.Errorf("%v; want %v\n", param.Value, "800")
	}

}
