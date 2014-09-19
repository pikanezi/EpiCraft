package world

import (
	"fmt"
	"testing"
	"time"
)

var (
	zones = []string{"zone_0"}
)

func TestLoadZone(t *testing.T) {
	for _, zoneName := range zones {
		timer := time.Now()
		zone, err := LoadZone(zoneName)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("Loaded %s in %v\n", zoneName, time.Since(timer)))
		if len(zone.Cells[0]) != zone.XMax {
			t.Fatal("Invalid width")
		}
		if len(zone.Cells) != zone.YMax {
			t.Fatal("Invalid height")
		}
		if zone.spawn == nil {
			t.Fatal("Spawn not set")
		}
		if zoneLoader[zoneName] == nil {
			t.Fatal(fmt.Sprintf("Zone %s not in zone loader", zoneName))
		}
	}
}

func BenchmarkLoadZone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := LoadZone(zones[0])
		if err != nil {
			b.Fatal(err)
		}
	}
}
