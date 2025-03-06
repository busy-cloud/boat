package smart

type Filter struct {
	Text  string `json:"text,omitempty"`
	Value any    `json:"value,omitempty"`
}

type Action struct {
	Type       string         `json:"type,omitempty"`
	Link       string         `json:"link,omitempty"`
	LinkFunc   string         `json:"link_func,omitempty"`
	Params     map[string]any `json:"params,omitempty"`
	ParamsFunc string         `json:"params_func,omitempty"`
	Script     string         `json:"script,omitempty"`
	Page       string         `json:"page,omitempty"`
	Dialog     string         `json:"dialog,omitempty"`
	External   string         `json:"external,omitempty"`
}

type Column struct {
	Key      string   `json:"key,omitempty"`
	Label    string   `json:"label,omitempty"`
	Keyword  bool     `json:"keyword,omitempty"`
	Sortable bool     `json:"sortable,omitempty"`
	Filter   []Filter `json:"filter,omitempty"`
	Date     bool     `json:"date,omitempty"`
	Ellipsis bool     `json:"ellipsis,omitempty"`
	Break    bool     `json:"break,omitempty"`
	Action   *Action  `json:"action,omitempty"`
}

type Operator struct {
	Icon     string  `json:"icon,omitempty"`
	Label    string  `json:"label,omitempty"`
	Title    string  `json:"title,omitempty"`
	Action   *Action `json:"action,omitempty"`
	Confirm  string  `json:"confirm,omitempty"`
	External bool    `json:"external,omitempty"`
}

type Button struct {
	Icon   string  `json:"icon,omitempty"`
	Label  string  `json:"label,omitempty"`
	Title  string  `json:"title,omitempty"`
	Action *Action `json:"action,omitempty"`
}

type Table struct {
	Buttons   []Button   `json:"buttons,omitempty"`
	Columns   []Column   `json:"columns,omitempty"`
	Operators []Operator `json:"operators,omitempty"`
}
