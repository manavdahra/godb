package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type CmdStatus string

const (
	MetaSuccess      CmdStatus = "META_SUCCESS"
	MetaUnrecognized CmdStatus = "META_UNRECOGNIZED"
	PrepSuccess      CmdStatus = "PREP_SUCCESS"
	PrepUnrecognized CmdStatus = "PREP_UNRECOGNIZED"
)

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
}

func main() {
	for {
		printPrompt()
		reader := bufio.NewReader(os.Stdin)
		b, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		cmd := string(b)
		if strings.HasPrefix(cmd, ".") {
			switch execMetaCmd(cmd) {
			case MetaSuccess:
				continue
			case MetaUnrecognized:
				printMetaCmdExecFail(cmd)
				continue
			}
		}

		var stmt Statement
		switch prepareStmt(cmd, &stmt) {
		case PrepUnrecognized:
			printPrepStmtExecFail(cmd)
			continue
		}
		execStmt(stmt)
	}
}

func prepareStmt(cmd string, stmt *Statement) CmdStatus {
	cmd = strings.ToUpper(cmd)
	switch {
	case strings.HasPrefix(cmd, INSERT.String()):
		stmt.StatementType = INSERT
	case strings.HasPrefix(cmd, SELECT.String()):
		stmt.StatementType = SELECT
	default:
		return PrepUnrecognized
	}
	return PrepSuccess
}

func printMetaCmdExecFail(cmd string) {
	fmt.Printf("Unrecognized command '%s'.\n", cmd)
}

func printPrepStmtExecFail(cmd string) {
	fmt.Printf("Unrecognized keyword at the start of '%s'.\n", cmd)
}

func execMetaCmd(cmd string) CmdStatus {
	if strings.EqualFold(cmd, ".exit") {
		os.Exit(0)
		return MetaSuccess
	} else {
		return MetaUnrecognized
	}
}

func execStmt(stmt Statement) {
	switch stmt.StatementType {
	case INSERT:
		fmt.Printf("Do insert\n")
	case SELECT:
		fmt.Printf("Do select\n")
	}
}

func printPrompt() {
	fmt.Print("db > ")
}
