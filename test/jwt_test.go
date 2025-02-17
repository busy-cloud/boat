package test

import (
	"github.com/busy-cloud/boat/web"
	"testing"
)

func TestJwt(t *testing.T) {
	d, e := web.JwtGenerate("id", false)
	t.Log(d, e)
}
