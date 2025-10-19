package at_commands

import (
	"fmt"
	"strings"
)

type ATCommand int

const (
	ManufacturerIdentification ATCommand = iota
	ModelIdentification
	RevisionIdentification
	ProductSerialNumberIdentification
	Clock
)

type CommandInfo struct {
	BaseCmd string
	Modes   CommandMode
}

var atCommands = map[ATCommand]CommandInfo{
	ManufacturerIdentification:        {BaseCmd: "AT+CGMI", Modes: ModeExecute},
	ModelIdentification:               {BaseCmd: "AT+CGMM", Modes: ModeExecute},
	RevisionIdentification:            {BaseCmd: "AT+CGMR", Modes: ModeExecute},
	ProductSerialNumberIdentification: {BaseCmd: "AT+CGSN", Modes: ModeExecute},
	// Clock docstring ?
	Clock: {
		BaseCmd: "AT+CCLK",
		Modes:   ModeRead | ModeTest | ModeWrite,
	},
}

func (at ATCommand) String() string {
	return atCommands[at].BaseCmd
}

func (c CommandInfo) Build(mode CommandMode, args ...string) (string, error) {
	if c.Modes&mode == 0 {
		return "", fmt.Errorf("%s does not support mode %v", c.BaseCmd, mode)
	}
	switch mode {
	case ModeExecute:
		return c.BaseCmd, nil
	case ModeRead:
		return c.BaseCmd + "?", nil
	case ModeTest:
		return c.BaseCmd + "=?", nil
	case ModeWrite:
		if len(args) == 0 {
			return "", fmt.Errorf("write mode requires args")
		}
		return c.BaseCmd + "=" + strings.Join(args, ","), nil
	default:
		return "", fmt.Errorf("unknown mode")
	}
}
