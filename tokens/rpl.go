package tokens

import (
    "math/big"
    "sync"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"

    "github.com/rocket-pool/rocketpool-go/rocketpool"
)


// Get RPL total supply
func GetRPLTotalSupply(rp *rocketpool.RocketPool, opts *bind.CallOpts) (*big.Int, error) {
    rocketTokenRPL, err := getRocketTokenRPL(rp)
    if err != nil {
        return nil, err
    }
    return totalSupply(rocketTokenRPL, "RPL", opts)
}


// Get RPL balance
func GetRPLBalance(rp *rocketpool.RocketPool, address common.Address, opts *bind.CallOpts) (*big.Int, error) {
    rocketTokenRPL, err := getRocketTokenRPL(rp)
    if err != nil {
        return nil, err
    }
    return balanceOf(rocketTokenRPL, "RPL", address, opts)
}


// Get RPL allowance
func GetRPLAllowance(rp *rocketpool.RocketPool, owner, spender common.Address, opts *bind.CallOpts) (*big.Int, error) {
    rocketTokenRPL, err := getRocketTokenRPL(rp)
    if err != nil {
        return nil, err
    }
    return allowance(rocketTokenRPL, "RPL", owner, spender, opts)
}


// Transfer RPL
func TransferRPL(rp *rocketpool.RocketPool, to common.Address, amount *big.Int, opts *bind.TransactOpts) (*types.Receipt, error) {
    rocketTokenRPL, err := getRocketTokenRPL(rp)
    if err != nil {
        return nil, err
    }
    return transfer(rocketTokenRPL, "RPL", to, amount, opts)
}


// Approve an RPL spender
func ApproveRPL(rp *rocketpool.RocketPool, spender common.Address, amount *big.Int, opts *bind.TransactOpts) (*types.Receipt, error) {
    rocketTokenRPL, err := getRocketTokenRPL(rp)
    if err != nil {
        return nil, err
    }
    return approve(rocketTokenRPL, "RPL", spender, amount, opts)
}


// Transfer RPL from a sender
func TransferFromRPL(rp *rocketpool.RocketPool, from, to common.Address, amount *big.Int, opts *bind.TransactOpts) (*types.Receipt, error) {
    rocketTokenRPL, err := getRocketTokenRPL(rp)
    if err != nil {
        return nil, err
    }
    return transferFrom(rocketTokenRPL, "RPL", from, to, amount, opts)
}


// Get contracts
var rocketTokenRPLLock sync.Mutex
func getRocketTokenRPL(rp *rocketpool.RocketPool) (*rocketpool.Contract, error) {
    rocketTokenRPLLock.Lock()
    defer rocketTokenRPLLock.Unlock()
    return rp.GetContract("rocketTokenRPL")
}

