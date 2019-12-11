package main

import (
	"context"
	"errors"
	"flag"
	"os"

	"github.com/cybozu-go/log"
	"github.com/cybozu-go/well"
	"github.com/fsnotify/fsnotify"
)

func main() {
	flag.Parse()
	err := well.LogConfig{}.Apply()
	if err != nil {
		log.ErrorExit(err)
	}

	if len(flag.Args()) == 0 {
		log.ErrorExit(errors.New("please specify a path to a monitored dir"))
	}

	fi, err := os.Stat(flag.Args()[0])
	if err != nil {
		log.ErrorExit(err)
	}

	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.ErrorExit(err)
	}
	defer w.Close()

	err = w.Add(fi.Name())
	if err != nil {
		log.ErrorExit(err)
	}

	well.Go(func(ctx context.Context) error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case e := <-w.Events:
				log.Info("received event", map[string]interface{}{
					"path": e.Name,
					"op":   e.Op,
				})
			case err := <-w.Errors:
				return err
			}
		}
	})

	well.Stop()
	err = well.Wait()
	if err != nil && !well.IsSignaled(err) {
		log.ErrorExit(err)
	}
}
