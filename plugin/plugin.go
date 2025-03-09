package plugin

import (
	"context"
	"github.com/busy-cloud/boat/lib"
	"github.com/busy-cloud/boat/log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var plugins lib.Map[Plugin]

type Plugin struct {
	Name         string   `json:"name"`
	Version      string   `json:"version"`
	Description  string   `json:"description,omitempty"`
	Type         string   `json:"type,omitempty"`
	Executable   string   `json:"executable,omitempty"`
	Arguments    []string `json:"arguments,omitempty"`
	Dependencies []string `json:"dependencies,omitempty"`
	Author       string   `json:"author,omitempty"`
	Email        string   `json:"email,omitempty"`
	Homepage     string   `json:"homepage,omitempty"`
	Socket       string   `json:"socket,omitempty"`

	dir     string
	proxy   *httputil.ReverseProxy
	process *os.Process
}

func (p *Plugin) Open() (err error) {

	if p.Executable != "" {
		attr := &os.ProcAttr{}
		attr.Dir = p.dir
		attr.Env = os.Environ()
		//TODO 可以添加环境变量
		attr.Files = append(attr.Files, os.Stdin, os.Stdout, os.Stderr)
		p.process, err = os.StartProcess(p.Executable, p.Arguments, attr)
		if err != nil {
			return err
		}
		log.Info("plugin start ", p.Name, ", pid ", p.process.Pid)
	}

	return nil
}

func (p *Plugin) Close() error {
	if p.process != nil {
		return p.process.Kill()
		//return p.process.Release()
	}
	return nil
}

func (p *Plugin) Proxy() {
	u, err := url.Parse(p.Socket)
	if err != nil {
		return
	}

	//创建反向代理
	p.proxy = &httputil.ReverseProxy{
		//Director: func(req *http.Request) {
		//	req.URL.Scheme = u.Scheme
		//	req.URL.Host = u.Host
		//	//设置User-Agent
		//	if _, ok := req.Header["User-Agent"]; !ok {
		//		// explicitly disable User-Agent so it's not set to default value
		//		req.Header.Set("User-Agent", "")
		//	}
		//},
	}

	//如果是unix
	if u.Scheme == "unix" {
		p.proxy.Transport = &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return net.Dial("unix", p.Socket)
			},
		}
	}
}
