package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestGenerateCommand_implement(t *testing.T) {
	var _ cli.Command = &GenerateCommand{}
}
