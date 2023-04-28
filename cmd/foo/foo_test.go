package foo

import (
	"testing"

	"github.com/fiveateooate/yacctl/cmd"
	"github.com/fiveateooate/yacctl/pkg/common"
)

func Test_FooCommand(t *testing.T) {
	args := []string{"foo"}
	cmdOutput, err := common.ExecuteCommand(cmd.RootCmd, args)
	if err != nil {
		t.Errorf("%s\n", err)
	} else {
		t.Logf("output: %s", cmdOutput)
	}
}
