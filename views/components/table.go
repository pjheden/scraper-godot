package components

type Table struct {
	Columns []TableColumn
	Rows    [][]string
}

type TableColumn struct {
	// Name of the column, will be displayed in the header.
	//
	// Example: "Email address"
	Name string

	// If the column is a href, the href will be used to create a link.
	IsHref bool

	// ConvertValueFunc is a function that can be used to convert the value of the cell.
	//
	// For example if the columns is boolean, we might want to show an icon instead of TRUE/FALSE
	ConvertValueFunc func(string) string
}

type TableRow struct {
	// ID of the specific row.
	//
	// Example: 1
	ID int

	// Cells of the row, will be displayed in the table.
	Cells []string
}
