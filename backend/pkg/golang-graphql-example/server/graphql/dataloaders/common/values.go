package common

import "time"

// Default wait for dataloaders.
const DefaultWait = 2 * time.Millisecond

// Default batch capacity for dataloaders.
const DefaultBatchCapacity = 100
