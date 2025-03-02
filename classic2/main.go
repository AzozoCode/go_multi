package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGBUS)

	select {

	case res := <-sig:
		signal.Stop(sig)
		fmt.Printf("%s recieved", res)
		os.Exit(0)
	}

}
