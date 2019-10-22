package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/ob-vss-ss19/ob-vss-ss19/proto.actor/echomessages"
)

type MyActor struct {
	count int
	wg    *sync.WaitGroup
}

func (state *MyActor) Receive(context actor.Context) {
	switch context.Message().(type) {
	case *echomessages.Response:
		state.count++
		fmt.Println(state.count)
	case *actor.Stopped:
		state.wg.Done()
	}
}

var (
	// nolint:gochecknoglobals
	flagBind = flag.String("bind", "localhost:8090", "Bind to address")
	// nolint:gochecknoglobals
	flagRemote = flag.String("remote", "localhost:8091", "remote host:port")
)

func main() {
	flag.Parse()

	remote.Start(*flagBind)

	var wg sync.WaitGroup

	props := actor.PropsFromProducer(func() actor.Actor {
		wg.Add(1)
		return &MyActor{0, &wg}
	})
	rootContext := actor.EmptyRootContext
	pid := rootContext.Spawn(props)
	message := &echomessages.Echo{Message: "hej"}

	fmt.Println("Sleeping 5 seconds...")
	time.Sleep(5 * time.Second)
	fmt.Println("Awake...")

	//this is the remote actor we want to communicate with
	fmt.Printf("Trying to connect to %s\n", *flagRemote)

	pidResp, err := remote.SpawnNamed(*flagRemote, "remote", "hello", 5*time.Second)
	if err != nil {
		panic(err)
	}

	remotePid := pidResp.Pid

	for i := 0; i < 10; i++ {
		rootContext.RequestWithCustomSender(remotePid, message, pid)
	}

	wg.Wait()
}
