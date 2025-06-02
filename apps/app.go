package apps

import (
	"context"
	"github.com/busy-cloud/boat/app"
	"github.com/busy-cloud/boat/log"
	"github.com/busy-cloud/boat/web"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
)

type App struct {
	app.App

	//可执行文件
	process *os.Process

	//代理
	proxy *httputil.ReverseProxy

	opened bool
}

func (a *App) Opened() bool {
	return a.opened
}

func (a *App) Open() (err error) {

	//内部插件不用打开
	if a.Internal {
		a.opened = true
		return nil
	}

	//打开标志
	if a.opened {
		return nil
	}
	a.opened = true

	//基础路径
	dir := filepath.Join("app", a.Id)

	//启动子进程
	if a.Executable != "" {
		attr := &os.ProcAttr{}
		attr.Dir = dir
		attr.Env = os.Environ()
		//TODO 可以添加环境变量
		attr.Files = append(attr.Files, os.Stdin, os.Stdout, os.Stderr)
		a.process, err = os.StartProcess(a.Executable, a.Arguments, attr)
		if err != nil {
			return err
		}
		log.Info("plugin start ", a.Name, ", pid ", a.process.Pid)
	}

	//附件
	//assets := filepath.Join(dir, "assets")
	//a.Assets = os.DirFS(assets)

	//前端页面
	if a.Static != "" {
		//a.static = http.Dir(a.Static)
		path := filepath.Join(dir, a.Static)

		//注册前端 TODO 可能有问题， 会与代理冲突
		web.StaticDir(path, "/app/"+a.Id+"/", "", "index.html")
	}

	//接口代理
	if a.ApiUrl != "" {
		u, err := url.Parse(a.ApiUrl)
		if err != nil {
			return err
		}
		a.proxy = httputil.NewSingleHostReverseProxy(u)
	}
	//UnixSocket方式（速度更快）
	if a.UnixSocket != "" {
		a.proxy = &httputil.ReverseProxy{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
					return net.Dial("unix", a.UnixSocket)
				},
			},
		}
	}

	return nil
}

func (a *App) Close() error {
	if a.process != nil {
		return a.process.Kill()
		//return a.process.Release()
	}

	//TODO 需要从pages store中移除

	return nil
}

func (a *App) ServeApi(ctx *gin.Context) bool {
	if a.proxy == nil {
		return false
	}

	ctx.Abort()
	a.proxy.ServeHTTP(ctx.Writer, ctx.Request)

	return true
}
