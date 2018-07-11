package main

import (
	"net/http"
	"time"
	"fmt"
	
	"github.com/facebookgo/grace/gracehttp"
)

func main() {
	go func() {
		ticker := time.NewTicker(time.Second * 10)
		for {
			select {
			case <-ticker.C:
				fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

			}
		}
	}()

	gracehttp.Serve(
		&http.Server{Addr: ":5001", Handler: newGraceHandler()},
	)
}

func newGraceHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	return mux
}
