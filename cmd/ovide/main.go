package main

import (
	"log"
	"time"

	"github.com/iivvoo/novi/logger"
	"github.com/iivvoo/novi/ovide"
)

func main() {
	logger.OpenLog("ovide.log")
	log.Printf("Starting at %s\n", time.Now())
	defer logger.CloseLog()
	ovide.Run()
}
