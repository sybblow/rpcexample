package main

import (
	"encoding/json"
	"os"
	"time"
	// "github.com/sybblow/rpcexample"
)

type Foo struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func main() {
	// conn := rpcexample.NewLineWriter(os.Stdout)
	enc := json.NewEncoder(os.Stdout)
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		enc.Encode(Foo{"me", 2})
	}
}
