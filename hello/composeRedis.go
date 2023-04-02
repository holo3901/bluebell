package hello

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
	"time"
)

type RedisHandler struct {
}

func (*RedisHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	val := r.FormValue("value")
	if key == "" {
		key = "key"
	}
	conn, err := getConn()
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "conn redis failed")
		return
	}
	defer conn.Close()

	if key != "" && val != "" {
		_, err = conn.Do("set", key, val)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "redis set failed")
			return
		}
		fmt.Fprintf(w, "redis set success")
	}

	if key != "" && val == "" {
		res, err := conn.Do("get", key)
		if err != nil {
			log.Println(err)
			fmt.Fprintf(w, "redis get failed")
			return
		}
		if res != nil {
			fmt.Fprintf(w, "redis get success value : %v", string(res.([]byte)))
		} else {
			fmt.Fprintf(w, "redis get value is empty ")
		}
	}
}

const (
	ip   = "redis"
	port = "6379"
	pwd  = "123456"
)

func getConn() (redis.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	c, err := redis.DialContext(ctx, "tcp", ip+":"+port)
	if err != nil {
		return nil, err
	}
	if _, err := c.Do("auth", pwd); err != nil {
		c.Close()
		return nil, err
	}
	return c, nil
}
