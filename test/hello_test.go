package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	api "go-demo/src/api"
)

func TestHelloWorld(t *testing.T) {

	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	os.Setenv("name", "Thaynara")
	api.HelloWorld(rr, req)
	os.Unsetenv("name")

	expected := `{"message":"Hello world Thaynara!"}`
	actual, err := ioutil.ReadAll(rr.Body)

	if err != nil {
		t.Fatal(err)
	}

	if string(actual) != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}
