package main

import (
	redis "github.com/cedriclam/go-redis-server"
)

func main() {

	server, err := redis.NewServer(redis.DefaultConfig().Host("0.0.0.0"))
	if err != nil {
		panic(err)
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
