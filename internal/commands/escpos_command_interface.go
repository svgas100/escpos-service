package commands

import (
	"regexp"
)

/* ESC POS Command as specifiec by "https://reference.epson-biz.com/modules/ref_escpos/index.php". */
type ESCPOSCommand interface {
	/** Command name. */
	ESCPOSCommandName() string

	/** Command sequence. */
	ESCPOSCommandSequence() []uint8
}

/* ESC POS command which uses parameters inside its command buffer. */
type ESCPOSParameterisedCommand interface {
	ESCPOSCommand // extend ESCPOSCommand

	/**
	* Validation patterns for the provided input.
	* The index defines for which position the supplied regex pattern
	* should be applied in order to verify provided commands.
	 */
	ESCPOSParameterValidation() []ValidationPattern
}

/* Validation pattern for ESCPOSParameterisedCommand */
type ValidationPattern struct {
	Index       int
	Regex       *regexp.Regexp
	Description string
}
