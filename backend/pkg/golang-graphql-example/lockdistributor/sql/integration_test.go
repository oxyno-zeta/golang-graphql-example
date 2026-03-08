//go:build integration

package sqllockdistributor

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	cmocks "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config/mocks"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/database"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing"
)

type LockDistributorTestSuite struct {
	suite.Suite

	cfg           *config.Config
	cfgManager    config.Manager
	logger        log.Logger
	tracingSvc    tracing.Service
	metricsSvc    metrics.Service
	db            database.DB
	ld            Service
	lockTableName string
}

func (suite *LockDistributorTestSuite) SetupSuite() {
	fmt.Println("SetupSuite phase")
	lockTableName := config.DefaultLockDistributorTableName

	cfg := &config.Config{
		Log:     &config.LogConfig{Level: "debug", Format: "human"},
		Tracing: &config.TracingConfig{Enabled: false},
		LockDistributor: &config.LockDistributorConfig{
			TableName:          lockTableName,
			LeaseDuration:      config.DefaultLockDistributorLeaseDuration,
			HeartbeatFrequency: config.DefaultLockDistributionHeartbeatFrequency,
		},
		Database: &config.DatabaseConfig{
			Driver: config.DefaultDatabaseDriver,
			ConnectionURL: &config.CredentialConfig{
				Value: "host=localhost port=5432 user=postgres dbname=postgres-integration password=postgres sslmode=disable",
			},
		},
	}

	ctrl := gomock.NewController(suite.T())
	cfgManagerMock := cmocks.NewMockManager(ctrl)
	cfgManagerMock.EXPECT().GetConfig().AnyTimes().Return(cfg)

	logger := log.NewLogger()
	err := logger.Configure("debug", "human", "")
	suite.NoError(err)

	metricsSvc := metrics.NewService()
	suite.metricsSvc = metricsSvc
	tracingSvc := tracing.New(cfgManagerMock, logger)
	err = tracingSvc.InitializeAndReload()
	suite.NoError(err)

	db := database.NewDatabase("main", cfgManagerMock, logger, metricsSvc, tracingSvc)
	err = db.Connect()
	suite.NoError(err)

	ld := NewService(cfgManagerMock, db)
	err = ld.InitializeAndReload(logger)
	suite.NoError(err)

	suite.db = db
	suite.ld = ld
	suite.lockTableName = lockTableName
	suite.cfg = cfg
	suite.cfgManager = cfgManagerMock
	suite.logger = logger
	suite.tracingSvc = tracingSvc
}

func (suite *LockDistributorTestSuite) TearDownSuite() {
	fmt.Println("TearDownSuite phase")
	if suite.db != nil {
		suite.NoError(suite.db.Close())
	}
}

func (suite *LockDistributorTestSuite) BeforeTest(_, _ string) {
	fmt.Println("BeforeTest phase")
	suite.cleanLocks()
}

func (suite *LockDistributorTestSuite) AfterTest(_, _ string) {
	fmt.Println("AfterTest phase")
	suite.cleanLocks()
}

func (suite *LockDistributorTestSuite) cleanLocks() {
	if suite.db == nil {
		return
	}

	gdb := suite.db.GetGormDB().Exec(fmt.Sprintf("TRUNCATE TABLE %s;", suite.lockTableName))
	suite.NoError(gdb.Error)
}

// What is tested?
// Concurrent acquire attempts on the same lock name.
// Expected results:
// First acquire succeeds, second concurrent acquire on same lock fails while the first lock is held.
func (suite *LockDistributorTestSuite) TestConcurrentAcquireWithContext_SameLockCannotBeAcquiredAtSameTime() {
	first := suite.ld.GetLock("same-lock")
	suite.NoError(first.AcquireWithContext(context.Background()))
	defer func() { suite.NoError(first.Release()) }()

	second := suite.ld.GetLock("same-lock")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := second.AcquireWithContext(ctx)
	suite.Error(err)
	suite.True(errors.Is(err, ErrLockNotAcquired) || errors.Is(err, context.DeadlineExceeded))
}

// What is tested?
// Same-lock contention in the window after heartbeat tick and before lease expiration.
// Expected results:
// Second acquire still fails after a delay just above heartbeat and below lease duration.
func (suite *LockDistributorTestSuite) TestConcurrentAcquireWithContext_SameLockCannotBeAcquired_BetweenHeartbeatAndLease() {
	heartbeat, err := time.ParseDuration(suite.cfg.LockDistributor.HeartbeatFrequency)
	suite.NoError(err)
	lease, err := time.ParseDuration(suite.cfg.LockDistributor.LeaseDuration)
	suite.NoError(err)

	delay := heartbeat + 300*time.Millisecond
	suite.Less(delay, lease)

	first := suite.ld.GetLock("same-lock-between-heartbeat-lease")
	suite.NoError(first.AcquireWithContext(context.Background()))
	defer func() { suite.NoError(first.Release()) }()

	time.Sleep(delay)

	second := suite.ld.GetLock("same-lock-between-heartbeat-lease")
	ctx, cancel := context.WithTimeout(context.Background(), 700*time.Millisecond)
	defer cancel()

	err = second.AcquireWithContext(ctx)
	suite.Error(err)
	suite.True(errors.Is(err, ErrLockNotAcquired) || errors.Is(err, context.DeadlineExceeded))
}

// What is tested?
// Same-lock contention very close to lease expiry but still before expiry.
// Expected results:
// Second acquire still fails when attempted shortly before lease duration is reached.
func (suite *LockDistributorTestSuite) TestConcurrentAcquireWithContext_SameLockCannotBeAcquired_JustBeforeLease() {
	lease, err := time.ParseDuration(suite.cfg.LockDistributor.LeaseDuration)
	suite.NoError(err)

	delay := lease - 300*time.Millisecond
	suite.Greater(delay, time.Duration(0))

	first := suite.ld.GetLock("same-lock-before-lease")
	suite.NoError(first.AcquireWithContext(context.Background()))
	defer func() { suite.NoError(first.Release()) }()

	time.Sleep(delay)

	second := suite.ld.GetLock("same-lock-before-lease")
	ctx, cancel := context.WithTimeout(context.Background(), 700*time.Millisecond)
	defer cancel()

	err = second.AcquireWithContext(ctx)
	suite.Error(err)
	suite.True(errors.Is(err, ErrLockNotAcquired) || errors.Is(err, context.DeadlineExceeded))
}

// What is tested?
// Regression test for commit 7b754565ee8995a12ef2e9df3dd01f4b57c8b586:
// successful acquire must keep lock alive even after waiting past lease duration.
// Expected results:
// A second acquire on the same lock remains blocked after waiting >3s (default lease duration).
// Before the fix, heartbeat could stop after acquire, lease would expire, and second acquire could succeed.
func (suite *LockDistributorTestSuite) TestAcquireWithContext_Regression_HeartbeatNotReleasedAfterAcquire() {
	secondDB := database.NewDatabase("secondary", suite.cfgManager, suite.logger, suite.metricsSvc, suite.tracingSvc)
	suite.NoError(secondDB.Connect())
	defer func() { suite.NoError(secondDB.Close()) }()

	secondService := NewService(suite.cfgManager, secondDB)
	suite.NoError(secondService.InitializeAndReload(suite.logger))

	first := suite.ld.GetLock("regression-heartbeat-lock")
	suite.NoError(first.AcquireWithContext(context.Background()))
	defer func() { suite.NoError(first.Release()) }()

	// Integration config lease is 3s and heartbeat is 1s.
	// Sleep past lease duration to ensure we detect missing heartbeat renewal.
	time.Sleep(4500 * time.Millisecond)

	released, err := first.IsReleased()
	suite.NoError(err)
	suite.False(released)

	second := secondService.GetLock("regression-heartbeat-lock")
	ctx, cancel := context.WithTimeout(context.Background(), 1200*time.Millisecond)
	defer cancel()

	err = second.AcquireWithContext(ctx)
	suite.Error(err)
	suite.True(errors.Is(err, ErrLockNotAcquired) || errors.Is(err, context.DeadlineExceeded))
}

// What is tested?
// Concurrent acquire attempts on different lock names.
// Expected results:
// Both acquires succeed concurrently because lock names are independent.
func (suite *LockDistributorTestSuite) TestConcurrentAcquireWithContext_DifferentLocksCanBeAcquiredAtSameTime() {
	lockA := suite.ld.GetLock("lock-a")
	lockB := suite.ld.GetLock("lock-b")

	errCh := make(chan error, 2)

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		errCh <- lockA.AcquireWithContext(ctx)
	}()

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		errCh <- lockB.AcquireWithContext(ctx)
	}()

	suite.NoError(<-errCh)
	suite.NoError(<-errCh)

	suite.NoError(lockA.Release())
	suite.NoError(lockB.Release())
}

// What is tested?
// Re-acquiring a lock after it has been released.
// Expected results:
// Acquire succeeds for first holder, release succeeds, then a second acquire on same lock succeeds.
func (suite *LockDistributorTestSuite) TestAcquireWithContext_SameLockCanBeAcquiredAfterRelease() {
	first := suite.ld.GetLock("reacquire-same-lock")
	suite.NoError(first.AcquireWithContext(context.Background()))
	suite.NoError(first.Release())

	second := suite.ld.GetLock("reacquire-same-lock")
	suite.NoError(second.AcquireWithContext(context.Background()))
	suite.NoError(second.Release())
}

// What is tested?
// Acquire path without explicit context.
// Expected results:
// Acquire() succeeds and lock can be released successfully.
func (suite *LockDistributorTestSuite) TestAcquire_WorksWithoutContext() {
	l := suite.ld.GetLock("acquire-no-context")
	suite.NoError(l.Acquire())
	suite.NoError(l.Release())
}

// What is tested?
// Releasing the same lock multiple times.
// Expected results:
// First release succeeds and second release is tolerated (idempotent behavior in wrapper).
func (suite *LockDistributorTestSuite) TestRelease_IsIdempotent() {
	l := suite.ld.GetLock("release-idempotent")
	suite.NoError(l.AcquireWithContext(context.Background()))
	suite.NoError(l.Release())
	suite.NoError(l.Release())
}

// What is tested?
// IsAlreadyTaken state before acquire and while lock is held.
// Expected results:
// false before acquire, true while held
func (suite *LockDistributorTestSuite) TestIsAlreadyTaken_Lifecycle() {
	l := suite.ld.GetLock("already-taken-lifecycle")

	taken, err := l.IsAlreadyTaken()
	suite.NoError(err)
	suite.False(taken)

	suite.NoError(l.AcquireWithContext(context.Background()))
	taken, err = l.IsAlreadyTaken()
	suite.NoError(err)
	suite.True(taken)

	suite.NoError(l.Release())
}

// What is tested?
// IsReleased state transition for a held lock.
// Expected results:
// false after acquire and true after release.
func (suite *LockDistributorTestSuite) TestIsReleased_Lifecycle() {
	l := suite.ld.GetLock("is-released-lifecycle")
	suite.NoError(l.AcquireWithContext(context.Background()))

	released, err := l.IsReleased()
	suite.NoError(err)
	suite.False(released)

	suite.NoError(l.Release())
	released, err = l.IsReleased()
	suite.NoError(err)
	suite.True(released)
}

// What is tested?
// High-concurrency acquire on many distinct lock names.
// Expected results:
// All goroutines acquire and release successfully with no cross-lock interference.
func (suite *LockDistributorTestSuite) TestConcurrentAcquireWithContext_ManyDifferentLocksSucceed() {
	const workers = 10
	errCh := make(chan error, workers)

	for i := range workers {
		go func(idx int) {
			lockName := fmt.Sprintf("many-locks-%d", idx)
			l := suite.ld.GetLock(lockName)

			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			if err := l.AcquireWithContext(ctx); err != nil {
				errCh <- err
				return
			}
			errCh <- l.Release()
		}(i)
	}

	for range workers {
		suite.NoError(<-errCh)
	}
}

// What is tested?
// Progress under contention where many goroutines target the same lock.
// Expected results:
// Each worker eventually acquires/releases the lock within timeout; total successful acquires equals worker count.
func (suite *LockDistributorTestSuite) TestAcquireWithContext_SameLockContentionMakesProgress() {
	const workers = 4

	var acquiredCount int32
	var wg sync.WaitGroup
	wg.Add(workers)

	for range workers {
		go func() {
			defer wg.Done()
			l := suite.ld.GetLock("single-contention-lock")
			err := suite.acquireEventually(l, 2*time.Second)
			if err != nil {
				return
			}
			atomic.AddInt32(&acquiredCount, 1)
			time.Sleep(20 * time.Millisecond)
			_ = l.Release()
		}()
	}

	wg.Wait()
	suite.Equal(int32(workers), acquiredCount)
}

// What is tested?
// Deadline behavior when lock is already held by another owner.
// Expected results:
// Acquire attempt returns an error quickly and respects the short context deadline bound.
func (suite *LockDistributorTestSuite) TestAcquireWithContext_RespectsShortDeadline() {
	holder := suite.ld.GetLock("short-deadline-lock")
	suite.NoError(holder.AcquireWithContext(context.Background()))
	defer func() { suite.NoError(holder.Release()) }()

	waiter := suite.ld.GetLock("short-deadline-lock")
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	start := time.Now()
	err := waiter.AcquireWithContext(ctx)
	elapsed := time.Since(start)

	suite.Error(err)
	suite.LessOrEqual(elapsed, 500*time.Millisecond)
}

// What is tested?
// Immediate cancellation behavior for acquire with a canceled context.
// Expected results:
// Acquire fails promptly with a cancellation-related error.
func (suite *LockDistributorTestSuite) TestAcquireWithContext_CanceledContextReturnsPromptly() {
	holder := suite.ld.GetLock("cancelled-context-lock")
	suite.NoError(holder.AcquireWithContext(context.Background()))
	defer func() { suite.NoError(holder.Release()) }()

	waiter := suite.ld.GetLock("cancelled-context-lock")
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	start := time.Now()
	err := waiter.AcquireWithContext(ctx)
	elapsed := time.Since(start)

	suite.Error(err)
	suite.True(errors.Is(err, context.Canceled) || errors.Is(err, ErrLockNotAcquired))
	suite.LessOrEqual(elapsed, 200*time.Millisecond)
}

// acquireEventually keeps trying to acquire a lock until success or total timeout.
// It uses short per-attempt context deadlines so each try returns quickly.
// Retryable errors are:
// - context.DeadlineExceeded: this specific attempt timed out
// - ErrLockNotAcquired: lock is currently held by someone else
// For retryable errors, it waits a short backoff and tries again.
// Any other error is considered unexpected and returned immediately.
// If the global deadline is reached first, it returns context.DeadlineExceeded.
func (suite *LockDistributorTestSuite) acquireEventually(l Lock, totalTimeout time.Duration) error {
	deadline := time.Now().Add(totalTimeout)
	for time.Now().Before(deadline) {
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		err := l.AcquireWithContext(ctx)
		cancel()
		if err == nil {
			return nil
		}
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, ErrLockNotAcquired) {
			time.Sleep(20 * time.Millisecond)
			continue
		}
		return err
	}
	return context.DeadlineExceeded
}

func TestLockDistributorTestSuite(t *testing.T) {
	suite.Run(t, new(LockDistributorTestSuite))
}
