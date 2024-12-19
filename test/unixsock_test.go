package test

import (
	"encoding/json"
	"net/url"
	"testing"
)

func TestUnixSockUrl(t *testing.T) {
	//"unix:///temp/boat.sock"
	uu := "C:\\Users\\Administrator\\AppData\\Local\\Temp/boat.sock"
	u, e := url.Parse(uu)
	t.Log(u, e)
	text, _ := json.Marshal(u)
	t.Log(string(text))
}
