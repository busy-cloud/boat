package page

import "github.com/busy-cloud/boat/smart"

type Page struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Template string `json:"template"` //模板 table form info chart

	*smart.Table
	*smart.Form
	*smart.Info

	SearchUrl  string         `json:"search_url,omitempty"` //查询URL
	SearchFunc string         `json:"search_func ,omitempty"`
	LoadUrl    string         `json:"load_url,omitempty"` //加载URL
	LoadFunc   string         `json:"load_func ,omitempty"`
	SubmitUrl  string         `json:"submit_url,omitempty"` //提交URL
	SubmitFunc string         `json:"submit_func ,omitempty"`
	Params     map[string]any `json:"params,omitempty"` //页面参数
	ParamsFunc string         `json:"params_func,omitempty"`

	Children []*Page `json:"children,omitempty"`
}
