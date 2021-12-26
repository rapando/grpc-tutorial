package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/rapando/grpc-tutorial/machine"
	"google.golang.org/grpc"
)

type MachineServer struct {
	machine.UnimplementedMachineServer
}

// Execute
// this function runs the set of instructions provided
func (s *MachineServer) Execute(ctx context.Context, instruction *machine.Instruction) (result *machine.Result, err error) {
	log.Println("instruction : ", instruction)
	return &machine.Result{
		Output: fmt.Sprintln("operator: ", instruction.Operator, "operand: ", instruction.Operand),
	}, nil

}

func main() {
	log.Printf("initializing grpc on port 9000")

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("unable to start listener : %v", err)
	}

	grpcServer := grpc.NewServer()
	machine.RegisterMachineServer(grpcServer, &MachineServer{})
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("unable to start server because %v", err)
	}
}
