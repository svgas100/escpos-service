package commands

import "regexp"

/** Print and feed paper Command **/
type PrintAndFeedPaperESCPOSCommand struct {
	Param_n uint8
}

func (command PrintAndFeedPaperESCPOSCommand) ESCPOSCommandName() string {
	return "Print and feed paper"
}

func (command PrintAndFeedPaperESCPOSCommand) ESCPOSCommandSequence() []uint8 {
	return []uint8{0x1B, 0x4A, command.Param_n}
}

func (command PrintAndFeedPaperESCPOSCommand) ESCPOSParameterValidation() []ValidationPattern {
	return []ValidationPattern{
		{
			2,
			regexp.MustCompile("^\\b([01]?[0-9][0-9]?|2[0-4][0-9]|25[0-5])$"),
			"Only values from 0 - 255",
		},
	}
}
