package selector_test

import (
	"fmt"
	"github.com/ptechen/config"
	"github.com/ptechen/html/parse"
	"io/ioutil"
	"testing"
)

func TestSelector(t *testing.T) {
	yml := map[string]*parse.FilterParams{}
	conf := &config.Config{}
	conf.YAML("test.yml", &yml)
	htmlBytes, err := ioutil.ReadFile("test.html")
	if err != nil {
		t.Error(err)
	}
	html := string(htmlBytes)
	res, err := parse.HtmlParse(html, yml)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}
