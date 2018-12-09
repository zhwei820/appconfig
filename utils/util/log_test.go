package util

import (
	"fmt"
	"testing"
	"time"
	"sync"
	"github.com/rs/zerolog/log"

)

func TestDestroy(t *testing.T) {
	fname:="./log/example.log"
	initLogger(fname)
	refreshLogger(fname)


	ticker = time.NewTicker(time.Second * 3)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		for t := range ticker.C {
			fmt.Printf("Backup at %s\n", t)
			NewLogrotate(fname, 0, 10, 100, false, refreshLogger) //

			wg.Done()
		}
	}()
	logTest()

	wg.Wait()
	ticker.Stop()

}

func logTest() {
	t1 := time.Now()
	for ii := 0; ii < 10000000; ii++ {
		log.Info().Msg("xxxxx")
	}
	t2 := time.Since(t1).Nanoseconds()

	println(float64(t2) / float64(time.Second))

}