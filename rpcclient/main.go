package main

import (
	"flag"
	"fmt"
	"net/rpc"
)

func readRemoteTxtFile(relativeFilePath string) {
	rpcClient, err := rpc.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result string
	err = rpcClient.Call("RPCService.ReadTxtFile", relativeFilePath, &result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully get file content...\n")
	fmt.Println(result)
}

func main(){
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Please specify file path to read...")
		return
	}
	readRemoteTxtFile(args[0])
}