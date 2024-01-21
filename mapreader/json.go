package mapreader

import (
	"encoding/json"
	"fmt"
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
