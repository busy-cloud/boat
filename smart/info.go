package smart

type Item struct {
	Key     string         `json:"key"`
	Type    string         `json:"type,omitempty"` //type object array
	Label   string         `json:"label"`
	Span    int            `json:"span,omitempty"`
	Options map[string]any `json:"options,omitempty"`
	Default any            `json:"default,omitempty"`
}

type Info []Item
