package auth

import (
	"github.com/wayn3h0/go-restful"
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
	recorder := httptest.NewRecorder()
	credentials := []*Credential{
		&Credential{
			Username: "foo",
			Password: "bar",
		},
		&Credential{
			Username: "username",
			Password: "password",
		},
	}

	api := restful.NewAPI()
	m := New(credentials...)
	m.AuthenticationFunc = func(username, password string) bool {
		return username == "x" && password == "y"
	}
	api.Use(m)
	api.Register("/", &Resource{"OK"})
	request, _ := http.NewRequest("GET", "/", nil)
	api.ServeHTTP(recorder, request)
	if recorder.Code != 401 {
		t.Fatalf("response should be 401, but got %d", recorder.Code)
	}

	recorder = httptest.NewRecorder()
	request.SetBasicAuth("username", "password")
	api.ServeHTTP(recorder, request)
	if recorder.Code == 401 {
		t.Fatal("response should not be 401")
	}

	recorder = httptest.NewRecorder()
	request.SetBasicAuth("x", "y")
	api.ServeHTTP(recorder, request)
	if recorder.Code == 401 {
		t.Fatal("response should not be 401")
	}
}
