package chain

import (
	"github.com/MarsOccupied/plasma-core/plasma/pkg/chain"
	"github.com/MarsOccupied/plasma-core/plasma/util"
	"github.com/MarsOccupied/plasma-core/plasma/pkg/rpc/pb"
	"github.com/pkg/errors"
	"github.com/ethereum/go-ethereum/rlp"
)

type BlockWithMeta struct {
	Block                 *Block                 `json:"block"`
	Metadata              *BlockMetadata         `json:"metadata"`
	ConfirmedTransactions []ConfirmedTransaction `json:"confirmedTransactions"`
}
