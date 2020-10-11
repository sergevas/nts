package main

import (
	"log"
	"time"

	"github.com/nats-io/go-nats"
)

type payload struct {
	RequestID string
	Data      string
}

func (p *payload) create(requestID, data string) payload {
	return payload{requestID, data}
}

func main() {
	requestSub := "cg/d8f15b10c89c"
	nc, err := nats.Connect("nats://192.168.1.70:4222",
		nats.Name("temp-connected-garden"),
		nats.UserInfo("foo", "secret"),
	)
	if err != nil {
		log.Fatalln(err)
	}

	nc.Subscribe("cg/status", func(m *nats.Msg) {
		log.Printf("[Received] %s", string(m.Data))
		log.Printf("[Received msg] %+v", *m)
		nc.Publish(m.Reply, []byte("{\"id\": \"5ccf7f2f1d04\", \"t\": 5.67}"))
	})

	for {
		resp, err := nc.Request(requestSub, []byte("req 5ccf7f2f1d04"), 5*time.Second)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("[Response]: %s", string(resp.Data))
	}
}
