package oraclehelper

import (
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
)

var (
	granularity = []string{"ALL", "AUTO", "DEFAULT", "PARTITION", "GLOBAL", "GLOBAL AND PARTITION", "SUBPARTITION"}
)

func TestStatsService(t *testing.T) {

	v := granularity[acctest.RandIntRange(0, 6)]

	resourceStats := ResourceStats{
		Pname: "GRANULARITY",
		Pvalu: v,
	}
	err := c.StatsService.SetGlobalPre(resourceStats)
	if err != nil {
		log.Fatalf("failed to stats, errormsg: %v\n", err)
	}

	globalGranularity, err := c.StatsService.ReadGlobalPre(resourceStats)
	if err != nil {
		log.Fatalf("failed to stats, errormsg: %v\n", err)
	}

	if resourceStats.Pvalu != globalGranularity.Pvalu {
		t.Errorf("got %s; want %s\n", globalGranularity.Pvalu, resourceStats.Pvalu)
	}

}
