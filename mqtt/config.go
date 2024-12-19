package mqtt

import (
	"github.com/god-jason/boat/config"
	"github.com/god-jason/boat/lib"
	"os"
	"runtime"
)

const MODULE = "mqtt"

func init() {
	url := "mqtt://localhost:1883"
	if runtime.GOOS != "windows" {
		//使用UnixSocket速度更快
		url = "unix://" + os.TempDir() + "/boat.sock" //windows下会出问题，Win10以上虽然支持，但是不能使用绝对路径，因为盘符会被错误解析
	}
	config.Register(MODULE, "url", url)
	config.Register(MODULE, "clientId", lib.RandomString(12))
	config.Register(MODULE, "username", "")
	config.Register(MODULE, "password", "")
}
