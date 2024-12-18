package boat

type Application struct {
	Name         string   `json:"name"`
	Version      string   `json:"version"`
	Description  string   `json:"description,omitempty"`
	Type         string   `json:"type,omitempty"`
	Executable   string   `json:"executable,omitempty"`
	Dependencies []string `json:"dependencies,omitempty"`
	Author       string   `json:"author,omitempty"`
	Email        string   `json:"email,omitempty"`
	Homepage     string   `json:"homepage,omitempty"`
}
