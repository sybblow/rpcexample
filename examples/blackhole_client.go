package main

import (
	"io"
	"log"
	"net/rpc/jsonrpc"
	"os"
	"sync"

	"github.com/haisum/rpcexample"
)

type stdioWrapper struct {
	in  io.ReadCloser
	out io.WriteCloser
}

func NewStdioWrapper() *stdioWrapper {
	return &stdioWrapper{os.Stdin, os.Stdout}
}

func (wr *stdioWrapper) Write(p []byte) (int, error) {
	return wr.out.Write(p)
}

func (wr *stdioWrapper) Read(p []byte) (int, error) {
	return wr.in.Read(p)
}

func (wr *stdioWrapper) Close() error {
	wr.out.Close()
	return wr.in.Close()
}

func main() {
	client := jsonrpc.NewClient(NewStdioWrapper())
	args := &rpcexample.Args{
		A: 2,
		B: 3,
	}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var result rpcexample.Result
			err := client.Call("Arith.Multiply", args, &result)
			if err != nil {
				log.Printf("error in Arith", err)
				return
			}
			log.Printf("%d*%d=%d\n", args.A, args.B, result)
		}()
	}
	wg.Wait()
}
