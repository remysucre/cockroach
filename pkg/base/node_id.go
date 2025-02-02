// Copyright 2016 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package base

import (
	"context"
	"strconv"
	"sync/atomic"

	"github.com/cockroachdb/cockroach/pkg/roachpb"
	"github.com/cockroachdb/cockroach/pkg/util"
	"github.com/cockroachdb/cockroach/pkg/util/errorutil"
	"github.com/cockroachdb/cockroach/pkg/util/log"
	"github.com/cockroachdb/errors"
	"github.com/cockroachdb/redact"
)

// NodeIDContainer is used to share a single roachpb.NodeID instance between
// multiple layers. It allows setting and getting the value. Once a value is
// set, the value cannot change.
type NodeIDContainer struct {
	_ util.NoCopy

	// nodeID is accessed atomically.
	nodeID int32

	// If nodeID has been set, str represents nodeID converted to string. We
	// precompute this value to speed up String() and keep it from allocating
	// memory dynamically.
	str atomic.Value
}

// String returns the node ID, or "?" if it is unset.
func (n *NodeIDContainer) String() string {
	s := n.str.Load()
	if s == nil {
		return "?"
	}
	return s.(string)
}

var _ redact.SafeValue = &NodeIDContainer{}

// SafeValue implements the redact.SafeValue interface.
func (n *NodeIDContainer) SafeValue() {}

// Get returns the current node ID; 0 if it is unset.
func (n *NodeIDContainer) Get() roachpb.NodeID {
	return roachpb.NodeID(atomic.LoadInt32(&n.nodeID))
}

// Set sets the current node ID. If it is already set, the value must match.
func (n *NodeIDContainer) Set(ctx context.Context, val roachpb.NodeID) {
	if val <= 0 {
		log.Fatalf(ctx, "trying to set invalid NodeID: %d", val)
	}
	oldVal := atomic.SwapInt32(&n.nodeID, int32(val))
	if oldVal == 0 {
		if log.V(2) {
			log.Infof(ctx, "NodeID set to %d", val)
		}
	} else if oldVal != int32(val) {
		log.Fatalf(ctx, "different NodeIDs set: %d, then %d", oldVal, val)
	}
	n.str.Store(strconv.Itoa(int(val)))
}

// Reset changes the NodeID regardless of the old value.
//
// Should only be used in testing code.
func (n *NodeIDContainer) Reset(val roachpb.NodeID) {
	atomic.StoreInt32(&n.nodeID, int32(val))
	n.str.Store(strconv.Itoa(int(val)))
}

// StoreIDContainer is added as a logtag in the pebbleLogger's context.
// The storeID value is later set atomically. The storeID can only be
// set after engine creation because the storeID is determined only after the
// pebbleLogger's context is created.
type StoreIDContainer struct {
	_ util.NoCopy

	// storeID is accessed atomically.
	storeID int32

	// If storeID has been set, str represents storeID converted to string. We
	// precompute this value to speed up String() and keep it from allocating
	// memory dynamically.
	str atomic.Value
}

// TempStoreID is used as the store id for a temp pebble engine's log
const TempStoreID = -1

// String returns "temp" for temp stores, and the storeID for main
// stores if they haven't been initialized. If a main store hasn't
// been initialized, then "?" is returned.
func (s *StoreIDContainer) String() string {
	str := s.str.Load()
	if str == nil {
		return "?"
	}
	return str.(string)
}

var _ redact.SafeValue = &StoreIDContainer{}

// SafeValue implements the redact.SafeValue interface.
func (s *StoreIDContainer) SafeValue() {}

// Get returns the current storeID; 0 if it is unset.
func (s *StoreIDContainer) Get() int32 {
	return atomic.LoadInt32(&s.storeID)
}

// Set sets the current storeID. If it is already set, the value should match.
func (s *StoreIDContainer) Set(ctx context.Context, val int32) {
	if val != TempStoreID && val <= 0 {
		if log.V(2) {
			log.Infof(
				ctx, "trying to set invalid storeID for the store in the Pebble log: %d",
				val)
		}
		return
	}
	oldVal := atomic.SwapInt32(&s.storeID, val)
	if oldVal != 0 && oldVal != val {
		if log.V(2) {
			log.Infof(
				ctx, "different storeIDs set for the store in the Pebble log: %d, then %d",
				oldVal, val)
		}
	}
	if val == TempStoreID {
		s.str.Store("temp")
	} else {
		s.str.Store(strconv.Itoa(int(val)))
	}
}

// A SQLInstanceID is an ephemeral ID assigned to a running instance of the SQL
// server. This is distinct from a NodeID, which is a long-lived identifier
// assigned to a node in the KV layer which is unique across all KV nodes in the
// cluster and persists across restarts. Instead, a Instance is similar to a
// process ID from the unix world: an integer assigned to the SQL server
// on process start which is unique across all SQL server processes running
// on behalf of the tenant, while the SQL server is running.
//
// NB: until https://github.com/cockroachdb/cockroach/issues/47899 is addressed,
// the properties of the SQLInstanceID hold trivially due to the constraint that
// only one SQL server must be running on behalf of the tenant at any given
// time. After that, it's likely that we'll allocate these IDs off a counter,
// so they will be completely unique (per tenant).
type SQLInstanceID int32

func (s SQLInstanceID) String() string {
	return strconv.FormatInt(int64(s), 10)
}

// SQLIDContainer wraps a SQLInstanceID and optionally a NodeID.
type SQLIDContainer struct {
	w             errorutil.TenantSQLDeprecatedWrapper // NodeID
	sqlInstanceID SQLInstanceID
}

// NewSQLIDContainer sets up an SQLIDContainer. It is handed either a positive SQLInstanceID
// (on tenants) or a positive NodeID, but not both.
//
// A zero sqlInstanceID falls back to the NodeID in SQLInstanceID().
// This is used in single-tenant deployments.
//
// In a multi-tenant deployment, we could initialize the SQLIDContainer with
// a nil nodeIDContainer and 0 as the instance ID. This is to aid bootstrapping.
// In such a case, SetSQLInstanceID needs to be invoked prior to the SQLIDContainer
// being used.
func NewSQLIDContainer(sqlInstanceID SQLInstanceID, nodeID *NodeIDContainer) *SQLIDContainer {
	return &SQLIDContainer{
		w:             errorutil.MakeTenantSQLDeprecatedWrapper(nodeID, nodeID != nil),
		sqlInstanceID: sqlInstanceID,
	}
}

// SetSQLInstanceID sets the SQL instance ID. It returns an error if
// we attempt to set an instance ID when the nodeID has already been
// initialized.
func (c *SQLIDContainer) SetSQLInstanceID(sqlInstanceID SQLInstanceID) error {
	if _, ok := c.OptionalNodeID(); ok {
		return errors.New("attempting to initialize instance ID when node ID is set")
	}
	c.sqlInstanceID = sqlInstanceID
	return nil
}

// OptionalNodeID returns the NodeID and true, if the former is exposed.
// Otherwise, returns zero and false.
func (c *SQLIDContainer) OptionalNodeID() (roachpb.NodeID, bool) {
	v, ok := c.w.Optional()
	if !ok {
		return 0, false
	}
	return v.(*NodeIDContainer).Get(), true
}

// OptionalNodeIDErr is like OptionalNodeID, but returns an error (referring to
// the optionally supplied Github issues) if the ID is not present.
func (c *SQLIDContainer) OptionalNodeIDErr(issue int) (roachpb.NodeID, error) {
	v, err := c.w.OptionalErr(issue)
	if err != nil {
		return 0, err
	}
	return v.(*NodeIDContainer).Get(), nil
}

// SQLInstanceID returns the wrapped SQLInstanceID.
func (c *SQLIDContainer) SQLInstanceID() SQLInstanceID {
	if n, ok := c.OptionalNodeID(); ok {
		return SQLInstanceID(n)
	}
	return c.sqlInstanceID
}

// TestingIDContainer is an SQLIDContainer with hard-coded SQLInstanceID of 10 and
// NodeID of 1.
var TestingIDContainer = func() *SQLIDContainer {
	var c NodeIDContainer
	c.Set(context.Background(), 1)
	return NewSQLIDContainer(10, &c)
}()
