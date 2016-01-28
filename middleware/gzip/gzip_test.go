package gzip

import (
	"compress/gzip"
	"github.com/wayn3h0/go-restful"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Resource struct {
	Message string
}

func (this Resource) Get(context *restful.Context) {
	context.WriteEntity(200, this)
}

func Test(t *testing.T) {
	message := "gzip test"
	resource := &Resource{message}

	request, _ := http.NewRequest("GET", "/", nil)
	request.Header.Add("Accept-Encoding", "gzip")
	recorder := httptest.NewRecorder()
	api := restful.NewAPI()
	api.Use(New())
	api.Register("/", resource)
	api.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("status code should be 200, but got %d. location: %s", recorder.Code, recorder.HeaderMap.Get("Location"))
	}

	hv := recorder.Header().Get("Content-Encoding")
	if hv != "gzip" {
		t.Errorf("error header [\"Context-Encoding\"], should be \"gzip\", but got \"%s\"", hv)
	}

	hv = recorder.Header().Get("vary")
	if hv != "Accept-Encoding" {
		t.Errorf("error header [\"Accept-Encoding\"], should be \"Accept-Encoding\", but got \"%s\"", hv)
	}

	reader, err := gzip.NewReader(recorder.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()
	body, _ := ioutil.ReadAll(reader)
	if string(body) != `{"Message":"gzip test"}` {
		t.Fatalf("error body, should be \"%s\", but got \"%s\"", message, string(body))
	}
}
