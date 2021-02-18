package txpool

import (
	. "yu/common"
	. "yu/txn"
)

type ItxPool interface {
	// insert into txCache for pending
	Pend(IsignedTxn) error
	// insert into txPool for tripods
	Insert(BlockNum, IsignedTxn) error
	// package some txns to send to tripods
	Package(numLimit uint64) ([]IsignedTxn, error)
	// get txn content of txn-hash from p2p network
	SyncTxns([]Hash) error
	// broadcast txns to p2p network
	BroadcastTxns() error
	// pop pending txns
	Pop() (IsignedTxn, error)
	// remove txns after execute all tripods
	Remove() error
}

type ItxCache interface {
	Push(IsignedTxn) error
	Pop() (IsignedTxn, error)
}