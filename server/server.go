package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/mipsmonsta/chunky/util"
	"github.com/mipsmonsta/epd"
	"github.com/mipsmonsta/epd/epd_config"
	"github.com/mipsmonsta/epd/imageutil"

	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":50051")

	if err != nil {
		e := fmt.Errorf("TCP Error: %w", err)
		fmt.Println(e)
		os.Exit(1)
	}
	defer func(){ //catch panic
		if err := recover(); err != nil {
			log.Printf("Error: %s\n", err)
		}
	}()


	file_dir := "./tmp"
	util.CreateFilesDir(&file_dir)
	//set the hook function for when the file is received at server end
	util.DefaultSomething = util.DoSomething{Function:hookProcessReceivedFile}

	s := grpc.NewServer()
	util.RegisterServices(s)
	log.Fatal(s.Serve(l))
}

func hookProcessReceivedFile(filePath string){
	e := epd.Epd{
		Config: epd_config.EpdConfig{},
	}

	e.Setup()
	e.Clear()
	img, err := imageutil.OpenImage(filePath)
	if err != nil {
		panic(fmt.Sprintf("Wrong filetype %w\n", err))
	}

	e.Display(&img, epd.MODE_MONO_DITHER_ON)

	e.Sleep()
}