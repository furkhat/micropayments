package worker

import "github.com/ethereum/go-ethereum/accounts/abi"

var (
	approvalNonIndexArgs abi.Arguments
)

func init() {
	abiUint256, err := abi.NewType("uint256")
	if err != nil {
		panic(err)
	}

	approvalNonIndexArgs = abi.Arguments{{
		Type: abiUint256,
	}}
}
