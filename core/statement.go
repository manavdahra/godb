package core

type StatementType string

const (
	INSERT StatementType = "INSERT"
	SELECT StatementType = "SELECT"
)

func (st StatementType) String() string {
	return string(st)
}

type Statement struct {
	StatementType
	Row
}
