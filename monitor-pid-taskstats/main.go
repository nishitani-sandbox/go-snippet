package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/cybozu-go/log"
	"github.com/cybozu-go/well"
	"github.com/mdlayher/taskstats"
	"github.com/mitchellh/go-ps"
)

const (
	interval = 10 // seconds
)

func main() {
	flag.Parse()
	err := well.LogConfig{}.Apply()
	if err != nil {
		log.ErrorExit(err)
	}

	if len(flag.Args()) == 0 {
		log.ErrorExit(errors.New("please specify PID"))
	}
	pid, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		log.ErrorExit(err)
	}

	proc, err := ps.FindProcess(pid)
	if err != nil {
		log.ErrorExit(err)
	}

	c, err := taskstats.New()
	if err != nil {
		log.ErrorExit(err)
	}

	prev := &taskstats.Stats{}
	well.Go(func(ctx context.Context) error {
		t := time.NewTicker(interval * time.Second)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-t.C:
				cur, err := c.PID(pid)
				if err != nil {
					return err
				}
				log.Info(fmt.Sprintf("stats of %d: %s", pid, proc.Executable()), map[string]interface{}{
					"user_cpu_time":          (cur.UserCPUTime.Nanoseconds() - prev.UserCPUTime.Nanoseconds()) / int64(time.Millisecond),
					"system_cpu_time":        (cur.SystemCPUTime.Nanoseconds() - prev.SystemCPUTime.Nanoseconds()) / int64(time.Millisecond),
					"minor_page_faults":      cur.MinorPageFaults - prev.MinorPageFaults,
					"major_page_faults":      cur.MajorPageFaults - prev.MajorPageFaults,
					"cpu_delay_count":        cur.CPUDelayCount - prev.CPUDelayCount,
					"cpu_delay":              (cur.CPUDelay.Nanoseconds() - prev.CPUDelay.Nanoseconds()) / int64(time.Millisecond),
					"block_io_delay_count":   cur.BlockIODelayCount - prev.BlockIODelayCount,
					"block_io_delay":         (cur.BlockIODelay.Nanoseconds() - prev.BlockIODelay.Nanoseconds()) / int64(time.Millisecond),
					"swap_in_delay_count":    cur.SwapInDelayCount - prev.SwapInDelayCount,
					"swap_in_delay":          (cur.SwapInDelay.Nanoseconds() - prev.SwapInDelay.Nanoseconds()) / int64(time.Millisecond),
					"free_pages_delay_count": cur.FreePagesDelayCount - prev.FreePagesDelayCount,
					"free_pages_delay":       (cur.FreePagesDelay.Nanoseconds() - prev.FreePagesDelay.Nanoseconds()) / int64(time.Millisecond),
				})
				prev = cur
			}
		}
		return nil
	})

	well.Stop()
	err = well.Wait()
	if err != nil && !well.IsSignaled(err) {
		log.ErrorExit(err)
	}
}
