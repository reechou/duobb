package main

import (
	"github.com/reechou/duobb/config"
	"github.com/reechou/duobb/controller"
)

func main() {
	controller.NewLogic(config.NewConfig()).Run()
}
