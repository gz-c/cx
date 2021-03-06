package webrpc

import (
	"github.com/amherag/skycoin/src/cipher"
	"github.com/amherag/skycoin/src/coin"
	"github.com/amherag/skycoin/src/visor"
	"github.com/amherag/skycoin/src/visor/historydb"
)

//go:generate mockery -name Gatewayer -case underscore -inpkg -testonly

// Gatewayer provides interfaces for getting skycoin related info.
type Gatewayer interface {
	GetLastBlocks(num uint64) ([]coin.SignedBlock, error)
	GetBlocksInRange(start, end uint64) ([]coin.SignedBlock, error)
	GetBlocks(vs []uint64) ([]coin.SignedBlock, error)
	GetUnspentOutputsSummary(filters []visor.OutputsFilter) (*visor.UnspentOutputsSummary, error)
	GetTransaction(txid cipher.SHA256) (*visor.Transaction, error)
	InjectBroadcastTransaction(tx coin.Transaction) error
	GetSpentOutputsForAddresses(addr []cipher.Address) ([][]historydb.UxOut, error)
}
