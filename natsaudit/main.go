package main

import (
	"log"
	"runtime"

	"github.com/nats-io/go-nats"
)

func main() {
	nc, err := nats.Connect("nats://192.168.1.70:4222",
		nats.Name("natsaudit"),
		nats.UserInfo("foo", "secret"),
	)
	if err != nil {
		log.Fatalln(err)
	}

	nc.Subscribe(">", func(m *nats.Msg) {
		log.Printf("%#v", *m)
	})
	runtime.Goexit()
}
