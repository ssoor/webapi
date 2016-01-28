package webapi

import (
	"io/ioutil"
	"net/http"
	"testing"
)

type BasicJsonItem struct{}

func (item BasicJsonItem) Get(values Values, request *http.Request) (int, interface{}, http.Header) {
	items := []string{"item1", "item2"}
	data := map[string][]string{"items": items}
	return 200, data, nil
}

func TestBasicJsonGet(t *testing.T) {

	item := new(BasicJsonItem)

	var api = NewJsonAPI()
	api.AddResource(item, "/items", "/bar", "/baz")
	go api.Start(3000)
	resp, err := http.Get("http://localhost:3000/items")
	if err != nil {
		t.Error(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "{\n  \"items\": [\n    \"item1\",\n    \"item2\"\n  ]\n}" {
		t.Error("Not equal.")
	}
}

type BasicByteItem struct{}

func (item BasicByteItem) Get(values Values, request *http.Request) (int, interface{}, http.Header) {
	items := "items++++++++++++items++++++++++items"
	data := []byte(items)
	return 200, data, nil
}

func TestBasicByteGet(t *testing.T) {

	item := new(BasicByteItem)

	var api = NewByteAPI()
	api.AddResource(item, "/items", "/bar", "/baz")
	go api.Start(3001)
	resp, err := http.Get("http://localhost:3001/items")
	if err != nil {
		t.Error(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "items++++++++++++items++++++++++items" {
		t.Error("Not equal.")
	}
}
