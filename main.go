package main

import (
	"chat/internal/bootstrap"
	_ "embed"
)

var (
	service = "chat"
	version = "latest"
)

//go:embed banner.txt
var banner string

func init() {
	bootstrap.Init(service, version, banner)
}

func main() {
	bootstrap.Run()
}
