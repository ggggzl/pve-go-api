package common

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

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
