package smart

type Column struct {
	Key      string
	Label    string
	Keyword  bool
	Sortable bool
	Filter   map[string]any
	Date     bool
	Ellipsis bool
	Break    bool
}

type Table []Column
