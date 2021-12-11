package node

import (
	"math/big"

	"github.com/PatriceVignola/rocketpool-go/node"
	"github.com/PatriceVignola/rocketpool-go/rocketpool"
	"github.com/PatriceVignola/rocketpool-go/tokens"

	"github.com/PatriceVignola/rocketpool-go/tests/testutils/accounts"
	rplutils "github.com/PatriceVignola/rocketpool-go/tests/testutils/tokens/rpl"
)

// Mint & stake an amount of RPL against a node
func StakeRPL(rp *rocketpool.RocketPool, ownerAccount, nodeAccount *accounts.Account, amount *big.Int) error {

	// Get RocketNodeStaking contract address
	rocketNodeStakingAddress, err := rp.GetAddress("rocketNodeStaking")
	if err != nil {
		return err
	}

	// Mint, approve & stake RPL
	if err := rplutils.MintRPL(rp, ownerAccount, nodeAccount, amount); err != nil {
		return err
	}
	if _, err := tokens.ApproveRPL(rp, *rocketNodeStakingAddress, amount, nodeAccount.GetTransactor()); err != nil {
		return err
	}
	if _, err := node.StakeRPL(rp, amount, nodeAccount.GetTransactor()); err != nil {
		return err
	}

	// Return
	return nil

}
