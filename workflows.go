package app

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func CheckFixityWorkflow(ctx workflow.Context, data FixityInput) (*Result, error) {
	// Activity Options
	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         defaultRetryPolicy(),
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	results := &Result{}
	for _, p := range data.Packages {
		res := &FixityResult{}
		if err := workflow.ExecuteActivity(ctx, Fixitycheck, p).Get(ctx, res); err != nil {
			fmt.Println(err.Error())
			// return err
		}
		results.Data = append(results.Data, res)
	}
	results.TranfersChecked = len(results.Data)
	// return results, errors.New("Handmade error")
	return results, nil
}

func defaultRetryPolicy() *temporal.RetryPolicy {
	return &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0, // exponential backoff
		MaximumInterval:    time.Minute,
		MaximumAttempts:    10,
	}
}
