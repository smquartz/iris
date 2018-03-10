package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/smquartz/iris/handler"
	"github.com/smquartz/iris/subscriber"

	example "github.com/smquartz/iris/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("video.quartz.famethyst.redeye.srv.iris"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("topic.video.quartz.famethyst.redeye.srv.iris", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("topic.video.quartz.famethyst.redeye.srv.iris", service.Server(), subscriber.Handler)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
