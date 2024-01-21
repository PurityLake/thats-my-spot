package mapreader

import (
	"fmt"
	"log"

	"github.com/PurityLake/thatsmyspot/data"
)

func ParseTilesetData(object JsonObject) map[string]data.Property {
	properties := make(map[string]data.Property)
	tiles, ok := object["tiles"].(JsonArray)
	if !ok {
		log.Fatal("Could not parse map layers")
	}
	tileList := make([]data.Property, len(tiles))
	for i, tile := range tiles {
		tileObj := tile.(JsonObject)
		tileProps := make(map[string]data.Property)
		t, ok := tileObj["properties"].(JsonArray)
		if !ok {
			log.Fatal("Could not parse tile properties")
		}
		for _, value := range t {
			prop := value.(JsonObject)
			tileProps[prop["name"].(string)] = *data.NewProperty(prop["name"].(string), prop["type"].(string), prop["value"])
		}
		tileList[i] = *data.NewProperty(fmt.Sprintf("%d", i), "map", tileProps)
	}
	properties["properties"] = *data.NewProperty("properties", "array", tileList)
	return properties
}
