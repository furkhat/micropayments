package worker

import "github.com/ethereum/go-ethereum/accounts/abi"

var (
	approvalNonIndexArgs       abi.Arguments
	channelCreatedNonIndexArgs abi.Arguments
)

func init() {
	abiUint256, err := abi.NewType("uint256")
	if err != nil {
		panic(err)
	}

	abiUint192, err := abi.NewType("uint192")
	if err != nil {
		panic(err)
	}

	abiBytes32, err := abi.NewType("bytes32")
	if err != nil {
		panic(err)
	}

	approvalNonIndexArgs = abi.Arguments{{
		Type: abiUint256,
	}}

	channelCreatedNonIndexArgs = abi.Arguments{
		{
			Type: abiUint192,
		},
		{
			Type: abiBytes32,
		},
	}
}
