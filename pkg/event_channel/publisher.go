package eventchannel

import (
	"fmt"
)

type Channels map[string]*Channel

type Publisher struct {
	channels Channels
}

func NewPublisher() *Publisher {
	return &Publisher{
		channels: Channels{},
	}
}

func (p *Publisher) AddChannel(name string, channel *Channel) {
	p.channels[name] = channel
}

func (p *Publisher) DeleteChannel(name string) {
	delete(p.channels, name)
}

func (p *Publisher) ListChannels() []string {
	list := make([]string, 0, len(p.channels))
	for name := range p.channels {
		list = append(list, name)
	}

	return list
}

func (p *Publisher) Send(msg string, channels ...string) error {
	for _, ch := range channels {
		channel, ok := p.channels[ch]
		if !ok {
			return fmt.Errorf("channel %s can't be found", ch)
		}
		channel.Send(msg)
	}

	return nil
}
