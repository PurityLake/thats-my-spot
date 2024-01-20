package mapreader

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/PurityLake/thatsmyspot/data"
)

type (
	JsonObject = map[string]interface{}
	JsonArray  = []interface{}
)

func ReadJson(filename string) (map[string]interface{}, error) {
	if runtime.GOOS == "js" {
		resp, err := http.Get("assets/maps/tiled/map0.json")
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		defer resp.Body.Close()

		byteValue, err := io.ReadAll(resp.Body)
		if err != nil {
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
	jsonFile, err := os.Open(filename)
	http.Get("asserts/maps/tiled/map0.json")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
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
	}
	properties["layers"] = *data.NewProperty("layers", "array", allLayers)

	return properties
}

func ParseTilesetData(object JsonObject) map[string]data.Property {
	properties := make(map[string]data.Property)
	tiles := object["tiles"].(JsonArray)
	tileList := make([]data.Property, len(tiles))
	for i, tile := range tiles {
		tileObj := tile.(JsonObject)
		tileProps := make(map[string]data.Property)
		for _, value := range tileObj["properties"].(JsonArray) {
			prop := value.(JsonObject)
			tileProps[prop["name"].(string)] = *data.NewProperty(prop["name"].(string), prop["type"].(string), prop["value"])
		}
		tileList[i] = *data.NewProperty(fmt.Sprintf("%d", i), "map", tileProps)
	}
	properties["properties"] = *data.NewProperty("properties", "array", tileList)
	return properties
}
