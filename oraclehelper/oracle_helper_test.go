package oraclehelper

import (
	"github.com/kelseyhightower/envconfig"

	"log"
	"testing"
)

var (
	cfg = Cfg{}

	c = &Client{}
)

func init() {
	if err := envconfig.Process("tf_ora_helper", &cfg); err != nil {
		log.Fatalf("failed to parse config: %v\n", err)
	}
	c = NewClient(cfg)
}
func TestDBConnection(t *testing.T) {
	var got string
	want := "foo"

	rows, err := c.DBClient.Query("select 'foo' as foo from dual")
	if err != nil {
		t.Errorf("error: %g", err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&got)
	}
	if want != got {
		t.Errorf("%v; want %v\n", got, want)
	}

}
