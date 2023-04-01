package main

import (
	"log"
	"runtime"

	"github.com/nats-io/go-nats"
)

func main() {
	nc, err := nats.Connect("nats://192.168.1.86:4222",
		nats.Name("natsaudit"),
		nats.UserInfo("cg_nats_usr", "cg_nats_passwd"),
	)
	if err != nil {
		log.Fatalln(err)
	}

	nc.Subscribe(">", func(m *nats.Msg) {
		log.Printf("Subject: %s\n", string(m.Subject))
		log.Printf("Message: %s\n", string(m.Data))
		log.Printf("--------------------")
	})
	runtime.Goexit()
}
