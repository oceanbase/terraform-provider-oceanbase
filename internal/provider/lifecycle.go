package provider

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

// Only an authoritative lookup of the managed identity may return this error.
var errResourceNotFound = errors.New("resource not found")

func operationTimeout(env string, fallback time.Duration) time.Duration {
	seconds, err := strconv.ParseInt(os.Getenv(env), 10, 32)
	if err != nil || seconds <= 0 {
		return fallback
	}
	return time.Duration(seconds) * time.Second
}

func waitForRemoval(ctx context.Context, timeout time.Duration, read func(context.Context) error) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return retry.RetryContext(ctx, timeout, func() *retry.RetryError {
		err := read(ctx)
		switch {
		case errors.Is(err, errResourceNotFound):
			return nil
		case err != nil:
			return retry.NonRetryableError(err)
		default:
			return retry.RetryableError(errors.New("resource is still present"))
		}
	})
}
