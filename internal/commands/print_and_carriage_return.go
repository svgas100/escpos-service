package commands

/** Print and carriage return Command **/
type PrintAndCarriageReturnESCPOSCommand struct {
}

func (command PrintAndCarriageReturnESCPOSCommand) ESCPOSCommandName() string {
	return "Print and carriage return"
}

func (command PrintAndCarriageReturnESCPOSCommand) ESCPOSCommandSequence() []uint8 {
	return []uint8{0x0D}
}
