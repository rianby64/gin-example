package main_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var host string

func init() {
	err := godotenv.Load()
	if err == nil {
		host = os.Getenv("HOST")
	}
}

func Test_check_connection(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("http://%s:8080/", host))

	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp.Status)
}
