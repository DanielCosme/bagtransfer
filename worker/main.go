package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	app "fixity/bagtransfer"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("could not connect to Temporal server")
	}
	defer c.Close()

	w := worker.New(c, app.CheckFixityTaskQueue, worker.Options{})
	w.RegisterWorkflow(app.CheckFixityWorkflow)
	w.RegisterActivity(app.GreetActivity)
	w.RegisterActivity(app.GreetActivity2)
	// start listening on the Task Queue
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("unable to run worker", err)
	}
}
