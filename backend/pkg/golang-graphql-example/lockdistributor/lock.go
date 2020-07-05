package lockdistributor

import "cirello.io/pglock"

type lock struct {
	name string
	pl   *pglock.Lock
	s    *service
}

func (l *lock) Acquire() error {
	ll, err := l.s.cl.Acquire(l.name)
	// Check error
	if err != nil {
		return err
	}
	// Save lock
	l.pl = ll

	return nil
}

func (l *lock) Release() error {
	return l.pl.Close()
}
