package common

import (
	"bytes"

	"github.com/spf13/cobra"
)

func ExecuteCommand(root *cobra.Command, args []string) (string, error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	err := root.Execute()
	return buf.String(), err
}
