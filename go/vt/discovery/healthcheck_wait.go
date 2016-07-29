package discovery

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"golang.org/x/net/context"

	log "github.com/golang/glog"

	"github.com/youtube/vitess/go/vt/concurrency"
	"github.com/youtube/vitess/go/vt/topo"

	topodatapb "github.com/youtube/vitess/go/vt/proto/topodata"
)

var (
	// ErrWaitForTabletsTimeout is returned if we cannot get the tablets in time
	ErrWaitForTabletsTimeout = errors.New("timeout waiting for tablets")

	// how much to sleep between each check
	waitAvailableTabletInterval = 100 * time.Millisecond
)

// keyspaceShard is a helper structure used internally
type keyspaceShard struct {
	keyspace string
	shard    string
}

// WaitForTablets waits for at least one tablet in the given cell /
// keyspace / shard before returning.
func WaitForTablets(ctx context.Context, hc HealthCheck, cell, keyspace, shard string, types []topodatapb.TabletType) error {
	keyspaceShards := map[keyspaceShard]bool{
		keyspaceShard{
			keyspace: keyspace,
			shard:    shard,
		}: true,
	}
	return waitForTablets(ctx, hc, keyspaceShards, types, false)
}

// WaitForAllServingTablets waits for at least one serving tablet in the given cell
// for all keyspaces / shards before returning.
func WaitForAllServingTablets(ctx context.Context, hc HealthCheck, ts topo.SrvTopoServer, cell string, types []topodatapb.TabletType) error {
	keyspaceShards, err := findAllKeyspaceShards(ctx, ts, cell)
	if err != nil {
		return err
	}

	return waitForTablets(ctx, hc, keyspaceShards, types, true)
}

// findAllKeyspaceShards goes through all serving shards in the topology
func findAllKeyspaceShards(ctx context.Context, ts topo.SrvTopoServer, cell string) (map[keyspaceShard]bool, error) {
	ksNames, err := ts.GetSrvKeyspaceNames(ctx, cell)
	if err != nil {
		return nil, err
	}

	keyspaceShards := make(map[keyspaceShard]bool)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errRecorder concurrency.AllErrorRecorder
	for _, ksName := range ksNames {
		wg.Add(1)
		go func(keyspace string) {
			defer wg.Done()

			// get SrvKeyspace for cell/keyspace
			ks, err := ts.GetSrvKeyspace(ctx, cell, keyspace)
			if err != nil {
				errRecorder.RecordError(err)
				return
			}

			// get all shard names that are used for serving
			mu.Lock()
			for _, ksPartition := range ks.Partitions {
				for _, shard := range ksPartition.ShardReferences {
					keyspaceShards[keyspaceShard{
						keyspace: keyspace,
						shard:    shard.Name,
					}] = true
				}
			}
			mu.Unlock()
		}(ksName)
	}
	wg.Wait()
	if errRecorder.HasErrors() {
		return nil, errRecorder.Error()
	}

	return keyspaceShards, nil
}

// waitForTablets is the internal method that polls for tablets
func waitForTablets(ctx context.Context, hc HealthCheck, keyspaceShards map[keyspaceShard]bool, types []topodatapb.TabletType, requireServing bool) error {
RetryLoop:
	for {
		select {
		case <-ctx.Done():
			break RetryLoop
		default:
			// Context is still valid. Move on.
		}

		for ks := range keyspaceShards {
			allPresent := true
			for _, tt := range types {
				tl := hc.GetTabletStatsFromTarget(ks.keyspace, ks.shard, tt)
				if requireServing {
					hasServingEP := false
					for _, t := range tl {
						if t.LastError == nil && t.Serving {
							hasServingEP = true
							break
						}
					}
					if !hasServingEP {
						allPresent = false
						break
					}
				} else {
					if len(tl) == 0 {
						allPresent = false
						break
					}
				}
			}

			if allPresent {
				delete(keyspaceShards, ks)
			}
		}

		if len(keyspaceShards) == 0 {
			// we found everything we needed
			return nil
		}

		// Unblock after the sleep or when the context has expired.
		timer := time.NewTimer(waitAvailableTabletInterval)
		select {
		case <-ctx.Done():
			timer.Stop()
		case <-timer.C:
		}
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Warningf("waitForTablets timeout for %v (context error: %v)", keyspaceShards, ctx.Err())
		return ErrWaitForTabletsTimeout
	}
	err := fmt.Errorf("waitForTablets failed for %v (context error: %v)", keyspaceShards, ctx.Err())
	log.Error(err)
	return err
}
