package app

import (
	"context"
	"github.com/busy-cloud/boat/log"
	"github.com/busy-cloud/boat/web"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Plugin struct {
	Id          string `json:"id"`
	Name        string `json:"name"`                  //插件名
	Version     string `json:"version,omitempty"`     //版本号 SEMVER v0.0.0
	Icon        string `json:"icon,omitempty"`        //图标
	Description string `json:"description,omitempty"` //说明
	Type        string `json:"type,omitempty"`        //类型
	Author      string `json:"author,omitempty"`
	Email       string `json:"email,omitempty"`
	Homepage    string `json:"homepage,omitempty"`

	//前端文件
	Static string `json:"static,omitempty"` //静态目录
	static http.FileSystem

	//可执行文件
	Executable   string   `json:"executable,omitempty"` //可执行文件
	Arguments    []string `json:"arguments,omitempty"`  //参数
	Dependencies []string `json:"dependencies,omitempty"`
	process      *os.Process

	//代理
	ApiUrl     string `json:"api_url,omitempty"`
	UnixSocket string `json:"unix_socket,omitempty"`
	proxy      *httputil.ReverseProxy

	//前端
	Menus []string `json:"menus,omitempty"`
	Pages []string `json:"pages,omitempty"` //模板页面目录

	dir string
}

func (p *Plugin) Open() (err error) {

	//启动子进程
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

	//前端页面
	if p.Static != "" {
		//p.static = http.Dir(p.Static)
		//注册前端
		web.StaticDir(p.Static, "/plugin/"+p.Id+"/", "", "index.html")
	}

	//接口代理
	if p.ApiUrl != "" {
		u, err := url.Parse(p.ApiUrl)
		if err != nil {
			return err
		}
		p.proxy = httputil.NewSingleHostReverseProxy(u)
	}
	//UnixSocket方式（速度更快）
	if p.UnixSocket != "" {
		p.proxy = &httputil.ReverseProxy{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
					return net.Dial("unix", p.UnixSocket)
				},
			},
		}
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

func (p *Plugin) ServeApi(ctx *gin.Context) {
	if p.proxy == nil {
		return
	}

	ctx.Abort()
	p.proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
