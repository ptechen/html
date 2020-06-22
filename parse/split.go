package parse

import (
	"encoding/json"
	"strings"
)

type Split struct {
	Key    string `json:"key" yaml:"key"`
	Index  int    `json:"index" yaml:"index"`
	Enable bool   `json:"enable" yaml:"enable"`
}

func (params *Split) split(text string) string {
	if params != nil {
		if params.Enable {
			if params.Key == "\\n" {
				text = strings.Split(text, "\n")[params.Index]
			} else if params.Key == "\\t" {
				text = strings.Split(text, "\t")[params.Index]
			} else {
				text = strings.Split(text, params.Key)[params.Index]
			}
		} else {
			dataBytes := make([]byte, 0)
			if params.Key == "\\n" {
				dataBytes, _ = json.Marshal(strings.Split(text, "\n"))
				text = string(dataBytes)
			} else if params.Key == "\\t" {
				dataBytes, _ = json.Marshal(strings.Split(text, "\t"))

			} else {
				dataBytes, _ = json.Marshal(strings.Split(text, params.Key))
			}
			text = string(dataBytes)
		}
	}
	return text
}