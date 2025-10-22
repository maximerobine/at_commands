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
	CharacterSet
	PIN
	SubscriberNumber
	SignalQuality
	Operator
	NetworkRegistration
	PreferredOperatorList
	ReadOperatorNames
	ErrorVerbosity
	MessageFormat
	ListMessage
	SendMessage
)

var atCommands = map[ATCommand]string{
	ManufacturerIdentification:        "AT+CGMI",
	ModelIdentification:               "AT+CGMM",
	RevisionIdentification:            "AT+CGMR",
	ProductSerialNumberIdentification: "AT+CGSN",
	Clock:                             "AT+CCLK",
	CharacterSet:                      "AT+CSCS",
	PIN:                               "AT+CPIN",
	SubscriberNumber:                  "AT+CNUM",
	SignalQuality:                     "AT+CSQ",
	Operator:                          "AT+COPS",
	NetworkRegistration:               "AT+CREG",
	PreferredOperatorList:             "AT+CPOL",
	ReadOperatorNames:                 "AT+COPN",
	ErrorVerbosity:                    "AT+CMEE",
	MessageFormat:                     "AT+CMGF",
	ListMessage:                       "AT+CMGL",
	SendMessage:                       "AT+CMGS",
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
