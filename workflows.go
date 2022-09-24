package app

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func CheckFixityWorkflow(ctx workflow.Context, data Data) error {
	// Activity Options
	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         defaultRetryPolicy(),
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// Workflow steps??
	// For now just say hello world.
	if err := workflow.ExecuteActivity(ctx, GreetActivity, data).Get(ctx, nil); err != nil {
		return err
	}

	if err := workflow.ExecuteActivity(ctx, GreetActivity2, data).Get(ctx, nil); err != nil {
		return err
	}

	return nil
}

func defaultRetryPolicy() *temporal.RetryPolicy {
	return &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0, // exponential backoff
		MaximumInterval:    time.Minute,
		MaximumAttempts:    10,
	}
}
