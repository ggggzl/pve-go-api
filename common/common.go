package common

import (
	"encoding/json"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//GetInfo read YAML from file
func GetInfo(filePath string, content interface{}) error {

	rawContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(rawContent, content)
	if err != nil {
		return err
	}

	return err
}

//GetRawData extract the "data" field from the Proxmox API query response and return raw JSON field for structure-specific unmarshalling
func GetRawData(b []byte) ([]*json.RawMessage, error) {
	var rawContent map[string]*json.RawMessage
	var rawData []*json.RawMessage

	if err := json.Unmarshal(b, &rawContent); err != nil {
		return nil, err
	}
	for key, value := range rawContent {
		if key == "data" {
			if err := json.Unmarshal(*value, &rawData); err != nil {
				return nil, err
			}
		}
	}

	return rawData, nil
}
