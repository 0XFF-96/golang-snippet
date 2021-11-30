package raft

import (
	"go-common/library/ecode/pb"
)

type raft struct {
	id uint64
	Term uint64
	Vote uint64

	readStates []ReadState
	// the log
	raftLog *raftLog

	maxMsgSize         uint64
	maxUncommittedSize uint64
	// TODO(tbg): rename to trk.
	prs tracker.ProgressTracker
	state StateType

	// isLearner is true if the local raft node is a learner.
	isLearner bool
	msgs []pb.Message
	lead uint64
	leadTransferee uint64
	pendingConfIndex uint64
	uncommittedSize uint64
	electionElapsed int
	heartbeatElapsed int

	checkQuorum bool // 这个是什么意思？
	preVote     bool

	heartbeatTimeout int
	electionTimeout  int
	// randomizedElectionTimeout is a random number between
	// [electiontimeout, 2 * electiontimeout - 1]. It gets reset
	// when raft changes its state to follower or candidate.
	randomizedElectionTimeout int
	disableProposalForwarding bool

	tick func()
	step stepFunc
	// logger Logger
}

