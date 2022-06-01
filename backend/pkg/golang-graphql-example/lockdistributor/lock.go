package lockdistributor

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"cirello.io/pglock"
)

const transactionSerializeErrorMaxRetry = 1000

type lock struct {
	name string
	pl   *pglock.Lock
	s    *service
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

func (l *lock) AcquireWithContext(ctx context.Context) error {
	err := l.internalAcquireWithRetry(func() (*pglock.Lock, error) { return l.s.cl.AcquireContext(ctx, l.name) })
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (l *lock) Acquire() error {
	err := l.internalAcquireWithRetry(func() (*pglock.Lock, error) { return l.s.cl.Acquire(l.name) })
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (l *lock) internalAcquireWithRetry(acquire func() (*pglock.Lock, error)) error {
	var lastSerializeError error
	// Initialize counter
	counter := 0

	// Loop until the max retry is reached
	for counter < transactionSerializeErrorMaxRetry {
		// Acquire lock
		ll, err := acquire()
		// Check error
		if err != nil {
			// Check if it is a transaction serialize error
			if strings.Contains(err.Error(), "could not serialize access due to") {
				// Yes, so increment and retry
				counter++

				lastSerializeError = err

				continue
			}

			// No, abort here
			return errors.WithStack(err)
		}

		// Save lock
		l.pl = ll

		// Return lock
		return nil
	}

	// By default in this case, returning the transaction serialize error
	// with last error message
	return errors.Wrap(ErrAcquireTransactionSerialize, lastSerializeError.Error())
}

func (l *lock) IsReleased() (bool, error) {
	return l.pl.IsReleased(), nil
}

func (l *lock) Release() error {
	// Close
	err := l.pl.Close()
	// Check error
	if err != nil && !errors.Is(err, pglock.ErrLockAlreadyReleased) {
		return errors.WithStack(err)
	}
	// Default
	return nil
}
