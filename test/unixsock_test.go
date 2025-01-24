package test

import (
	"github.com/busy-cloud/boat/mqtt"
	"net"
	"net/url"
	"os"
	"testing"
	"time"
)

func parseUrl(u string, t *testing.T) {
	uu, e := url.Parse(u)
	if e != nil {
		t.Error(e)
		return
	}
	t.Log(uu.Scheme, "->", uu.Host, "->", uu.Path)
}

func TestUnixSockUrl(t *testing.T) {
	u1 := "unix:///temp/boat.sock"
	u2 := "unix:///" + url.PathEscape("C:/temp/boat.sock")
	for _, u := range []string{u1, u2} {
		parseUrl(u, t)
	}
}

func Test2(t *testing.T) {
	l, e := net.Listen("unix", os.TempDir()+"/boat.sock")
	if e != nil {
		t.Error(e)
		return
	}
	defer l.Close()

	go Connect(t)
	go Connect(t)
	go Connect(t)
	go Connect(t)
	go Connect(t)

	buf := make([]byte, 1024)
	for i := 0; i < 5; i++ {
		c, e := l.Accept()
		if e != nil {
			t.Error(e)
		}
		n, e := c.Read(buf)
		if e != nil {
			t.Error(e)
		}
		_, _ = c.Write(buf[:n])
		_ = c.Close()
	}
}

func Connect(t *testing.T) {
	c, e := net.Dial("unix", "/"+os.TempDir()+"/boat.sock")
	if e != nil {
		t.Error(e)
		return
	}
	defer c.Close()

	_, e = c.Write([]byte("hello"))
	if e != nil {
		t.Error(e)
		return
	}
	buf := make([]byte, 1024)
	n, e := c.Read(buf)
	if e != nil {
		t.Error(e)
	}
	t.Log(string(buf[:n]))
}

func TestMqtt(t *testing.T) {
	err := mqtt.Startup()
	if err != nil {
		t.Error(err)
	}
	defer mqtt.Shutdown()

	time.Sleep(time.Second)
	tk := mqtt.Publish("test", []byte("hello world"))
	//tk.Wait()
	e := tk.Error()
	if e != nil {
		t.Error(e)
	}

	mqtt.Client.Disconnect(0)
}
