package main

import (
	"godb/core"
	"os"
)

func main() {
	core.NewCli(os.Stdin, os.Stdout).Run()
}
