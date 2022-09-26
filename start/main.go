package main

import (
	"context"
	app "fixity/bagtransfer"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("could not create Temporal client", err)
	}
	defer c.Close()
	opts := client.StartWorkflowOptions{
		ID:        "check_fixity_workflow",
		TaskQueue: app.CheckFixityTaskQueue,
	}

	input := app.FixityInput{
		Packages: []string{
			"fixity_pass",
			"fails_checksum",
			"missing_checksum",
		},
	}
	_, err = c.ExecuteWorkflow(
		context.Background(),
		opts, app.CheckFixityWorkflow,
		input,
	)
	if err != nil {
		log.Fatalln("error checking fixity", err)
	}
}
