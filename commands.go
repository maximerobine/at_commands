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

var atCommands = map[ATCommand]string{
	ManufacturerIdentification:        "AT+CGMI",
	ModelIdentification:               "AT+CGMM",
	RevisionIdentification:            "AT+CGMR",
	ProductSerialNumberIdentification: "AT+CGSN",
	Clock:                             "AT+CCLK",
}

func (at ATCommand) String() string {
	return atCommands[at]
}

func (at ATCommand) Read() string {
	return at.String() + "?"
}

func (at ATCommand) Execute() string {
	return at.String()
}

func (at ATCommand) Test() string {
	return at.String() + "=?"
}

func (at ATCommand) Write(args ...string) string {
	if len(args) == 0 {
		return at.String() + "="
	}
	return fmt.Sprintf("%s=%s", at.String(), strings.Join(args, ","))
}
