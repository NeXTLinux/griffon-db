package pkg

import (
	"github.com/wagoodman/go-partybus"

	"github.com/anchore/go-logger"
	"github.com/nextlinux/griffon-db/internal/bus"
	"github.com/nextlinux/griffon-db/internal/log"
)

func SetLogger(l logger.Logger) {
	log.Set(l)
}

func SetBus(b *partybus.Bus) {
	bus.SetPublisher(b)
}
