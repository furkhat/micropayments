package monitor

import (
	"math/big"
	"reflect"
	"testing"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/furkhat/micropayments/contract"
)

var timeout = time.Millisecond * time.Duration(30)

func TestCollectQuery(t *testing.T) {
	acc := newTestAccount()
	db.Insert(acc)
	defer db.Delete(acc)
	fakeConn := &fakeEthConn{number: 100}
	mon := NewMonitor(db, fakeConn, 1, pscAddr, ptcAddr)
	mon.lastSeenBlock = 10

	ticker := newMockTicker()
	go mon.start(ticker.C)
	ticker.tick()
	time.Sleep(timeout)

	addresses := []common.Hash{common.HexToAddress(acc.EthAddr).Hash()}
	expected := []ethereum.FilterQuery{
		{
			FromBlock: new(big.Int).SetUint64(11),
			ToBlock:   new(big.Int).SetUint64(99),
			Addresses: []common.Address{pscAddr, ptcAddr},
			Topics:    [][]common.Hash{nil, addresses},
		},
		{
			FromBlock: new(big.Int).SetUint64(11),
			ToBlock:   new(big.Int).SetUint64(99),
			Addresses: []common.Address{pscAddr, ptcAddr},
			Topics:    [][]common.Hash{nil, nil, addresses},
		},
	}

	if !reflect.DeepEqual(expected, fakeConn.filterQueries) {
		t.Fatal("unexpected filter queries")
	}
}

func TestHandlerCalls(t *testing.T) {
	fakeConn := &fakeEthConn{}
	fakeConn.fakeLogs = []ethtypes.Log{
		ethtypes.Log{
			Topics:      []common.Hash{contract.EthTokenApproval},
			BlockNumber: 100},
		ethtypes.Log{
			Topics:      []common.Hash{contract.EthTokenTransfer},
			BlockNumber: 101},
		ethtypes.Log{
			Topics:      []common.Hash{contract.EthCooperativeChannelClose},
			BlockNumber: 102},
		ethtypes.Log{
			Topics:      []common.Hash{contract.EthChannelCreated},
			BlockNumber: 103},
	}
	mon := NewMonitor(db, fakeConn, 1, pscAddr, ptcAddr)
	mon.lastSeenBlock = 10
	var approves, transfers, chanCloses, chanCreates int
	mon.RegisterWorker(contract.EthTokenApproval, func(*ethtypes.Log) error {
		approves++
		return nil
	})
	mon.RegisterWorker(contract.EthTokenTransfer, func(*ethtypes.Log) error {
		transfers++
		return nil
	})
	mon.RegisterWorker(contract.EthCooperativeChannelClose, func(*ethtypes.Log) error {
		chanCloses++
		return nil
	})
	mon.RegisterWorker(contract.EthChannelCreated, func(*ethtypes.Log) error {
		chanCreates++
		return nil
	})
	go mon.Start()
	time.Sleep(time.Millisecond * time.Duration(10))

	if approves != 1 || transfers != 1 || chanCloses != 1 || chanCreates != 1 {
		t.Fatal("wrong number of logs collected")
	}

	if mon.lastSeenBlock != 103 {
		t.Fatal("last seen block not updated")
	}
}
