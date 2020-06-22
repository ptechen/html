package parse

import "strings"

type Replace struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

func (params *Replace) replace(text string) string {
	if params != nil {
		if params.Before == "\\n" {
			text = strings.ReplaceAll(text, "\n", params.After)
		} else if params.Before == "\\t" {
			text = strings.ReplaceAll(text, "\t", params.After)
		} else {
			text = strings.ReplaceAll(text, params.Before, params.After)
		}
	}
	return text
}

func replaces(params[]*Replace, text string) string {
	if len(params) > 0 {
		for i := 0; i < len(params); i++ {
			rep := params[i]
			text = rep.replace(text)
		}
	}
	return text
}