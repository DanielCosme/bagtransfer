package app

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func CheckFixityWorkflow(ctx workflow.Context, data Data) (*Data, error) {
	// Activity Options
	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         defaultRetryPolicy(),
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// Package data
	if err := workflow.ExecuteActivity(ctx, BagItActivity, data).Get(ctx, nil); err != nil {
		return nil, err
	}
	return &data, nil
}

func defaultRetryPolicy() *temporal.RetryPolicy {
	return &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0, // exponential backoff
		MaximumInterval:    time.Minute,
		MaximumAttempts:    10,
	}
}
