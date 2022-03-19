package main

import (
	"github.com/expiron/gopet/pkg/keeper"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infof("Hello, world!")

	keeper.FeedAPet("Doggie")
}
