package utils

import (
	"ChangePicBed/model"
	"bytes"
	"encoding/json"
	"fmt"
)

func PrintConfig(config model.Config) error {
	var err error
	var bf []byte
	var out bytes.Buffer

	bf, err = json.Marshal(config)
	if err != nil {
		return err
	}

	err = json.Indent(&out, bf, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println("Configuration details:")
	fmt.Printf("%v\n", out.String())

	return nil
}
