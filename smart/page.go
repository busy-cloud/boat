package smart

type Page struct {
	Name    string
	Type    string //table edit detail
	Table   string
	Columns []TableColumn
	Fields  []Field
}