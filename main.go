package main

import (
	"gallery/comfy"
	"gallery/core"
	"gallery/db"
	"gallery/db/redis"
	"gallery/fiber"
	"gallery/nats"
)

var (
	log = core.GetLogger()
)

func main() {
	log.Info("getting comfy config..")
	comfyErr := comfy.Startup()
	if comfyErr != nil {
		// log.Error(comfyErr)
		log.Fatal(comfyErr)
	}
	defer core.HandlePanic()
	log.Info("Connecting to Mongo")
	db.Connect()
	log.Info("Connecting to Redis")
	redis.Start()
	redis.WaitActive()
	log.Info("starting nats server..")
	nats.Start()
	log.Info("nats setup done")
	log.Info("starting server")
	fiber.Start()
}
