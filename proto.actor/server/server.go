package main

import (
	"flag"
	"sync"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/log"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/ob-vss-ss19/ob-vss-ss19/proto.actor/echomessages"
)

type MyActor struct{}

func (state *MyActor) Receive(context actor.Context) {
	switch context.Message().(type) {
	case *echomessages.Echo:
		context.Respond(&echomessages.Response{
			SomeValue: "result",
		})
	default: // just for linter
	}
}

func NewMyActor() actor.Actor {
	log.Message("Hello-Actor is up and running")
	return &MyActor{}
}

// nolint:gochecknoglobals
var flagBind = flag.String("bind", "localhost:8091", "Bind to address")

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	defer wg.Wait()

	flag.Parse()
	remote.Start(*flagBind)

	remote.Register("hello", actor.PropsFromProducer(NewMyActor))
}
