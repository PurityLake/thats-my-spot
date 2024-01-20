package data

import (
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
)

func GetFile(filename string) ([]byte, error) {
	if runtime.GOOS == "js" {
		resp, err := http.Get(filename)
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
		return byteValue, nil
	} else {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		defer file.Close()

		byteValue, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		return byteValue, nil
	}
}
