package app

type Menu struct {
	Name     string      `json:"name"`
	Title    string      `json:"title,omitempty"`
	Icon     string      `json:"icon,omitempty"`
	Children []*MenuItem `json:"children,omitempty"`
	Index    int         `json:"index,omitempty"`
	//Domain     []string `json:"domain"` //域 admin project 或 dealer等
	Privileges []string `json:"privileges,omitempty"`
}

type MenuItem struct {
	//Type       string         `json:"type,omitempty"` //route 路由, web 嵌入web, window 独立弹出
	Name  string `json:"name,omitempty"`
	Title string `json:"title,omitempty"`
	Url   string `json:"url,omitempty"`
	//Query      map[string]any `json:"query,omitempty"`
	External   bool     `json:"external,omitempty"`
	Privileges []string `json:"privileges,omitempty"`
}
