package parse

import "strings"

func deletes(params []string, text string) string {
	if len(params) > 0 {
		for i := 0; i < len(params); i++ {
			curDelete := params[i]
			if curDelete == "\\n" {
				text = strings.ReplaceAll(text, "\n", "")
			} else if curDelete == "\\t" {
				text = strings.ReplaceAll(text, "\t", "")
			} else {
				text = strings.ReplaceAll(text, curDelete, "")
			}
		}
	}
	return text
}