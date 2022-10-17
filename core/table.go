package core

type Table struct {
	Rows [][]byte
}

func NewTable() Table {
	return Table{Rows: make([][]byte, 0)}
}
