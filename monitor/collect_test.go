package monitor

import (
	"math/big"
	"reflect"
	"sync"
	"testing"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/furkhat/micropayments/contract"
	"github.com/furkhat/micropayments/data"
)

func TestCollectQuery(t *testing.T) {
	td := newTestData(t, 100, 10)
	defer td.cleanUp(t)

	acc := data.NewAccount()
	insertToTestDB(t, acc)
	defer deleteFromTestDB(t, acc)
	ticker := newMockTicker()
	go td.mon.start(ticker.C)
	ticker.tick()
	time.Sleep(time.Millisecond * time.Duration(30))

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

	if !reflect.DeepEqual(expected, td.fakeConn.filterQueries) {
		t.Fatal("unexpected filter queries")
	}
}

func TestHandlerCalls(t *testing.T) {
	td := newTestData(t, 103, 10)
	defer td.cleanUp(t)

	td.fakeConn.fakeLogs = []ethtypes.Log{
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
	var wg sync.WaitGroup
	wg.Add(4)
	success := make(chan struct{})
	go func() {
		wg.Wait()
		success <- struct{}{}
	}()
	markOneDone := func(*ethtypes.Log) error {
		wg.Done()
		return nil
	}
	td.mon.RegisterWorker(contract.EthTokenApproval, markOneDone)
	td.mon.RegisterWorker(contract.EthTokenTransfer, markOneDone)
	td.mon.RegisterWorker(contract.EthCooperativeChannelClose, markOneDone)
	td.mon.RegisterWorker(contract.EthChannelCreated, markOneDone)
	go td.mon.Start()

	select {
	case <-time.After(time.Second):
		t.Fatal("not all workers has not been invoked")
	case <-success:
	}

	if td.mon.lastSeenBlock != 103 {
		t.Fatal("last seen block not updated")
	}

	blockSettings := &data.Setting{}
	err := db.FindByPrimaryKeyTo(blockSettings, data.SettingsLastSeenBlock)
	if err != nil {
		t.Fatal("could not find last seen block settings: ", err)
	}
	if blockSettings.Value != "103" {
		t.Fatal("last seen block not updated")
	}
}
