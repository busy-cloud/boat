package app

type Base struct {
	Id          string `json:"id"`
	Icon        string `json:"icon,omitempty"`        //图标
	Name        string `json:"name"`                  //插件名
	Description string `json:"description,omitempty"` //说明
	Version     string `json:"version,omitempty"`     //版本号 SEMVER v0.0.0
	Internal    bool   `json:"internal,omitempty"`    //内部插件
}

type Menu struct {
	Name       string   `json:"name"`
	Title      string   `json:"title,omitempty"`
	NzIcon     string   `json:"nz-icon,omitempty"` //ant.design图标库
	Items      []*Entry `json:"items,omitempty"`
	Index      int      `json:"index,omitempty"`
	Privileges []string `json:"privileges,omitempty"`
	//Domain     []string `json:"domain"` //域 admin project 或 dealer等
}

type Entry struct {
	Name       string   `json:"name"`
	Title      string   `json:"title,omitempty"`
	Icon       string   `json:"icon,omitempty"`
	Url        string   `json:"url,omitempty"`
	External   bool     `json:"external,omitempty"`
	Privileges []string `json:"privileges,omitempty"`
}

type App struct {
	Base //继承基础信息

	//扩展信息
	Type     string `json:"type,omitempty"` //类型
	Author   string `json:"author,omitempty"`
	Email    string `json:"email,omitempty"`
	Homepage string `json:"homepage,omitempty"`

	//资源
	Shortcuts []*Entry `json:"shortcuts,omitempty"` //桌面快捷方式
	Menus     []*Menu  `json:"menus,omitempty"`     //菜单项
	Pages     string   `json:"pages,omitempty"`     //模板页面目录，支持通配符

	//前端文件
	Static string `json:"static,omitempty"` //静态目录

	//可执行文件
	Executable   string   `json:"executable,omitempty"` //可执行文件
	Arguments    []string `json:"arguments,omitempty"`  //参数
	Dependencies []string `json:"dependencies,omitempty"`

	//代理
	ApiUrl     string `json:"api_url,omitempty"`
	UnixSocket string `json:"unix_socket,omitempty"`
}
