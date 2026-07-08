package main

import (
	"flag"
	"fmt"

	"github.com/anan112pcmec/Burung-backend-2/watcher_app"
)

func main() {
	var rebootcass bool
	flag.BoolVar(&rebootcass, "cassreboot", false, "contoh penggunaan")
	flag.Parse()

	if rebootcass {
		fmt.Println("akan restart cassandra")
	} else {
		fmt.Println("cassandra gaakan di restart")
	}
	watcher_app.Run(rebootcass)
}
