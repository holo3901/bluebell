package hello

import (
	"fmt"
	"net"
	"net/http"
)

type PongV2 struct {
}

func (*PongV2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	fmt.Printf("v1.0.1 ip = %s", ip)
	fmt.Println()
	fmt.Fprintf(w, "v1.0.1 ip = %s", ip)
}

type PongV2ex struct {
}

func (*PongV2ex) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	fmt.Printf("v1.0.2 ip = %s", ip)
	fmt.Println()
	fmt.Fprintf(w, "v1.0.2 ip = %s", ip)
}
