package world

import (
	"EpiCraft/game/entities"
	"github.com/mewmew/tmx"
	"strconv"
)

const (
	zonePath = "../../assets/zones/"
)

var (
	zoneLoader = make(map[string]*Zone)
)

// GetZones return the zones loaded.
func GetZones() []*Zone {
	zones := make([]*Zone, len(zoneLoader))
	for _, zone := range zoneLoader {
		zones = append(zones, zone)
	}
	return zones
}

// LoadZone returns the zone asked.
// If the zone has already been loaded, it is a light operation.
func LoadZone(zoneName string) (zone *Zone, err error) {
	zone, ok := zoneLoader[zoneName]
	if ok {
		return
	}
	tmxMap, err := tmx.Open(zonePath + zoneName + ".tmx")
	if err != nil {
		return
	}
	if zone, err = zoneFromTMXMap(tmxMap); err != nil {
		return
	}
	zoneLoader[zoneName] = zone
	return
}

// zoneFromTMXMap returns a new instantiated zone from a TMX (Tiled Map XML) map.
func zoneFromTMXMap(tmxMap *tmx.Map) (zone *Zone, err error) {
	zone = &Zone{}
	zone.Name = tmxMap.Layers[0].Name
	zone.XMax = tmxMap.Width
	zone.YMax = tmxMap.Height
	zone.Cells = entities.NewCells(zone.XMax, zone.YMax)
	for y := 0; y < tmxMap.Height; y++ {
		for x := 0; x < tmxMap.Width; x++ {
			cell := entities.NewCell(x, y, entities.Type(tmxMap.Layers[0].GetGID(x, y)))
			switch cell.Type {
			case entities.EntityPlayerSpawner:
				cellType, err := strconv.Atoi(tmxMap.Tilesets[2].TilesInfo[0].Properties[0].Value)
				if err != nil {
					return nil, err
				}
				cell.Entities = append(cell.Entities, entities.NewEntity(0, 0, entities.EntityPlayerSpawner))
				cell.Type = entities.Type(cellType)
				zone.spawn = cell
			}
			zone.Cells[y][x] = cell
		}
	}
	return
}
