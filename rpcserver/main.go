package main

import(
	"fmt"
	"net"
	"net/rpc"
	"os"
	"path/filepath"
	"errors"
)

type RPCService struct {
	WorkFolder string
}

func (rpcService *RPCService) ReadTxtFile(relativeFilePath string, result *string) error {
	fullPathOfFile := filepath.Join(rpcService.WorkFolder, relativeFilePath)
	fmt.Println("Request to read file :", relativeFilePath)
	if _, err := os.Stat(fullPathOfFile); os.IsNotExist(err) {
		fmt.Println(err)
		return errors.New(relativeFilePath +" doesn't exist!")
	}
	file, err := os.Open(fullPathOfFile)
	if err != nil {
		fmt.Println( err )
		return errors.New("Failed to open file " + relativeFilePath)
	}
	defer file.Close()
	stat , err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return errors.New("Failed to get file size" + relativeFilePath)
	}
	byteBuf := make([]byte, stat.Size())
	_, err = file.Read(byteBuf)
	if err != nil {
		fmt.Println(err)
		return errors.New("Failed to read file content: " + relativeFilePath)
	}
	(*result) = string(byteBuf)
	fmt.Println("Successfully processed request Reading file ", relativeFilePath)
	return nil
}

func RegisterRPCService() {
	rpcService := RPCService{ WorkFolder :"d:\\code\\go\\scrape"}
	rpc.Register(&rpcService)
	ln, err := net.Listen("tcp", "localhost:9999")	
	if err != nil {
		fmt.Println(err)
		return
	}
	for{
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main(){
	go RegisterRPCService()
	var temp string
	fmt.Println("RPC Service is ready. Press ANY Key to interrupt...")
	fmt.Scanln(&temp)
}