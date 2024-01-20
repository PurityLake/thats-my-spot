package mapreader

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
)

func ReadMap(filename string) (map[string]interface{}, error) {
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
