package redis

import (
	"context"
	"strconv"
	"time"
)

func CreateLogin(userid int64, toke string) error {
	_, err := client.Set(context.Background(), getRedisKey(keyLoginZsetPF+strconv.Itoa(int(userid))), toke, 2*time.Hour).Result()
	if err != nil {
		return err
	}
	return nil
}
func GetLogin(userid int64) (token string, err error) {
	result, err := client.Get(context.Background(), getRedisKey(keyLoginZsetPF+strconv.Itoa(int(userid)))).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
