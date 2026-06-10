package raft

import (
	"go-common/library/ecode/pb"
)

// StateType represents the role of a node in a cluster.
type StateType uint64

type stepFunc func(r *raft, m pb.Message) error


type raftLog struct {
	// storage contains all stable entries since the last snapshot.
	storage Storage

	// unstable contains all unstable entries and snapshot.
	// they will be saved into storage.
	unstable unstable
	// committed is the highest log position that is known to be in
	// stable storage on a quorum of nodes.
	committed uint64
	// applied is the highest log position that the application has
	// been instructed to apply to its state machine.
	// Invariant: applied <= committed
	applied uint64

	logger Logger
	maxNextEntsSize uint64
}
