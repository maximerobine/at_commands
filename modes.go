package at_commands

type CommandMode uint8

const (
	ModeExecute CommandMode = 1 << iota
	ModeRead
	ModeTest
	ModeWrite
)
