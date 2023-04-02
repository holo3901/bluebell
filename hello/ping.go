package hello

import (
	"fmt"
	"net"
	"net/http"
)

type Pong struct {
}

func (*Pong) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var ip string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}
	fmt.Printf("ip = %s", ip)
	fmt.Println()
	fmt.Fprintf(w, "ip = %s", ip)
}
