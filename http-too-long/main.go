package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/cybozu-go/log"
	"github.com/cybozu-go/well"
)

func main() {
	flag.Parse()
	err := well.LogConfig{}.Apply()
	if err != nil {
		log.ErrorExit(err)
	}

	s := &well.HTTPServer{
		Server: &http.Server{
			Addr: ":8888",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(10 * time.Second)
				w.WriteHeader(http.StatusOK)
			}),
		},
	}
	err = s.ListenAndServe()
	if err != nil {
		log.ErrorExit(err)
	}
	well.Stop()
	err = well.Wait()
	if err != nil && !well.IsSignaled(err) {
		log.ErrorExit(err)
	}
}
