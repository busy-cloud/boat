package smart

type Filter struct {
	Text  string `json:"text,omitempty"`
	Value any    `json:"value,omitempty"`
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
	Link     string   `json:"link,omitempty"`
}

type Operator struct {
	Icon     string `json:"icon,omitempty"`
	Label    string `json:"label,omitempty"`
	Title    string `json:"title,omitempty"`
	Link     string `json:"link,omitempty"`
	Action   string `json:"action,omitempty"`
	Confirm  string `json:"confirm,omitempty"`
	External bool   `json:"external,omitempty"`
}

type Button struct {
	Icon   string `json:"icon,omitempty"`
	Label  string `json:"label,omitempty"`
	Title  string `json:"title,omitempty"`
	Link   string `json:"link,omitempty"`
	Action string `json:"action,omitempty"`
}

type Table struct {
	Buttons   []Button   `json:"buttons,omitempty"`
	Columns   []Column   `json:"columns,omitempty"`
	Operators []Operator `json:"operators,omitempty"`
}
