package hello

import (
	"flag"
	"fmt"
	"net/http"
)

const (
	defaultStartUpParam = "default"
)

var (
	param1 = flag.String("param1", defaultStartUpParam, "param1 to hello world")
	param2 = flag.String("param2", defaultStartUpParam, "param2 to hello world")
)

func init() {
	flag.Parse()
}

type PrintStartParam struct {
}

func (*PrintStartParam) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("start up params:     param1 = %s and param2 = %s ", *param1, *param2)
	fmt.Println()
	fmt.Fprintf(w, "start up params:     param1 = %s and param2 = %s ", *param1, *param2)
}
