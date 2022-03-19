package main

import (
	"app/pkg/keeper"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infof("Hello, world!")

	keeper.FeedAPet("Doggie")
}
