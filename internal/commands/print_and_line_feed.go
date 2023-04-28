package commands

/** Print and line feed Command **/
type PrintAndLineFeedESCPOSCommand struct {
}

func (command PrintAndLineFeedESCPOSCommand) ESCPOSCommandName() string {
	return "Print and line feed"
}

func (command PrintAndLineFeedESCPOSCommand) ESCPOSCommandSequence() []uint8 {
	return []uint8{0x0A}
}
