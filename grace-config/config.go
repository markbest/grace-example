package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/markbest/grace-example/grace-config/config"
)

func main() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR1)
	go func() {
		for {
			<-s
			g.ReloadConfig()
			log.Println("Reloaded config")
			fmt.Println(g.Config())
		}
	}()

	fmt.Println(g.Config())
	select {}
}
