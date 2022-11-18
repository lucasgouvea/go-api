package main

import (
	Redis "go-api/internal/redis"
)

func Setup() {

	Redis.NewPool()
	connection := Redis.GetConnection()
	connectionError := connection.Err()
	if connectionError != nil {
		panic(connectionError)
	}

}
