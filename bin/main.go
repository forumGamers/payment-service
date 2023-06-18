package main

import (
	cfg "github.com/forumGamers/payment-service/config"
	r "github.com/forumGamers/payment-service/routes"
)

func main() {
	cfg.Connection()

	r.Routes()
}