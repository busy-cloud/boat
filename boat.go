package boat

import (
	"embed"
	_ "github.com/busy-cloud/boat/apis"
	_ "github.com/busy-cloud/boat/broker"
	"github.com/busy-cloud/boat/menu"
	_ "github.com/busy-cloud/boat/menu"
	"github.com/busy-cloud/boat/page"
	_ "github.com/busy-cloud/boat/plugin"
	_ "github.com/busy-cloud/boat/setting"
)

//go:embed pages
var pages embed.FS

//go:embed menus
var menus embed.FS

func init() {
	//注册页面
	page.EmbedFS(pages, "pages")

	//注册菜单
	menu.EmbedFS(menus, "menus")
}
