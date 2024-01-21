package mapreader

import (
	"encoding/json"
	"fmt"
	"image/color"
	"log"

	"github.com/PurityLake/thatsmyspot/data"
)

type (
	JsonObject = map[string]interface{}
	JsonArray  = []interface{}
)

func ReadJson(filename string) (map[string]interface{}, error) {
	byteValue, err := data.GetFile(filename)
	if err != nil {
		fmt.Println("Error reading file")
		log.Fatal(err)
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

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
