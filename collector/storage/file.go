package storage

import (
	"encoding/json"
	"os"
)

func SaveJSONL(path string, v any) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	_, err = file.Write(append(data, '\n'))
	return err
}