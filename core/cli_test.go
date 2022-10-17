package core_test

import (
	"godb/core"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCli_ProcessCmd(t *testing.T) {
	cli := core.NewCli(os.Stdin, os.Stdout)
	type testCase struct {
		input  string
		output core.CmdStatus
	}
	tcs := map[string]testCase{
		"select_none": {input: "select", output: core.ExecSuccess},
		"insert":      {input: "insert 1 manav manav.dahra@gmail.com", output: core.ExecSuccess},
		"select_some": {input: "select", output: core.ExecSuccess},
		"exit":        {input: ".exit", output: core.MetaExit},
	}
	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, cli.ProcessCmd(tc.input), tc.output)
		})
	}
}
