package main_test

import (
	"net/http"
	"testing"
)

func Test_check_connection(t *testing.T) {
	resp, err := http.Get("http://gin-example-go:8080/")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp.Status)
}
