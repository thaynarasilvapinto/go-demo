package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	api "go-demo/src/api"
)

func setupTestEnv(t *testing.T) {
	t.Cleanup(teardownTestEnv)
	os.Setenv("name", "Thaynara")
}

func teardownTestEnv() {
	os.Unsetenv("name")
}

func TestHelloWorld(t *testing.T) {
	t.Run("TestHelloWorld", func(t *testing.T) {
		setupTestEnv(t)

		req, err := http.NewRequest("GET", "/hello", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		api.HelloWorld(rr, req)

		expected := `{"message":"Hello world Thaynara!"}`
		actual, err := ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Fatal(err)
		}

		if string(actual) != expected {
			t.Errorf("expected %s, got %s", expected, actual)
		}
	})
}
