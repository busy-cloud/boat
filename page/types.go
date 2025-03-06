package page

import "github.com/busy-cloud/boat/smart"

type Table struct {
	smart.Table

	SearchUrl    string `json:"search_url,omitempty"`
	SearchAction string `json:"search_action,omitempty"`
}

type Form struct {
	smart.Form

	DataUrl      string `json:"data_url,omitempty"`
	DataAction   string `json:"data_action,omitempty"`
	SubmitUrl    string `json:"submit_url,omitempty"`
	SubmitAction string `json:"submit_action,omitempty"`
}

type Info struct {
	smart.Info

	DataUrl    string `json:"data_url,omitempty"`
	DataAction string `json:"data_action,omitempty"`
}

type Page struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Template string `json:"template"`
	Content  any    `json:"content"` //只能是以下的类型
	//Table    *Table `json:"table,omitempty"`
	//Form     *Form  `json:"form,omitempty"`
	//Info     *Info  `json:"info,omitempty"`
}
