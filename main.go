package main

import (
	"fmt"
	"github.com/tirava/event_channel/pkg/event_channel"
	"log"
)

func main() {

	user1 := eventchannel.NewUser("jkulvich")
	user2 := eventchannel.NewUser("vasya")

	ch1 := eventchannel.NewChannel()
	ch1.Subscribe(user1)
	ch1.Subscribe(user2)

	ch2 := eventchannel.NewChannel()
	ch2.Subscribe(user2)

	pub := eventchannel.NewPublisher()
	pub.AddChannel("test", ch1)
	pub.AddChannel("test2", ch2)

	if err := pub.Send("HELLO!", "test"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.Send("HELLO FROM CH2", "test2"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	fmt.Println("Channels:", pub.ListChannels())
	pub.DeleteChannel("test")
	fmt.Println("Channels:", pub.ListChannels())

	// error here
	if err := pub.Send("HELLO!", "test"); err != nil {
		log.Println("can't send:", err)
	}
}
