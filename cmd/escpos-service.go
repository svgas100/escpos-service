package main

import (
	"errors"
	"fmt"
	"strconv"

	command "github.com/svgas100/escpos-service/internal/commands"
)

func main() {

	testESCPOSCommand(command.PrintAndLineFeedESCPOSCommand{})
	testESCPOSCommand(command.PrintAndCarriageReturnESCPOSCommand{})
	testESCPOSCommand(command.SelectDefaultLineSpacingESCPOSCommand{})
	err := testESCPOSParameterisedCommand(command.PrintAndFeedPaperESCPOSCommand{Param_n: 252})
	if err != nil {
		fmt.Println(err)
	}
}

func testESCPOSCommand(command command.ESCPOSCommand) {
	fmt.Println("###################################")
	fmt.Println(command.ESCPOSCommandName())
	fmt.Println(command.ESCPOSCommandSequence())
}

func testESCPOSParameterisedCommand(command command.ESCPOSParameterisedCommand) error {
	testESCPOSCommand(command)
	validationPatterns := command.ESCPOSParameterValidation()

	commandSequence := command.ESCPOSCommandSequence()
	commandLength := len(commandSequence)
	// validate arguments!

	for i := 0; i < len(validationPatterns); i++ {
		validationPattern := validationPatterns[i]
		if validationPattern.Index >= commandLength {
			return errors.New(fmt.Sprintf("Command sequence length does not match the supplied validation pattern! Max command index '%v', required validation index '%v'.", commandLength-1, validationPattern.Index))
		}

		argumentString := strconv.FormatInt(int64(commandSequence[validationPattern.Index]), 10)
		result := validationPattern.Regex.MatchString(argumentString)
		if !result {
			return errors.New(fmt.Sprintf("Argument '%v' at position %v does not match argument constraint! %v", argumentString, validationPattern.Index, validationPattern.Description))
		}
	}
	return nil
}
