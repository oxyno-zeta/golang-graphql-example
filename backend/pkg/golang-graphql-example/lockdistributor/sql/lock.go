package sqllockdistributor

import (
	"context"

	"cirello.io/pglock"
	"emperror.dev/errors"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
)

type lock struct {
	pl    *pglock.Lock
	s     *service
	trace tracing.Trace
	ctx   context.Context //nolint:containedctx // Keep the first context
	name  string
}

func (l *lock) IsAlreadyTaken() (bool, error) {
	// Get lock
	lo, err := l.s.cl.Get(l.name)
	// Check error
	if err != nil {
		// Check if error is a not found error
		if errors.Is(err, pglock.ErrLockNotFound) {
			return false, nil
		}

		return false, errors.WithStack(err)
	}

	// Check if lock exists or not
	return lo != nil, nil
}

func (l *lock) AcquireWithContext(ctx context.Context) (err error) {
	// Get trace
	trace := tracing.GetTraceFromContext(ctx)
	// Save it
	l.trace = trace
	l.ctx = ctx

	// Start trace
	ctx, ct := trace.GetChildTrace(ctx, "lockdistributor.Acquiring")
	// Add tags
	ct.SetTag("lock.name", l.name)
	ct.SetTag("lock.engine", "postgresql")
	// Defer end
	defer func() {
		// Check error
		if err != nil {
			ct.MarkAsError()
			ct.SetTag("lock.error", err.Error())
		}
		// End
		ct.Finish()
	}()

	// Create cancellable context
	cancelCtx, cancel := context.WithCancel(ctx) //nolint:govet
	// Create timeout
	timeoutCtx, cancelTimeout := context.WithTimeout(ctx, acquireTimeoutDuration)
	// Defer the cancel in case it is finishing earlier
	defer cancelTimeout()

	// Building error chan
	errChan := make(chan error)

	// Start acquire in routine to manage timeout
	go func() {
		// Acquire lock
		ll, err2 := l.s.cl.AcquireContext(cancelCtx, l.name)
		// Check error
		if err2 != nil {
			// Check if it is a not acquired error to wrap it
			if errors.Is(err2, pglock.ErrNotAcquired) {
				errChan <- ErrLockNotAcquired

				return
			}

			errChan <- errors.WithStack(err2)
		}
		// Save lock
		l.pl = ll
		// Inform
		errChan <- nil
	}()

	// Wait for timeout or result
	select {
	case <-timeoutCtx.Done():
		// Timeout raised => Need to cancel context
		cancel()

		return timeoutCtx.Err()
	case err = <-errChan:
		return err //nolint:govet
	}
}

func (l *lock) Acquire() error {
	return l.AcquireWithContext(context.TODO())
}

func (l *lock) IsReleased() (bool, error) {
	return l.pl.IsReleased(), nil
}

func (l *lock) Release() (err error) {
	// Get child trace
	_, ct := l.trace.GetChildTrace(l.ctx, "lockdistributor.Release")
	// Add tags
	ct.SetTag("lock.name", l.name)
	ct.SetTag("lock.engine", "postgresql")
	// Defer
	defer func() {
		// Check error
		if err != nil {
			ct.MarkAsError()
			ct.SetTag("lock.error", err.Error())
		}
		// End
		ct.Finish()
	}()

	// Close
	err = l.pl.Close()
	// Check error
	if err != nil && !errors.Is(err, pglock.ErrLockAlreadyReleased) {
		return errors.WithStack(err)
	}
	// Default
	return nil
}
