package page

import "github.com/busy-cloud/boat/smart"

type Table struct {
	smart.Table

	SearchUrl  string `json:"search_url,omitempty"`
	SearchFunc string `json:"search_func ,omitempty"`
}

type Form struct {
	smart.Form

	LoadUrl    string `json:"load_url,omitempty"`
	LoadFunc   string `json:"load_func ,omitempty"`
	SubmitUrl  string `json:"submit_url,omitempty"`
	SubmitFunc string `json:"submit_func ,omitempty"`
}

type Info struct {
	smart.Info

	LoadUrl  string `json:"load_url,omitempty"`
	LoadFunc string `json:"load_func ,omitempty"`
}

type Content struct {
	Template string `json:"template"` //模板 table form info chart

	*Table
	*Form
	*Info
}

type Page struct {
	Name    string    `json:"name"`
	Content []Content `json:"content"` //只能是以下的类型
}
