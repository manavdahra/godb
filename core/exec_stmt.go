package core

import (
	"fmt"
)

func (c *Cli) execInsertStmt(stmt *Statement) CmdStatus {
	row := stmt.Row
	rowBytes, err := row.Serialize()
	if err != nil {
		return ExecFail
	}
	c.Table.Rows = append(c.Table.Rows, rowBytes)
	return ExecSuccess
}

func (c *Cli) execSelectStmt(stmt *Statement) CmdStatus {
	for _, rowBytes := range c.Table.Rows {
		row := Row{}
		if err := row.Deserialize(rowBytes); err != nil {
			return ExecFail
		}
		c.printResult(fmt.Sprintf("(%d %s %s)\n", row.ID, row.UserName, row.Email))
	}
	return ExecSuccess
}
