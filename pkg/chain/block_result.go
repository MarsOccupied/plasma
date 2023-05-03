package chain

import (
	"github.com/MarsOccupied/plasma-core/plasma/pkg/chain"
	"github.com/MarsOccupied/plasma-core/plasma/util"
	"github.com/MarsOccupied/plasma-core/plasma/pkg/rpc/pb"
	"github.com/pkg/errors"
	"github.com/ethereum/go-ethereum/rlp"
)

import (
	"math/big"
)

type BlockResult struct {
	MerkleRoot         util.Hash
	NumberTransactions uint32
	BlockFees          *big.Int
	BlockNumber        *big.Int
}
