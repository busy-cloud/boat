package smart

type Chart struct {
	Type    string           `json:"type"`
	Options map[string]any   `json:"options,omitempty"`
	Legend  bool             `json:"legend,omitempty"`
	Tooltip bool             `json:"tooltip,omitempty"`
	Time    bool             `json:"time,omitempty"`
	Radar   map[string]int64 `json:"radar,omitempty"`
}
