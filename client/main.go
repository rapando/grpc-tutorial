package main

import (
	"context"
	"log"
	"time"

	"github.com/rapando/grpc-tutorial/machine"
	"google.golang.org/grpc"
)

func main() {
	log.Printf("starting client to server on port 9000")

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial("localhost:9000", opts...)
	if err != nil {
		log.Fatalf("unable to connect to server because %v", err)
	}
	defer conn.Close()

	client := machine.NewMachineClient(conn)

	// try to execute
	instructions := []*machine.Instruction{
		{Operator: "PUSH", Operand: 1},
		{Operator: "PUSH", Operand: 2},
		{Operator: "ADD"},
		{Operator: "PUSH", Operand: 3},
		{Operator: "DIV"},
		{Operator: "PUSH", Operand: 4},
		{Operator: "MUL"},
		{Operator: "PUSH", Operand: 5},
		{Operator: "PUSH", Operand: 6},
		{Operator: "SUB"},
	}

	runExecute(client, instructions)
}

func runExecute(client machine.MachineClient, instructions []*machine.Instruction) {
	// sends a stream of calls
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	for _, instruction := range instructions {
		log.Printf("sending : operator : %s, operand : %v", instruction.Operator, instruction.Operand)
		res, err := client.Execute(ctx, instruction)
		if err != nil {
			log.Fatalf("send failed : %v", err)
		}

		log.Printf("received : %v", res.GetOutput())
		log.Printf("\n\n")
	}
}
