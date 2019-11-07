package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/cybozu-go/log"
	"github.com/cybozu-go/well"
)

const (
	DataSize = 1 << 30
)

type CustomReader struct {
	count int
}

func (r *CustomReader) Read(p []byte) (int, error) {
	fmt.Println("CustomReader: read")
	ds := len(p)
	if r.count+ds > DataSize {
		ds = DataSize - r.count
	}
	n := copy(p, strings.Repeat("a", ds))
	r.count += n
	fmt.Printf("CustomReader: count: %d\n", r.count)
	time.Sleep(1 * time.Second)
	return n, nil
}

func main() {
	flag.Parse()
	err := well.LogConfig{}.Apply()
	if err != nil {
		log.ErrorExit(err)
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	srv := &well.HTTPServer{
		Server: &http.Server{
			Addr: ":8888",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				go func() {
					<-r.Context().Done()
					log.Info("server: context canceled", nil)
					r.Body.Close()
				}()

				_, err := io.CopyN(ioutil.Discard, r.Body, DataSize)
				if err != nil {
					log.Info("server: read request body failed", map[string]interface{}{
						log.FnError: err,
					})
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				w.WriteHeader(http.StatusOK)
			}),
		},
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.ErrorExit(err)
	}

	well.Go(func(ctx context.Context) error {
		resp, err := http.Post("http://localhost:8888", "text/plain", &CustomReader{})
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		log.Info("client: got response", map[string]interface{}{
			"status": resp.StatusCode,
		})
		return nil
	})

	well.Stop()
	if err = well.Wait(); err != nil {
		log.ErrorExit(err)
	}
}
