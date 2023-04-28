package commands

/** Select default line spacing Command **/
type SelectDefaultLineSpacingESCPOSCommand struct {
}

func (command SelectDefaultLineSpacingESCPOSCommand) ESCPOSCommandName() string {
	return "Select default line spacing"
}

func (command SelectDefaultLineSpacingESCPOSCommand) ESCPOSCommandSequence() []uint8 {
	return []uint8{0x1B, 0x32}
}
