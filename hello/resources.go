package hello

import (
	"fmt"
	"net/http"
	"time"
)

type Cpu struct {
}

func (*Cpu) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				return
			default:
			}
		}
	}()
	time.Sleep(time.Second * 90)
	close(done)
	fmt.Fprintf(w, "done")
}
