package hello

import (
	"log"
	"net/http"
)

type HealthHandler struct {
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Health 健康检查成功")
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}
