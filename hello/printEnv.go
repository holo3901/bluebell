package hello

import (
	"fmt"
	"net/http"
	"os"
)

type EnvParam struct {
}

func (*EnvParam) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	env1 := os.Getenv("env1")
	env2 := os.Getenv("env2")
	fmt.Printf("env list : env1 = %s and env2 = %s", env1, env2)
	fmt.Println()
	fmt.Fprintf(w, "env list : env1 = %s and env2 = %s", env1, env2)
}
