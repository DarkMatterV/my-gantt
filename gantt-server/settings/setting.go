package settings

import (
	"encoding/json"
	"os"
)

type setting struct {
	DB database `json:"database" yaml:"database"`
}

type database struct {
	Addr string `json:"addr" yaml:"addr"`
}

var Setting *setting

func InitSetting(file string) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &Setting)
	return err
}
