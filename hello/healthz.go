package hello

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	started = time.Now()
)

type HealthzHandler struct {
}

func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	duration := time.Now().Sub(started)
	if duration.Seconds() > 20 {
		log.Println("Healthz 健康检查失败")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
		return
	} else {
		log.Println("Healthz 健康检查成功")
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}
