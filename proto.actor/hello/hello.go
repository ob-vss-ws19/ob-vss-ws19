package main

import (
	"fmt"

	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
)

type hello struct{ who string }
type goodbye struct{ until string }
type pleaseReply struct{ from actor.PID }
type helloActor struct{}

func (*helloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *hello:
		fmt.Printf("Hello %+v\n", *context.Sender)
	case *goodbye:
		fmt.Printf("ok cu %v\n", msg.until)
	case int:
		fmt.Println("got int")
	}
}

func main() {
	context := actor.EmptyRootContext
	props := actor.PropsFromProducer(func() actor.Actor {
		return &helloActor{}
	})
	pid := context.Spawn(props)
	context.Send(pid, &hello{who: "Roger"})
	context.Send(pid, &goodbye{until: "Tomorrow"})
	context.Send(pid, 1)

	console.ReadLine() // nolint:errcheck
}
