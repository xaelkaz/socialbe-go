package main

import (
	"social-golang/src/config"
	"social-golang/src/route"
	"social-golang/src/storage"
)

func main() {
	config.Init()
	storage.Init()
	route.Init()
}
