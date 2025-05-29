package app

type App struct {
	Id          string `json:"id"`
	Name        string `json:"name"`                  //插件名
	Version     string `json:"version,omitempty"`     //版本号 SEMVER v0.0.0
	Icon        string `json:"icon,omitempty"`        //图标
	Description string `json:"description,omitempty"` //说明
	Type        string `json:"type,omitempty"`        //类型
	Author      string `json:"author,omitempty"`
	Email       string `json:"email,omitempty"`
	Homepage    string `json:"homepage,omitempty"`

	Menus []*Menu `json:"menus,omitempty"` //菜单项
	Pages string  `json:"pages,omitempty"` //模板页面目录，支持通配符

	//前端文件
	Static string `json:"static,omitempty"` //静态目录

	//可执行文件
	Executable   string   `json:"executable,omitempty"` //可执行文件
	Arguments    []string `json:"arguments,omitempty"`  //参数
	Dependencies []string `json:"dependencies,omitempty"`

	//代理
	ApiUrl     string `json:"api_url,omitempty"`
	UnixSocket string `json:"unix_socket,omitempty"`

	//内部插件
	Internal bool `json:"internal,-"`
}
