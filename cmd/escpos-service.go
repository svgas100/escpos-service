package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	command "github.com/svgas100/escpos-service/internal/commands"
	"github.com/svgas100/escpos-service/internal/grpc/escposgrpcservice"
	usb "github.com/svgas100/escpos-service/internal/usb"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	escposgrpcservice.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *escposgrpcservice.HelloRequest) (*escposgrpcservice.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &escposgrpcservice.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func publishGRPCEndpoint() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	escposgrpcservice.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	usb.TestUsb()

	/*testESCPOSCommand(command.PrintAndLineFeedESCPOSCommand{})
	testESCPOSCommand(command.PrintAndCarriageReturnESCPOSCommand{})
	testESCPOSCommand(command.SelectDefaultLineSpacingESCPOSCommand{})
	err := testESCPOSParameterisedCommand(command.PrintAndFeedPaperESCPOSCommand{Param_n: 252})
	if err != nil {
		fmt.Println(err)
	}*/
	publishGRPCEndpoint()
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
