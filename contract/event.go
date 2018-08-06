package contract

import (
	"github.com/ethereum/go-ethereum/common"
)

// Logs digests.
var (
	// PSC logs.
	EthChannelCreated = common.HexToHash(
		"a6153987181667023837aee39c3f1a702a16e5e146323ef10fb96844a526143c")
	EthCooperativeChannelClose = common.HexToHash(
		"b488ea0f49970f556cf18e57588e78dcc1d3fd45c71130aa5099a79e8b06c8e7")

	// PTC logs.
	EthTokenApproval = common.HexToHash(
		"8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
	EthTokenTransfer = common.HexToHash(
		"ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
)
