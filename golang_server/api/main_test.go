package api

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	go Init()
	time.Sleep(time.Second * 10)
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Errorf(fmt.Sprint(err))
	}
	if resp.StatusCode != 200 {
		t.Errorf("Ping/pong status code: %q", resp.Status)
	}
}
