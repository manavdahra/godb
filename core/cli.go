package core

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

type CmdStatus string

const (
	MetaSuccess      CmdStatus = "META_SUCCESS"
	MetaExit         CmdStatus = "META_EXIT"
	MetaUnrecognized CmdStatus = "META_UNRECOGNIZED"

	ExecSuccess CmdStatus = "EXEC_SUCCESS"
	ExecFail    CmdStatus = "EXEC_FAIL"

	PrepSuccess      CmdStatus = "PREP_SUCCESS"
	PrepUnrecognized CmdStatus = "PREP_UNRECOGNIZED"
	PrepSyntaxError  CmdStatus = "PREP_SYNTAX_ERROR"
)

type Cli struct {
	in  io.Reader
	out io.Writer
	Table
}

func NewCli(in io.Reader, out io.Writer) *Cli {
	return &Cli{Table: NewTable(), in: in, out: out}
}

func (c *Cli) Run() {
	for {
		c.printPrompt()
		rdr := bufio.NewReader(c.in)
		b, _, err := rdr.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		cmd := string(b)
		switch c.ProcessCmd(cmd) {
		case MetaSuccess, ExecSuccess:
			c.printResult("Executed\n")
		case MetaUnrecognized:
			c.printResult(fmt.Sprintf("Unrecognized command '%s'.\n", cmd))
		case PrepSyntaxError:
			c.printResult(fmt.Sprintf("Syntax error. Could not prepare the statement.\n"))
		case PrepUnrecognized:
			c.printResult(fmt.Sprintf("Unrecognized keyword at the start of '%s'.\n", cmd))
		case ExecFail:
			c.printResult(fmt.Sprintf("failed to execute statement\n"))
		case MetaExit:
			c.printResult(fmt.Sprintf("exiting\n"))
			return
		}
	}
}

func (c *Cli) ProcessCmd(cmd string) CmdStatus {
	if strings.HasPrefix(cmd, ".") {
		return c.execMetaCmd(cmd)
	}
	var stmt Statement
	res := c.prepareStmt(cmd, &stmt)
	if res == PrepSuccess {
		return c.execStmt(&stmt)
	}
	return res
}

func (c *Cli) prepareStmt(cmd string, stmt *Statement) CmdStatus {
	switch {
	case strings.HasPrefix(strings.ToUpper(cmd), INSERT.String()):
		if _, err := fmt.Sscanf(cmd, "insert %d %s %s", &stmt.ID, &stmt.UserName, &stmt.Email); err != nil {
			return PrepSyntaxError
		}
		stmt.StatementType = INSERT
	case strings.HasPrefix(strings.ToUpper(cmd), SELECT.String()):
		stmt.StatementType = SELECT
	default:
		return PrepUnrecognized
	}
	return PrepSuccess
}

func (c *Cli) execMetaCmd(cmd string) CmdStatus {
	if strings.EqualFold(cmd, ".exit") {
		return MetaExit
	} else {
		return MetaUnrecognized
	}
}

func (c *Cli) execStmt(stmt *Statement) CmdStatus {
	switch stmt.StatementType {
	case INSERT:
		return c.execInsertStmt(stmt)
	case SELECT:
		return c.execSelectStmt(stmt)
	default:
		return ExecFail
	}
}

func (c *Cli) printPrompt() {
	if _, err := c.out.Write([]byte("db > ")); err != nil {
		log.Fatal(err)
	}
}

func (c *Cli) printResult(res string) {
	if _, err := c.out.Write([]byte(res)); err != nil {
		log.Fatal(err)
	}
}
