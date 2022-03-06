package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/mipsmonsta/chunky/util"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":50051")

	if err != nil {
		e := fmt.Errorf("TCP Error: %w", err)
		fmt.Println(e)
		os.Exit(1)
	}
	file_dir := "./tmp"
	util.CreateFilesDir(&file_dir)
	//set the hook function for when the file is received at server end
	util.DefaultSomething = util.DoSomething{Function:hookProcessReceivedFile}

	s := grpc.NewServer()
	util.RegisterServices(s)
	log.Fatal(s.Serve(l))
}

func hookProcessReceivedFile(filePath string){
	fmt.Println(filePath)
}