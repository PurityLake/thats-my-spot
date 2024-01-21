package mapreader

import (
	"image/color"

	"github.com/PurityLake/thatsmyspot/data"
)

func ParseMapData(object JsonObject) map[string]data.Property {
	properties := make(map[string]data.Property)
	bgcolor, ok := object["backgroundcolor"].(string)
	if !ok {
		properties["backgroundcolor"] = *data.NewProperty("backgroundcolor", "color", color.RGBA{0, 0, 0, 255})
	} else {
		colorValue := data.HexToRGBA(bgcolor)
		properties["backgroundcolor"] = *data.NewProperty("backgroundcolor", "color", colorValue)
	}
	layers := object["layers"].(JsonArray)
	allLayers := make([]data.Property, len(layers))
	for i, layer := range layers {
		layerObj := layer.(JsonObject)
		dataArr := layerObj["data"].(JsonArray)
		length := len(dataArr)
		mapList := make([]int, length)
		for j := 0; j < length; j++ {
			mapList[j] = int(dataArr[j].(float64))
		}
		allLayers[i] = *data.NewProperty(layerObj["name"].(string), "array", mapList)
		for _, mapProps := range layerObj["properties"].(JsonArray) {
			prop := mapProps.(JsonObject)
			properties[prop["name"].(string)] = *data.NewProperty(prop["name"].(string), prop["type"].(string), prop["value"])
		}
	}
	properties["layers"] = *data.NewProperty("layers", "array", allLayers)

	return properties
}
